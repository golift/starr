package sonarr_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

func TestAPIVersionConstructors(t *testing.T) {
	t.Parallel()

	client := sonarr.New(starr.New("k", "http://sonarr", 0))
	assert.Equal(t, sonarr.APIver, client.APIVersion())
	assert.Equal(t, sonarr.APIv5, client.WithAPI(sonarr.APIv5).APIVersion())
	assert.Equal(t, sonarr.APIver, client.WithAPI("nope").APIVersion())
	assert.Equal(t, sonarr.APIv5, sonarr.NewV5(starr.New("k", "http://sonarr", 0)).APIVersion())

	separate := client.WithAPI(sonarr.APIv5)
	client.UseAPI(sonarr.APIv5)
	assert.Equal(t, sonarr.APIv5, client.APIVersion())
	client.UseAPI("nope")
	assert.Equal(t, sonarr.APIver, client.APIVersion())
	assert.Equal(t, sonarr.APIv5, separate.APIVersion())
}

func TestUseAPIConcurrent(t *testing.T) {
	t.Parallel()

	client := sonarr.New(starr.New("k", "http://sonarr", 0))

	var wait sync.WaitGroup
	for range 32 {
		wait.Go(func() {
			client.UseAPI(sonarr.APIv5)
			_ = client.APIVersion()
			client.UseAPI(sonarr.APIver)
		})
	}

	wait.Wait()
	assert.Equal(t, sonarr.APIver, client.APIVersion())
}

func TestUseAPIRewritesTheNextRequest(t *testing.T) {
	t.Parallel()

	var path string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		path = req.URL.Path
		_, _ = writer.Write([]byte(`[]`))
	}))
	t.Cleanup(server.Close)

	client := sonarr.New(starr.New("k", server.URL, 0))
	_, err := client.GetQualityProfiles()
	require.NoError(t, err)
	assert.Equal(t, "/api/v3/qualityProfile", path)

	client.UseAPI(sonarr.APIv5)
	_, err = client.GetQualityProfiles()
	require.NoError(t, err)
	assert.Equal(t, "/api/v5/qualityProfile", path)

	client.UseAPI(sonarr.APIver)
	_, err = client.GetQualityProfiles()
	require.NoError(t, err)
	assert.Equal(t, "/api/v3/qualityProfile", path)
}

func TestV5Paths(t *testing.T) {
	t.Parallel()

	var gotMethod, gotPath, gotBody string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		gotMethod = req.Method
		gotPath = req.URL.RequestURI()

		body, err := io.ReadAll(req.Body)
		assert.NoError(t, err)

		gotBody = string(body)

		writer.WriteHeader(http.StatusOK)

		response := []byte(`{}`)
		if req.URL.Path == "/api/v5/connection" || req.URL.Path == "/api/v5/qualityDefinition" {
			response = []byte(`[]`)
		}

		_, _ = writer.Write(response)
	}))
	t.Cleanup(server.Close)

	client := sonarr.NewV5(starr.New("k", server.URL, 0))

	_, err := client.GetNaming()
	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, gotMethod)
	assert.Equal(t, "/api/v5/settings/naming", gotPath)

	_, err = client.GetNotifications()
	require.NoError(t, err)
	assert.Equal(t, "/api/v5/connection", gotPath)

	size := 1.5
	_, err = client.UpdateQualityDefinitions([]*sonarr.QualityDefinition{{
		ID:       3,
		Title:    "HDTV",
		MinSize:  size,
		MaxSize:  size,
		PrefSize: size,
	}})
	require.NoError(t, err)
	assert.Equal(t, http.MethodPut, gotMethod)
	assert.Equal(t, "/api/v5/qualityDefinition", gotPath)
	assert.NotContains(t, gotBody, "minSize")
	assert.NotContains(t, gotBody, "maxSize")
	assert.NotContains(t, gotBody, "preferredSize")
	assert.Contains(t, gotBody, `"title":"HDTV"`)

	_, err = client.GetFeed(sonarr.Feed{})
	require.NoError(t, err)
	assert.Contains(t, gotPath, "feed/v5/calendar/sonarr.ics")
	assert.NotContains(t, gotPath, "feed/v3/")

	err = client.Ping()
	require.NoError(t, err)
	assert.Equal(t, "/ping", gotPath)
}

