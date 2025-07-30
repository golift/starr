//nolint:lll
package starrcmd

/*
All events accounted for; 1/30/2022.
https://github.com/Sonarr/Sonarr/blob/develop/src/NzbDrone.Core/Notifications/CustomScript/CustomScript.cs
*/

import (
	"time"
)

// SonarrApplicationUpdate is the ApplicationUpdate event.
type SonarrApplicationUpdate struct {
	PreviousVersion string `env:"sonarr_update_previousversion"` // 4.0.3.5875
	NewVersion      string `env:"sonarr_update_newversion"`      // 4.0.4.5909
	Message         string `env:"sonarr_update_message"`         // Sonarr updated from 4.0.3.5875 to 4.0.4.5909
}

// SonarrHealthIssue is the HealthIssue event.
type SonarrHealthIssue struct {
	Message   string `env:"sonarr_health_issue_message"` // Lists unavailable due to failures: Listnamehere
	IssueType string `env:"sonarr_health_issue_type"`    // ImportListStatusCheck
	Wiki      string `env:"sonarr_health_issue_wiki"`    // https://wiki.servarr.com/
	Level     string `env:"sonarr_health_issue_level"`   // Warning
}

// SonarrGrab is the Grab event.
type SonarrGrab struct {
	Quality            string      `env:"sonarr_release_quality"`                  // HDTV-720p
	Title              string      `env:"sonarr_series_title"`                     // This Is Us
	DownloadClient     string      `env:"sonarr_download_client"`                  // NZBGet
	ReleaseTitle       string      `env:"sonarr_release_title"`                    // This.is.Us.S06E04.720p.HDTV.x264-SYNCOPY
	DownloadID         string      `env:"sonarr_download_id"`                      // a87bda3c0e7f40a1b8fa011b421a5201
	ReleaseIndexer     string      `env:"sonarr_release_indexer"`                  // Indexor (Prowlarr)
	SeriesType         string      `env:"sonarr_series_type"`                      // Standard
	ReleaseGroup       string      `env:"sonarr_release_releasegroup"`             // SYNCOPY
	IMDbID             string      `env:"sonarr_series_imdbid"`                    // tt5555260
	EpisodeNumbers     []int       `env:"sonarr_release_episodenumbers,,"`         // 4
	EpisodeAirDates    []string    `env:"sonarr_release_episodeairdates,,"`        // 2022-01-25
	EpisodeTitles      []string    `env:"sonarr_release_episodetitles,|"`          // Don't Let Me Keep You
	AbsEpisodeNumbers  []int       `env:"sonarr_release_absoluteepisodenumbers,,"` // 92
	EpisodeAirDatesUTC []time.Time `env:"sonarr_release_episodeairdatesutc,,"`     // 1/26/2022 2:00:00 AM
	QualityVersion     int64       `env:"sonarr_release_qualityversion"`           // 1
	SeriesID           int64       `env:"sonarr_series_id"`                        // 47
	EpisodeCount       int         `env:"sonarr_release_episodecount"`             // 1
	Size               int64       `env:"sonarr_release_size"`                     // 885369406
	TVDbID             int64       `env:"sonarr_series_tvdbid"`                    // 311714
	TVMazeID           int64       `env:"sonarr_series_tvmazeid"`                  // 17128
	SeasonNumber       int         `env:"sonarr_release_seasonnumber"`             // 6
}

