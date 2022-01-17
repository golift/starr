package starr

import (
	"net/url"
	"strconv"
	"strings"
)

/* This file containers helper methods and types for page-able API calls.
 * Like GetHistory() and GetQueue().
 */

// Req is the input to search requests that have page-able responses.
// These are turned into HTTP parameters.
type Req struct {
	PageSize   int       // 10 is default if not provided.
	Page       int       // 1 is default if not provided.
	SortKey    string    // date, timeleft, others?
	SortDir    Sorting   // ascending, descending
	Filter     Filtering // enums for eventTypes. App specific.
	url.Values           // Additional values that may be set.
}

// Sorting is used as a request parameter value to sort lists, like History and Queue.
type Sorting string

const (
	// SortAsc is the default, and sorts lists in ascending order.
	SortAscend Sorting = "ascending"
	// SortDesc flips the sort order to descending.
	SortDescend Sorting = "descending"
)

// Filtering is used as a request parameter value to filter lists, like History and Queue.
// The filter values are different per-app, so find their values in their respective modules.
type Filtering int

// Set makes sure the sort direction is valid.
func (s *Sorting) Set(val string) {
	switch Sorting(strings.ToLower(val)) {
	default:
		fallthrough
	case SortAscend:
		*s = SortAscend
	case SortDescend:
		*s = SortDescend
	}
}

// Param returns the string value of a Filter eventType.
func (f Filtering) Param() string {
	return strconv.Itoa(int(f))
}

// Params returns a brand new url.Values with all request parameters combined.
func (r *Req) Params() url.Values {
	params := make(url.Values)

	if r.Filter > 0 {
		params.Set("eventType", r.Filter.Param())
	}

	if r.Page > 0 {
		params.Set("page", strconv.Itoa(r.Page))
	} else {
		params.Set("page", "1")
	}

	if r.PageSize > 0 {
		params.Set("pageSize", strconv.Itoa(r.PageSize))
	} else {
		params.Set("pageSize", "10")
	}

	if r.SortKey != "" {
		params.Set("sortKey", r.SortKey)
	} else {
		params.Set("sortKey", "date") // timeleft, title, id
	}

	if r.SortDir != "" {
		params.Set("sortDirection", string(r.SortDir))
	} else {
		params.Set("sortDirection", "ascending") // descending
	}

	for k, v := range r.Values {
		for _, val := range v {
			params.Set(k, val)
		}
	}

	return params
}

// Encode turns our request parameters into a URI string.
func (r *Req) Encode() string {
	return r.Params().Encode()
}

// CheckSet sets a request parameter if it's not already set.
func (r *Req) CheckSet(key, value string) { //nolint:cyclop
	switch strings.ToLower(key) {
	case "page":
		if r.Page == 0 {
			r.Page, _ = strconv.Atoi(value)
		}
	case "pagesize":
		if r.PageSize == 0 {
			r.PageSize, _ = strconv.Atoi(value)
		}
	case "sortkey":
		if r.SortKey == "" {
			r.SortKey = value
		}
	case "sortdirection":
		if r.SortDir == "" {
			r.SortDir.Set(value)
		}
	default:
		if r.Values == nil {
			r.Values = make(url.Values)
		}

		if r.Values.Get(key) == "" {
			r.Values.Set(key, value)
		}
	}
}

// Set sets a request parameter.
func (r *Req) Set(key, value string) {
	switch strings.ToLower(key) {
	case "page":
		r.Page, _ = strconv.Atoi(value)
	case "pagesize":
		r.PageSize, _ = strconv.Atoi(value)
	case "sortkey":
		r.SortKey = value
	case "sortdirection":
		r.SortDir.Set(value)
	default:
		if r.Values == nil {
			r.Values = make(url.Values)
		}

		r.Values.Set(key, value)
	}
}

// SetPerPage returns a proper perPage value that is not equal to zero,
// and not larger than the record count desired. If the count is zero, then
// perPage can be anything other than zero.
// This is used by paginated methods in the starr modules.
func SetPerPage(records, perPage int) int {
	const perPageDefault = 500

	if perPage <= 1 {
		if records > perPageDefault || records == 0 {
			perPage = perPageDefault
		} else {
			perPage = records
		}
	} else if perPage > records && records != 0 {
		perPage = records
	}

	return perPage
}

// AdjustPerPage to make sure we don't go over, or ask for more records than exist.
// This is used by paginated methods in the starr modules.
// 'records' is the number requested, 'total' is the number in the app,
// 'collected' is how many we have so far, and 'perPage' is the current perPage setting.
func AdjustPerPage(records, total, collected, perPage int) int {
	// Do not ask for more than was requested.
	if d := records - collected; perPage > d && d > 0 {
		perPage = d
	}

	// Ask for only the known total.
	if d := total - collected; perPage > d {
		perPage = d
	}

	return perPage
}