func TestV3QualityDefinitionSendsSizes(t *testing.T) {
	t.Parallel()

	var gotPath, gotBody string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		gotPath = req.URL.RequestURI()

		body, err := io.ReadAll(req.Body)
		assert.NoError(t, err)

		gotBody = string(body)

		writer.WriteHeader(http.StatusOK)

		if req.Method == http.MethodPut && req.URL.Path != "/api/v3/qualityDefinition/update" {
			_, _ = writer.Write([]byte(`{}`))
			return
		}

		_, _ = writer.Write([]byte(`[]`))
	}))
	t.Cleanup(server.Close)

	client := sonarr.New(starr.New("k", server.URL, 0))
	definition := &sonarr.QualityDefinition{ID: 3, Title: "HDTV", MinSize: 1, MaxSize: 2, PrefSize: 1.5}

	_, err := client.UpdateQualityDefinitions([]*sonarr.QualityDefinition{definition})
	require.NoError(t, err)
	assert.Equal(t, "/api/v3/qualityDefinition/update", gotPath)
	assert.Contains(t, gotBody, `"minSize":1`)
	assert.Contains(t, gotBody, `"maxSize":2`)
	assert.Contains(t, gotBody, `"preferredSize":1.5`)

	_, err = client.UpdateQualityDefinition(definition)
	require.NoError(t, err)
	assert.Equal(t, "/api/v3/qualityDefinition/3", gotPath)
	assert.Contains(t, gotBody, `"minSize":1`)
	assert.Contains(t, gotBody, `"maxSize":2`)
	assert.Contains(t, gotBody, `"preferredSize":1.5`)
}

func TestV5QualityDefinitionOmitsSizes(t *testing.T) {
	t.Parallel()

	var gotPath, gotBody string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		gotPath = req.URL.RequestURI()

		body, err := io.ReadAll(req.Body)
		assert.NoError(t, err)

		gotBody = string(body)

		writer.WriteHeader(http.StatusOK)

		if req.URL.Path == "/api/v5/qualityDefinition/3" {
			_, _ = writer.Write([]byte(`{}`))
			return
		}

		_, _ = writer.Write([]byte(`[]`))
	}))
	t.Cleanup(server.Close)

	client := sonarr.NewV5(starr.New("k", server.URL, 0))
	definition := &sonarr.QualityDefinition{ID: 3, Title: "HDTV", MinSize: 1, MaxSize: 2, PrefSize: 1.5}

	_, err := client.UpdateQualityDefinition(definition)
	require.NoError(t, err)
	assert.Equal(t, "/api/v5/qualityDefinition/3", gotPath)
	assert.NotContains(t, gotBody, "minSize")
	assert.NotContains(t, gotBody, "maxSize")
	assert.NotContains(t, gotBody, "preferredSize")
	assert.Contains(t, gotBody, `"title":"HDTV"`)
}

func TestQualityProfileItemSizesRoundTrip(t *testing.T) {
	t.Parallel()

	const payload = `{"name":"HD","allowed":true,"minSize":1.5,"maxSize":9,"preferredSize":4.25,` +
		`"quality":{"id":4,"name":"HDTV-720p"}}`

	var item starr.Quality
	require.NoError(t, json.Unmarshal([]byte(payload), &item))
	require.NotNil(t, item.MinSize)
	require.NotNil(t, item.MaxSize)
	require.NotNil(t, item.PreferredSize)
	assert.InDelta(t, 1.5, *item.MinSize, 0)
	assert.InDelta(t, 9.0, *item.MaxSize, 0)
	assert.InDelta(t, 4.25, *item.PreferredSize, 0)

	out, err := json.Marshal(&item)
	require.NoError(t, err)
	assert.JSONEq(t, payload, string(out))

	bare, err := json.Marshal(&starr.Quality{Name: "x"})
	require.NoError(t, err)
	assert.NotContains(t, string(bare), "Size")
}