// SonarrDownload is the Download event.
type SonarrDownload struct {
	Title                string      `env:"sonarr_series_title"`                     // Puppy Dog Pals
	SourceFolder         string      `env:"sonarr_episodefile_sourcefolder"`         // /downloads/completed/Series/Puppy.Dog.Pals.S05E03e04.The.Puppy.Outdoor.Play.Day.Games.for.the.Glove.of.the.Game.HULU.WEB-DL.AAC2.0.H.264-LAZY
	Quality              string      `env:"sonarr_episodefile_quality"`              // WEBDL-480p
	ReleaseGroup         string      `env:"sonarr_episodefile_releasegroup"`         // LAZY
	DownloadClient       string      `env:"sonarr_download_client"`                  // NZBGET
	EpisodePath          string      `env:"sonarr_episodefile_path"`                 // /tv/Puppy Dog Pals/Season 5/Puppy Dog Pals - S05E03-04 - The Puppy Outdoor Play Day Games + For the Glove of the Game WEBDL-480p.mkv
	SceneName            string      `env:"sonarr_episodefile_scenename"`            // Puppy.Dog.Pals.S05E03e04.The.Puppy.Outdoor.Play.Day.Games.for.the.Glove.of.the.Game.HULU.WEB-DL.AAC2.0.H.264-LAZY
	Path                 string      `env:"sonarr_series_path"`                      // /tv/Puppy Dog Pals
	SourcePath           string      `env:"sonarr_episodefile_sourcepath"`           // /downloads/completed/Series/Puppy.Dog.Pals.S05E03e04.The.Puppy.Outdoor.Play.Day.Games.for.the.Glove.of.the.Game.HULU.WEB-DL.AAC2.0.H.264-LAZY/9ZMAepAkHwQsOn.mkv
	DownloadID           string      `env:"sonarr_download_id"`                      // 977d4bd4ac3845c0a2d5c890cc5a10e4
	SeriesType           string      `env:"sonarr_series_type"`                      // Standard
	IMDbID               string      `env:"sonarr_series_imdbid"`                    // tt6688750
	RelativePath         string      `env:"sonarr_episodefile_relativepath"`         // Season 5/Puppy Dog Pals - S05E03-04 - The Puppy Outdoor Play Day Games + For the Glove of the Game WEBDL-480p.mkv
	EpisodeIDs           []int64     `env:"sonarr_episodefile_episodeids,,"`         // 22691,22692
	EpisodeNumbers       []int       `env:"sonarr_episodefile_episodenumbers,,"`     // 3,4
	EpisodeAirDates      []string    `env:"sonarr_episodefile_episodeairdates,,"`    // 2022-01-21,2022-01-21
	EpisodeTitles        []string    `env:"sonarr_episodefile_episodetitles,|"`      // The Puppy Outdoor Play Day Games|For the Glove of the Game
	EpisodeAirDatesUTC   []time.Time `env:"sonarr_episodefile_episodeairdatesutc,,"` // 1/21/2022 2:00:00 PM,1/21/2022 2:12:00 PM
	DeletedRelativePaths []string    `env:"sonarr_deletedrelativepaths,|"`           // Not always present.
	DeletedPaths         []string    `env:"sonarr_deletedpaths,|"`                   // Not always present.
	SeriesID             int64       `env:"sonarr_series_id"`                        // 108
	QualityVersion       int64       `env:"sonarr_episodefile_qualityversion"`       // 1
	FileID               int64       `env:"sonarr_episodefile_id"`                   // 14996
	TVDbID               int64       `env:"sonarr_series_tvdbid"`                    // 325978
	TVMazeID             int64       `env:"sonarr_series_tvmazeid"`                  // 26341
	EpisodeCount         int         `env:"sonarr_episodefile_episodecount"`         // 2
	SeasonNumber         int         `env:"sonarr_episodefile_seasonnumber"`         // 5
	IsUpgrade            bool        `env:"sonarr_isupgrade"`                        // False
}

// SonarrRename is the Rename event.
type SonarrRename struct {
	Title                 string   `env:"sonarr_series_title"`                        // series.Title)
	Path                  string   `env:"sonarr_series_path"`                         // series.Path)
	IMDbID                string   `env:"sonarr_series_imdbid"`                       // series.ImdbId ?? string.Empty)
	SeriesType            string   `env:"sonarr_series_type"`                         // series.SeriesType.ToString())
	FileIDs               []int64  `env:"sonarr_episodefile_ids,,"`                   // string.Join(",", renamedFiles.Select(e => e.EpisodeFile.Id)))
	RelativePaths         []string `env:"sonarr_episodefile_relativepaths,|"`         // string.Join("|", renamedFiles.Select(e => e.EpisodeFile.RelativePath)))
	Paths                 []string `env:"sonarr_episodefile_paths,|"`                 // string.Join("|", renamedFiles.Select(e => e.EpisodeFile.Path)))
	PreviousRelativePaths []string `env:"sonarr_episodefile_previousrelativepaths,|"` // string.Join("|", renamedFiles.Select(e => e.PreviousRelativePath)))
	PreviousPaths         []string `env:"sonarr_episodefile_previouspaths,|"`         // string.Join("|", renamedFiles.Select(e => e.PreviousPath)))
	ID                    int64    `env:"sonarr_series_id"`                           // series.Id.ToString())
	TVDbID                int64    `env:"sonarr_series_tvdbid"`                       // series.TvdbId.ToString())
	TVMazeID              int64    `env:"sonarr_series_tvmazeid"`                     // series.TvMazeId.ToString())
}

// SonarrSeriesDelete is the SeriesDelete event.
type SonarrSeriesDelete struct {
	Title        string `env:"sonarr_series_title"`        // series.Title)
	Path         string `env:"sonarr_series_path"`         // series.Path)
	IMDbID       string `env:"sonarr_series_imdbid"`       // series.ImdbId ?? string.Empty)
	SeriesType   string `env:"sonarr_series_type"`         // series.SeriesType.ToString())
	DeletedFiles string `env:"sonarr_series_deletedfiles"` // deleteMessage.DeletedFiles.ToString())
	ID           int64  `env:"sonarr_series_id"`           // series.Id.ToString())
	TVDbID       int64  `env:"sonarr_series_tvdbid"`       // series.TvdbId.ToString())
	TVMazeID     int64  `env:"sonarr_series_tvmazeid"`     // series.TvMazeId.ToString())
}

