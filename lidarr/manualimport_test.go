package lidarr_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/starrtest"
)

func TestManualImport(t *testing.T) {
	t.Parallel()

	basePath := path.Join("/", starr.API, lidarr.APIver, "manualimport")
	params := &lidarr.ManualImportParams{
		Folder:               "/music/album",
		DownloadID:           "dl123",
		ArtistID:             329,
		ReplaceExistingFiles: true,
		FilterExistingFiles:  false,
	}

	query := make(url.Values)
	query.Add("folder", params.Folder)
	query.Add("downloadId", params.DownloadID)
	query.Add("artistId", starr.Str(params.ArtistID))
	query.Add("replaceExistingFiles", starr.Str(params.ReplaceExistingFiles))
	query.Add("filterExistingFiles", starr.Str(params.FilterExistingFiles))

	expectedPath := basePath + "?" + query.Encode()
	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   expectedPath,
			ResponseStatus: http.StatusOK,
			ResponseBody: `[
				{"id":1,"path":"/music/album/01-track.flac","name":"01-track.flac","size":12345,"artist":{"id":329},"album":{"id":2826},"albumReleaseId":11727,"tracks":[{"id":220502}],"quality":{"quality":{"id":21,"name":"FLAC 24bit"}},"downloadId":"dl123","disableReleaseSwitching":false},
				{"id":2,"path":"/music/album/02-track.flac","name":"02-track.flac","size":12345,"artist":{"id":329},"album":{"id":2826},"albumReleaseId":11727,"tracks":[{"id":220503}],"quality":{"quality":{"id":21,"name":"FLAC 24bit"}},"downloadId":"dl123","disableReleaseSwitching":false}
			]`,
			WithError:       nil,
			ExpectedMethod:  "GET",
			ExpectedRequest: "",
			WithResponse: []*lidarr.ManualImportOutput{
				{ID: 1, Path: "/music/album/01-track.flac", Name: "01-track.flac", Size: 12345, Artist: &lidarr.Artist{ID: 329}, Album: &lidarr.Album{ID: 2826}, AlbumReleaseID: 11727, Tracks: []*lidarr.Track{{ID: 220502}}, Quality: &starr.Quality{Quality: &starr.BaseQuality{ID: 21, Name: "FLAC 24bit"}}, DownloadID: "dl123", DisableReleaseSwitching: false},
				{ID: 2, Path: "/music/album/02-track.flac", Name: "02-track.flac", Size: 12345, Artist: &lidarr.Artist{ID: 329}, Album: &lidarr.Album{ID: 2826}, AlbumReleaseID: 11727, Tracks: []*lidarr.Track{{ID: 220503}}, Quality: &starr.Quality{Quality: &starr.BaseQuality{ID: 21, Name: "FLAC 24bit"}}, DownloadID: "dl123", DisableReleaseSwitching: false},
			},
		},
		{
			Name:           "404",
			ExpectedPath:   expectedPath,
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			ExpectedMethod: "GET",
			WithResponse:   []*lidarr.ManualImportOutput(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.ManualImport(params)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestManualImportContext(t *testing.T) {
	t.Parallel()

	basePath := path.Join("/", starr.API, lidarr.APIver, "manualimport")
	params := &lidarr.ManualImportParams{
		Folder:               "/music/album",
		DownloadID:           "dl456",
		ArtistID:             100,
		ReplaceExistingFiles: false,
		FilterExistingFiles:  true,
	}

	query := make(url.Values)
	query.Add("folder", params.Folder)
	query.Add("downloadId", params.DownloadID)
	query.Add("artistId", starr.Str(params.ArtistID))
	query.Add("replaceExistingFiles", starr.Str(params.ReplaceExistingFiles))
	query.Add("filterExistingFiles", starr.Str(params.FilterExistingFiles))

	expectedPath := basePath + "?" + query.Encode()
	mockData := &starrtest.MockData{
		Name:           "200_empty",
		ExpectedPath:   expectedPath,
		ResponseStatus: http.StatusOK,
		ResponseBody:   "[]",
		ExpectedMethod: "GET",
		WithResponse:   []*lidarr.ManualImportOutput{},
	}
	mockServer := mockData.GetMockServer(t)
	client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))

	output, err := client.ManualImportContext(context.Background(), params)
	require.NoError(t, err)
	assert.Empty(t, output)
}