func TestV5QueueDeleteQuery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		opts    *starr.QueueDeleteOpts
		removal string
		block   string
	}{
		{
			name:    "defaults",
			removal: "RemoveFromClient",
			block:   "DoNotBlocklist",
		},
		{
			name: "change category and search",
			opts: &starr.QueueDeleteOpts{
				RemoveFromClient: starr.False(),
				ChangeCategory:   true,
				BlockList:        true,
			},
			removal: "ChangeCategory",
			block:   "BlocklistAndSearch",
		},
		{
			name: "ignore and block only",
			opts: &starr.QueueDeleteOpts{
				RemoveFromClient: starr.False(),
				BlockList:        true,
				SkipRedownload:   true,
			},
			removal: "Ignore",
			block:   "BlocklistOnly",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var got *http.Request

			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
				got = req

				writer.WriteHeader(http.StatusOK)
			}))
			t.Cleanup(server.Close)

			client := sonarr.NewV5(starr.New("k", server.URL, 0))
			require.NoError(t, client.DeleteQueue(7, test.opts))
			require.NotNil(t, got)
			assert.Equal(t, "/api/v5/queue/7", got.URL.Path)
			assert.Equal(t, test.removal, got.URL.Query().Get("removalMethod"))
			assert.Equal(t, test.block, got.URL.Query().Get("blocklistMethod"))
			assert.Empty(t, got.URL.Query().Get("removeFromClient"))
			assert.Empty(t, got.URL.Query().Get("blocklist"))
		})
	}
}

func TestV5QueueDecode(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/api/v5/queue", req.URL.Path)

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"page":1,"pageSize":1,"totalRecords":1,"records":[{` +
			`"id":3,"sizeLeft":12.5,"episodeIds":[9,8],"title":"ep"}]}`))
	}))
	t.Cleanup(server.Close)

	client := sonarr.NewV5(starr.New("k", server.URL, 0))
	queue, err := client.GetQueuePage(&starr.PageReq{Page: 1, PageSize: 1})
	require.NoError(t, err)
	require.Len(t, queue.Records, 1)
	assert.InDelta(t, 12.5, queue.Records[0].Sizeleft, 0.001)
	assert.Equal(t, []int64{9, 8}, queue.Records[0].EpisodeIDs)
	assert.Equal(t, int64(9), queue.Records[0].EpisodeID)
}

func TestReleaseProfileWrite(t *testing.T) {
	t.Parallel()

	t.Run("v5 all indexers", func(t *testing.T) {
		t.Parallel()

		body := releaseProfileBody(t, sonarr.NewV5, &sonarr.ReleaseProfile{
			Name:       "all",
			Enabled:    true,
			IndexerID:  0,
			IndexerIDs: []int64{0},
		})
		assert.Contains(t, body, `"indexerIds":[]`)
		assert.NotContains(t, body, `"indexerId"`)
		assert.NotContains(t, body, "preferred")
		assert.Contains(t, body, `"airDateGracePeriod":0`)
	})

	t.Run("v5 single indexer from v3 id", func(t *testing.T) {
		t.Parallel()

		body := releaseProfileBody(t, sonarr.NewV5, &sonarr.ReleaseProfile{IndexerID: 4})
		assert.Contains(t, body, `"indexerIds":[4]`)
		assert.NotContains(t, body, `"indexerId"`)
	})

	t.Run("v5 explicit empty list ignores legacy id", func(t *testing.T) {
		t.Parallel()

		body := releaseProfileBody(t, sonarr.NewV5, &sonarr.ReleaseProfile{
			Name:       "all",
			IndexerID:  4,
			IndexerIDs: []int64{},
		})
		assert.Contains(t, body, `"indexerIds":[]`)
		assert.NotContains(t, body, `"indexerIds":[4]`)
	})

	t.Run("v3 keeps indexerId", func(t *testing.T) {
		t.Parallel()

		body := releaseProfileBody(t, sonarr.New, &sonarr.ReleaseProfile{Name: "one", IndexerID: 4, IndexerIDs: []int64{4, 9}})
		assert.Contains(t, body, `"indexerId":4`)
		assert.NotContains(t, body, "indexerIds")
	})
}

func releaseProfileBody(t *testing.T, newClient func(*starr.Config) *sonarr.Sonarr, profile *sonarr.ReleaseProfile) string {
	t.Helper()

	var body string

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		raw, err := io.ReadAll(req.Body)
		assert.NoError(t, err)

		body = string(raw)

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"name":"ok","indexerId":0}`))
	}))
	t.Cleanup(server.Close)

	_, err := newClient(starr.New("k", server.URL, 0)).AddReleaseProfile(profile)
	require.NoError(t, err)

	return body
}

