package starr_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
)

func TestCalendarTimeFilterFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{"midnight", time.Date(2020, 2, 20, 0, 30, 0, 0, time.UTC), "2020-02-20T00:30:00.000Z"},
		{"afternoon", time.Date(2020, 2, 20, 15, 4, 5, 0, time.UTC), "2020-02-20T15:04:05.000Z"},
		{"last minute", time.Date(2020, 2, 20, 23, 59, 59, 0, time.UTC), "2020-02-20T23:59:59.000Z"},
		{"noon", time.Date(2020, 2, 20, 12, 0, 0, 0, time.UTC), "2020-02-20T12:00:00.000Z"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// The apps read this filter as a 24 hour clock; a 12 hour format shifts afternoon times by 12 hours.
			assert.Equal(t, test.want, test.input.UTC().Format(starr.CalendarTimeFilterFormat))
		})
	}
}

func TestQueueDeleteOpts_Values(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	var opts *starr.QueueDeleteOpts

	params := opts.Values() // it's nil.
	require.Equal(t, "removeFromClient=true", params.Encode(),
		"default queue delete parameters encoded incorrectly")

	opts = &starr.QueueDeleteOpts{
		BlockList:        true,
		RemoveFromClient: starr.False(),
		SkipRedownload:   true,
	}
	params = opts.Values()

	assert.Equal("false", params.Get("removeFromClient"), "delete parameters encoded incorrectly")
	assert.Equal("true", params.Get("blocklist"), "delete parameters encoded incorrectly")
	assert.Equal("true", params.Get("skipRedownload"), "delete parameters encoded incorrectly")
}

func TestPlayTime_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		json     string
		wantOrig string
		wantDur  time.Duration
	}{
		{"hh:mm:ss", `"01:23:45"`, "01:23:45", 1*time.Hour + 23*time.Minute + 45*time.Second},
		{"hh:mm:ss.fraction", `"00:03:31.1750999"`, "00:03:31.1750999", 3*time.Minute + 31*time.Second + 175099900*time.Nanosecond},
		{"mm:ss", `"12:34"`, "12:34", 12*time.Minute + 34*time.Second},
		{"mm:ss.fraction", `"05:30.5"`, "05:30.5", 5*time.Minute + 30*time.Second + 500*time.Millisecond},
		{"ss only", `"90"`, "90", 90 * time.Second},
		{"ss.fraction", `"45.25"`, "45.25", 45*time.Second + 250*time.Millisecond},
		{"zero", `"00:00:00"`, "00:00:00", 0},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var playTime starr.PlayTime

			err := json.Unmarshal([]byte(testCase.json), &playTime)
			require.NoError(t, err)
			assert.Equal(t, testCase.wantOrig, playTime.Original, "Original")
			assert.Equal(t, testCase.wantDur, playTime.Duration, "Duration")
		})
	}
}

func TestPlayTime_MarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("round_trip_preserves_original", func(t *testing.T) {
		t.Parallel()

		inputs := []string{`"01:23:45"`, `"00:03:31.1750999"`, `"12:34"`, `"90"`}
		for _, input := range inputs {
			var playTime starr.PlayTime
			require.NoError(t, json.Unmarshal([]byte(input), &playTime))
			data, err := json.Marshal(&playTime)
			require.NoError(t, err)
			assert.Equal(t, input, string(data), "round-trip for %s", input)
		}
	})

	t.Run("programmatic_formats_as_hh_mm_ss", func(t *testing.T) {
		t.Parallel()

		playTime := starr.PlayTime{Duration: 3*time.Minute + 31*time.Second + 175099900*time.Nanosecond}
		data, err := json.Marshal(&playTime)
		require.NoError(t, err)

		var back starr.PlayTime

		require.NoError(t, json.Unmarshal(data, &back))
		// Allow small drift from float rounding in format/parse
		assert.InDelta(t, float64(playTime.Duration), float64(back.Duration), float64(time.Millisecond),
			"marshal from Duration only should round-trip to equivalent duration")
	})

	t.Run("zero_duration", func(t *testing.T) {
		t.Parallel()

		playTime := starr.PlayTime{}
		data, err := json.Marshal(&playTime)
		require.NoError(t, err)
		assert.Equal(t, `"00:00:00"`, string(data))
	})
}