// SonarrEpisodeFileDelete is the EpisodeFileDelete event.
type SonarrEpisodeFileDelete struct {
	Reason             string      `env:"sonarr_episodefile_deletereason"`         // deleteMessage.Reason.ToString())
	Title              string      `env:"sonarr_series_title"`                     // series.Title)
	Path               string      `env:"sonarr_series_path"`                      // series.Path)
	IMDbID             string      `env:"sonarr_series_imdbid"`                    // series.ImdbId ?? string.Empty)
	SeriesType         string      `env:"sonarr_series_type"`                      // series.SeriesType.ToString())
	RelativePath       string      `env:"sonarr_episodefile_relativepath"`         // episodeFile.RelativePath)
	FilePath           string      `env:"sonarr_episodefile_path"`                 // Path.Combine(series.Path, episodeFile.RelativePath))
	SeasonNumber       string      `env:"sonarr_episodefile_seasonnumber"`         // episodeFile.SeasonNumber.ToString())
	Quality            string      `env:"sonarr_episodefile_quality"`              // episodeFile.Quality.Quality.Name)
	QualityVersion     string      `env:"sonarr_episodefile_qualityversion"`       // episodeFile.Quality.Revision.Version.ToString())
	ReleaseGroup       string      `env:"sonarr_episodefile_releasegroup"`         // episodeFile.ReleaseGroup ?? string.Empty)
	SceneName          string      `env:"sonarr_episodefile_scenename"`            // episodeFile.SceneName ?? string.Empty)
	EpisodeIDs         []int64     `env:"sonarr_episodefile_episodeids,,"`         // string.Join(",", episodeFile.Episodes.Value.Select(e => e.Id)))
	EpisodeNumbers     []int       `env:"sonarr_episodefile_episodenumbers,,"`     // string.Join(",", episodeFile.Episodes.Value.Select(e => e.EpisodeNumber)))
	EpisodeAirDates    []string    `env:"sonarr_episodefile_episodeairdates,,"`    // string.Join(",", episodeFile.Episodes.Value.Select(e => e.AirDate)))
	EpisodeAirDatesUTC []time.Time `env:"sonarr_episodefile_episodeairdatesutc,,"` // string.Join(",", episodeFile.Episodes.Value.Select(e => e.AirDateUtc)))
	EpisodeTitles      []string    `env:"sonarr_episodefile_episodetitles,|"`      // string.Join("|", episodeFile.Episodes.Value.Select(e => e.Title)))
	ID                 int64       `env:"sonarr_series_id"`                        // series.Id.ToString())
	TVDbID             int64       `env:"sonarr_series_tvdbid"`                    // series.TvdbId.ToString())
	TVMazeID           int64       `env:"sonarr_series_tvmazeid"`                  // series.TvMazeId.ToString())
	FileID             int64       `env:"sonarr_episodefile_id"`                   // episodeFile.Id.ToString())
	EpisodeCount       int         `env:"sonarr_episodefile_episodecount"`         // episodeFile.Episodes.Value.Count.ToString())
}

// SonarrTest has no members.
type SonarrTest struct{}

// GetSonarrApplicationUpdate returns the ApplicationUpdate event data.
func (c *CmdEvent) GetSonarrApplicationUpdate() (output SonarrApplicationUpdate, err error) {
	return output, c.get(EventApplicationUpdate, &output)
}

// GetSonarrHealthIssue returns the ApplicationUpdate event data.
func (c *CmdEvent) GetSonarrHealthIssue() (output SonarrHealthIssue, err error) {
	return output, c.get(EventHealthIssue, &output)
}

// GetSonarrTest returns the ApplicationUpdate event data.
func (c *CmdEvent) GetSonarrTest() (output SonarrTest, err error) {
	return output, c.get(EventTest, &output)
}

// GetSonarrGrab returns the Grab event data.
func (c *CmdEvent) GetSonarrGrab() (output SonarrGrab, err error) {
	return output, c.get(EventGrab, &output)
}

// GetSonarrDownload returns the Download event data.
func (c *CmdEvent) GetSonarrDownload() (output SonarrDownload, err error) {
	return output, c.get(EventDownload, &output)
}

// GetSonarrRename returns the Rename event data.
func (c *CmdEvent) GetSonarrRename() (output SonarrRename, err error) {
	return output, c.get(EventRename, &output)
}

// GetSonarrSeriesDelete returns the SeriesDelete event data.
func (c *CmdEvent) GetSonarrSeriesDelete() (output SonarrSeriesDelete, err error) {
	return output, c.get(EventSeriesDelete, &output)
}

// GetSonarrEpisodeFileDelete returns the EpisodeFileDelete event data.
func (c *CmdEvent) GetSonarrEpisodeFileDelete() (output SonarrEpisodeFileDelete, err error) {
	return output, c.get(EventEpisodeFileDelete, &output)
}