func TestReleaseUnmarshal(t *testing.T) {
	t.Parallel()

	var nested sonarr.Release

	err := json.Unmarshal([]byte(`{
		"release":{"guid":"g1","title":"Show.S01E01","indexerId":2,"size":10},
		"parsedInfo":{"seriesTitle":"Show","seasonNumber":1,"isPossibleSpecialEpisode":true,"quality":{"quality":{"id":1,"name":"HDTV"}}},
		"decision":{"approved":true,"rejections":[{"message":"nope"}]}
	}`), &nested)
	require.NoError(t, err)
	assert.Equal(t, "g1", nested.GUID)
	assert.Equal(t, "Show.S01E01", nested.Title)
	assert.True(t, nested.Approved)
	assert.Equal(t, []string{"nope"}, nested.Rejections)
	assert.Equal(t, "Show", nested.SeriesTitle)
	assert.Equal(t, int64(2), nested.IndexerID)
	assert.True(t, nested.IsPossibleSpecialEpisode)

	var flat sonarr.Release

	err = json.Unmarshal([]byte(`{"guid":"flat","title":"Old","approved":true,"indexerId":4,"isPossibleSpecialEpisode":true}`), &flat)
	require.NoError(t, err)
	assert.Equal(t, "flat", flat.GUID)
	assert.Equal(t, "Old", flat.Title)
	assert.True(t, flat.Approved)
	assert.Equal(t, int64(4), flat.IndexerID)
	assert.True(t, flat.IsPossibleSpecialEpisode)
}

func TestReleaseProfileUnmarshalDropsZeroIndexer(t *testing.T) {
	t.Parallel()

	var profile sonarr.ReleaseProfile

	err := json.Unmarshal([]byte(`{"indexerId":0,"indexerIds":[0,6]}`), &profile)
	require.NoError(t, err)
	assert.Equal(t, []int64{6}, profile.IndexerIDs)

	err = json.Unmarshal([]byte(`{"indexerId":0}`), &profile)
	require.NoError(t, err)
	assert.Empty(t, profile.IndexerIDs)

	err = json.Unmarshal([]byte(`{"indexerId":3}`), &profile)
	require.NoError(t, err)
	assert.Equal(t, []int64{3}, profile.IndexerIDs)

	err = json.Unmarshal([]byte(`{"indexerId":4,"indexerIds":[]}`), &profile)
	require.NoError(t, err)
	assert.NotNil(t, profile.IndexerIDs)
	assert.Empty(t, profile.IndexerIDs)

	body := releaseProfileBody(t, sonarr.NewV5, &profile)
	assert.Contains(t, body, `"indexerIds":[]`)
	assert.NotContains(t, body, `"indexerIds":[4]`)

	err = json.Unmarshal([]byte(`{"name":"legacy","indexerId":4}`), &profile)
	require.NoError(t, err)
	assert.Equal(t, []int64{4}, profile.IndexerIDs)

	profile.IndexerIDs = []int64{}
	body = releaseProfileBody(t, sonarr.NewV5, &profile)
	assert.Contains(t, body, `"indexerIds":[]`)
	assert.NotContains(t, body, `"indexerIds":[4]`)
}

