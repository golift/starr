// Package starrcmd provides the bindings to consume a custom script command hook from any Starr app.
// Create these by going into Settings->Connect->Custom Script in Lidarr, Prowlarr, Radarr, Readarr, or Sonarr.
// See the included example_test.go file for examples on how to use this module.
package starrcmd

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// get offloads the error checking from all the other routines.
// This is where our journey into the data truly begins.
func (c *CmdEvent) get(wanted Event, output interface{}) error {
	if c.Type != wanted {
		return fmt.Errorf("%w: requested '%s' have '%s'", ErrInvalidEvent, wanted, c.Type)
	}

	if err := fillStructFromEnv(output); err != nil {
		return fmt.Errorf("reading environment: %w", err)
	}

	return nil
}

// This does not traverse structs and will only stay on normal members.
func fillStructFromEnv(dataStruct interface{}) error {
	field := reflect.ValueOf(dataStruct)
	if field.Kind() != reflect.Ptr || field.Elem().Kind() != reflect.Struct {
		panic("yuh dun ate in sumthin bahd! This is a bug in the starrcmd library.")
	}

	t := field.Type().Elem()
	for idx := 0; idx < t.NumField(); idx++ { // Loop each struct member
		split := strings.SplitN(t.Field(idx).Tag.Get("env"), ",", 2) //nolint:gomnd

		tag := strings.ToLower(split[0]) // lower to protect naming mistakes.
		if !field.Elem().Field(idx).CanSet() || tag == "-" || tag == "" {
			continue // This only works with non-empty reflection tags on exported members.
		}

		// If the tag has a comma, the value that follows is used to split strings into []string.
		var splitVal string
		if len(split) == 2 { //nolint:gomnd
			splitVal = split[1]
		}

		value := os.Getenv(tag)
		if value == "" {
			// fmt.Println("skipping", tag)
			continue
		}

		err := parseStructMember(field.Elem().Field(idx), value, splitVal)
		if err != nil {
			return fmt.Errorf("%s: (%s) %w", tag, os.Getenv(tag), err)
		}
	}

	return nil
}

/* All of the code below was taken from the golift.io/cnfg module. */

// This is trimmed and does not parse some types.
func parseStructMember(field reflect.Value, value, splitVal string) error {
	var err error

	switch fieldType := field.Type().String(); fieldType {
	// Handle each member type appropriately (differently).
	case "string":
		// SetString is a reflect package method to update a struct member by index.
		field.SetString(value)
	case "int", "int64":
		var val int64

		val, err = parseInt(fieldType, value)
		field.SetInt(val)
		/*
			case "float64":
				// uncomment float64 if needed.
				var val float64
				//nolint:gomnd
				val, err = strconv.ParseFloat(value, 64)
				field.SetFloat(val)
			case "time.Duration":
				// this needs to be fixed to work with any duration values we find in starr apps.
				var val time.Duration

				val, err = time.ParseDuration(value)
				field.Set(reflect.ValueOf(val))
		*/
	case "time.Time":
		var val time.Time

		val, err = time.Parse(DateFormat, value)
		field.Set(reflect.ValueOf(val))

	case "bool":
		var val bool

		val, err = strconv.ParseBool(value)
		field.SetBool(val)
	default:
		if missing, err := parseSlices(field, value, splitVal); err != nil {
			return fmt.Errorf("%s: %w", value, err)
		} else if missing {
			panic(fmt.Sprintf("invalid type provided to parser, this is a bug in the starrcmd library: %s (%s)",
				fieldType, value))
		}
	}

	if err != nil {
		return fmt.Errorf("%s: %w", value, err)
	}

	return nil
}

func parseSlices(field reflect.Value, value, splitVal string) (bool, error) {
	if splitVal == "" {
		// this will trigger a panic() if you forget a splitVal on an env tag.
		return true, nil
	}

	var err error

	switch fieldType := field.Type().String(); fieldType {
	default:
		return true, nil
	case "[]time.Time":
		split := strings.Split(value, splitVal)
		vals := make([]time.Time, len(split))

		for idx, val := range split {
			if vals[idx], err = time.Parse(DateFormat, val); err != nil {
				return false, fmt.Errorf("(%s) %s: %w", splitVal, val, err)
			}
		}

		field.Set(reflect.ValueOf(vals))
	case "[]int64":
		split := strings.Split(value, splitVal)
		vals := make([]int64, len(split))

		for idx, val := range split {
			if vals[idx], err = parseInt(fieldType, val); err != nil {
				return false, fmt.Errorf("%s: %w", value, err)
			}
		}

		field.Set(reflect.ValueOf(vals))
	case "[]string":
		vals := strings.Split(value, splitVal)
		field.Set(reflect.ValueOf(vals))
	}

	return false, err
}

// parseInt parses an integer from a string as specific size.
// If you need int8, 16 or 32, add them...
func parseInt(intType, envval string) (i int64, err error) {
	//nolint:gomnd
	switch intType {
	default:
		i, err = strconv.ParseInt(envval, 10, 0)
	case "int64":
		i, err = strconv.ParseInt(envval, 10, 64)
	}

	if err != nil { // this error may prove to suck...
		return i, fmt.Errorf("parsing integer: %w", err)
	}

	return i, nil
}