func TestManualImportCommandFromOutputs(t *testing.T) {
	t.Parallel()

	t.Run("nil_returns_nil", func(t *testing.T) {
		t.Parallel()
		out := lidarr.ManualImportCommandFromOutputs(nil, true)
		assert.Nil(t, out)
	})

	t.Run("empty_slice_returns_nil", func(t *testing.T) {
		t.Parallel()
		out := lidarr.ManualImportCommandFromOutputs([]*lidarr.ManualImportOutput{}, true)
		assert.Nil(t, out)
	})

	t.Run("all_nil_entries_returns_nil", func(t *testing.T) {
		t.Parallel()
		out := lidarr.ManualImportCommandFromOutputs([]*lidarr.ManualImportOutput{nil, nil}, true)
		assert.Nil(t, out)
	})

	t.Run("single_output", func(t *testing.T) {
		t.Parallel()

		outputs := []*lidarr.ManualImportOutput{
			{
				Path:                    "/music/01.flac",
				Artist:                  &lidarr.Artist{ID: 329},
				Album:                   &lidarr.Album{ID: 2826},
				AlbumReleaseID:          11727,
				Tracks:                  []*lidarr.Track{{ID: 220502}},
				Quality:                 &starr.Quality{Quality: &starr.BaseQuality{ID: 21, Name: "FLAC 24bit"}},
				DownloadID:              "dl123",
				DisableReleaseSwitching: true,
			},
		}

		cmd := lidarr.ManualImportCommandFromOutputs(outputs, true)
		require.NotNil(t, cmd)
		assert.Equal(t, "ManualImport", cmd.Name)
		assert.True(t, cmd.ReplaceExistingFiles)
		assert.Equal(t, "auto", cmd.ImportMode)
		require.Len(t, cmd.Files, 1)
		assert.Equal(t, "/music/01.flac", cmd.Files[0].Path)
		assert.Equal(t, int64(329), cmd.Files[0].ArtistID)
		assert.Equal(t, int64(2826), cmd.Files[0].AlbumID)
		assert.Equal(t, int64(11727), cmd.Files[0].AlbumReleaseID)
		assert.Equal(t, []int64{220502}, cmd.Files[0].TrackIDs)
		assert.Equal(t, "dl123", cmd.Files[0].DownloadID)
		assert.True(t, cmd.Files[0].DisableReleaseSwitching)
		assert.Equal(t, 0, cmd.Files[0].IndexerFlags)
	})

	t.Run("nil_artist_album_use_zero_ids", func(t *testing.T) {
		t.Parallel()

		outputs := []*lidarr.ManualImportOutput{
			{
				Path:       "/music/01.flac",
				Artist:     nil,
				Album:      nil,
				Tracks:     []*lidarr.Track{{ID: 99}},
				Quality:    nil,
				DownloadID: "dl",
			},
		}

		cmd := lidarr.ManualImportCommandFromOutputs(outputs, false)
		require.NotNil(t, cmd)
		require.Len(t, cmd.Files, 1)
		assert.Equal(t, int64(0), cmd.Files[0].ArtistID)
		assert.Equal(t, int64(0), cmd.Files[0].AlbumID)
		assert.Equal(t, []int64{99}, cmd.Files[0].TrackIDs)
		assert.NotNil(t, cmd.Files[0].Quality)
		assert.False(t, cmd.ReplaceExistingFiles)
	})

	t.Run("skips_nil_tracks", func(t *testing.T) {
		t.Parallel()

		outputs := []*lidarr.ManualImportOutput{
			{
				Path:    "/music/01.flac",
				Artist:  &lidarr.Artist{ID: 1},
				Album:   &lidarr.Album{ID: 2},
				Tracks:  []*lidarr.Track{nil, {ID: 10}, nil},
				Quality: &starr.Quality{},
			},
		}
		cmd := lidarr.ManualImportCommandFromOutputs(outputs, true)
		require.NotNil(t, cmd)
		require.Len(t, cmd.Files, 1)
		assert.Equal(t, []int64{10}, cmd.Files[0].TrackIDs)
	})
}

func TestManualImportReprocess(t *testing.T) {
	t.Parallel()

	manualImportReq := &lidarr.ManualImportInput{
		ID:                      1,
		Path:                    "/music/01.flac",
		Name:                    "01.flac",
		ArtistID:                329,
		AlbumID:                 2826,
		AlbumReleaseID:          11727,
		TrackIDs:                []int64{220502},
		Quality:                 &starr.Quality{Quality: &starr.BaseQuality{ID: 21, Name: "FLAC 24bit"}},
		DownloadID:              "dl123",
		ReplaceExistingFiles:    true,
		DisableReleaseSwitching: false,
	}
	expectedBodyBuf := new(bytes.Buffer)
	require.NoError(t, json.NewEncoder(expectedBodyBuf).Encode(manualImportReq))

	expectedBody := expectedBodyBuf.String()
	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "manualimport"),
			ResponseStatus:  http.StatusOK,
			ResponseBody:    "{}",
			WithError:       nil,
			WithRequest:     manualImportReq,
			ExpectedRequest: expectedBody,
			ExpectedMethod:  "POST",
		},
		{
			Name:            "404",
			ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "manualimport"),
			ResponseStatus:  http.StatusNotFound,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithRequest:     manualImportReq,
			ExpectedRequest: expectedBody,
			ExpectedMethod:  "POST",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			reqCopy := *manualImportReq
			err := client.ManualImportReprocess(&reqCopy)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}

func TestManualImportReprocessContext(t *testing.T) {
	t.Parallel()

	manualImportReq := &lidarr.ManualImportInput{
		ID:         2,
		Path:       "/music/02.flac",
		Name:       "02.flac",
		ArtistID:   100,
		AlbumID:    200,
		DownloadID: "dl999",
	}
	expectedBodyBuf := new(bytes.Buffer)
	require.NoError(t, json.NewEncoder(expectedBodyBuf).Encode(manualImportReq))

	expectedBody := expectedBodyBuf.String()
	mockData := &starrtest.MockData{
		Name:            "200",
		ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "manualimport"),
		ResponseStatus:  http.StatusOK,
		ResponseBody:    "{}",
		ExpectedRequest: expectedBody,
		ExpectedMethod:  "POST",
	}
	mockServer := mockData.GetMockServer(t)
	client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))

	err := client.ManualImportReprocessContext(context.Background(), manualImportReq)
	require.NoError(t, err)
}