func TestForServer(t *testing.T) {
	t.Parallel()

	t.Run("v5", func(t *testing.T) {
		t.Parallel()

		client := forServerClient(t, http.StatusOK, `{"version":"5.0.0.2569"}`)
		selected, err := client.ForServer(context.Background())
		require.NoError(t, err)
		assert.Equal(t, sonarr.APIv5, selected.APIVersion())
	})

	t.Run("v4", func(t *testing.T) {
		t.Parallel()

		client := forServerClient(t, http.StatusOK, `{"version":"4.0.20.3012"}`)
		selected, err := client.ForServer(context.Background())
		require.NoError(t, err)
		assert.Equal(t, sonarr.APIver, selected.APIVersion())
	})

	t.Run("probe failure stays on v3", func(t *testing.T) {
		t.Parallel()

		client := forServerClient(t, http.StatusInternalServerError, `{"message":"nope"}`)
		selected, err := client.ForServer(context.Background())
		require.Error(t, err)
		assert.Equal(t, sonarr.APIver, selected.APIVersion())
	})
}

func forServerClient(t *testing.T, status int, body string) *sonarr.Sonarr {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/api/v3/system/status", req.URL.Path)

		writer.WriteHeader(status)
		_, _ = writer.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	return sonarr.NewV5(starr.New("k", server.URL, 0))
}

func TestV5LanguageProfilesUnsupported(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		t.Fatal("language profiles must not call Sonarr 5")
	}))
	t.Cleanup(server.Close)

	client := sonarr.NewV5(starr.New("k", server.URL, 0))
	_, err := client.GetLanguageProfiles()
	require.ErrorIs(t, err, sonarr.ErrUnsupported)

	_, err = client.GetLanguageProfile(1)
	require.ErrorIs(t, err, sonarr.ErrUnsupported)

	_, err = client.AddLanguageProfile(&sonarr.LanguageProfile{Name: "English"})
	require.ErrorIs(t, err, sonarr.ErrUnsupported)

	_, err = client.UpdateLanguageProfile(&sonarr.LanguageProfile{ID: 1, Name: "English"})
	require.ErrorIs(t, err, sonarr.ErrUnsupported)

	err = client.DeleteLanguageProfile(1)
	require.ErrorIs(t, err, sonarr.ErrUnsupported)
	require.NotErrorIs(t, err, starr.ErrRequestError)
}

func TestExclusionsPaging(t *testing.T) {
	t.Parallel()

	t.Run("v3 paged", func(t *testing.T) {
		t.Parallel()

		records := exclusionPages(t, sonarr.New, "/api/v3/importlistexclusion/paged")
		require.Len(t, records, 2)
		assert.Equal(t, "one", records[0].Title)
		assert.Equal(t, "two", records[1].Title)
	})

	t.Run("v5 collection", func(t *testing.T) {
		t.Parallel()

		records := exclusionPages(t, sonarr.NewV5, "/api/v5/importlistexclusion")
		require.Len(t, records, 2)
	})
}

func exclusionPages(
	t *testing.T,
	newClient func(*starr.Config) *sonarr.Sonarr,
	wantPath string,
) []*sonarr.Exclusion {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		assert.Equal(t, wantPath, req.URL.Path)
		assert.Equal(t, "1000", req.URL.Query().Get("pageSize"))

		writer.WriteHeader(http.StatusOK)

		if req.URL.Query().Get("page") == "1" {
			_, _ = writer.Write([]byte(`{"page":1,"totalRecords":2,"records":[{"id":1,"title":"one","tvdbId":10}]}`))

			return
		}

		_, _ = writer.Write([]byte(`{"page":2,"totalRecords":2,"records":[{"id":2,"title":"two","tvdbId":11}]}`))
	}))
	t.Cleanup(server.Close)

	records, err := newClient(starr.New("k", server.URL, 0)).GetExclusions()
	require.NoError(t, err)

	return records
}
