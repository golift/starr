package starrcmd

/*
Prowlarr only has 3 events, all accounted for; 1/30/2022.
https://github.com/Prowlarr/Prowlarr/blob/develop/src/NzbDrone.Core/Notifications/CustomScript/CustomScript.cs
*/

// ProwlarrApplicationUpdate is the ApplicationUpdate event.
type ProwlarrApplicationUpdate struct {
	PreviousVersion string `env:"prowlarr_update_previousversion"` // 4.0.3.5875
	NewVersion      string `env:"prowlarr_update_newversion"`      // 4.0.4.5909
	Message         string `env:"prowlarr_update_message"`         // Prowlarr updated from 4.0.3.5875 to 4.0.4.5909
}

// ProwlarrHealthIssue is the HealthIssue event.
type ProwlarrHealthIssue struct {
	Message   string `env:"prowlarr_health_issue_message"` // some message about sme problem
	IssueType string `env:"prowlarr_health_issue_type"`    // NeverSeenOne
	Wiki      string `env:"prowlarr_health_issue_wiki"`    // something something something
	Level     string `env:"prowlarr_health_issue_level"`   // Warning
}

// ProwlarrTest has no members.
type ProwlarrTest struct{}

// GetProwlarrApplicationUpdate returns the ApplicationUpdate event data.
func GetProwlarrApplicationUpdate() (output ProwlarrApplicationUpdate, err error) {
	return output, get(EventApplicationUpdate, &output)
}

// GetProwlarrHealthIssue returns the ApplicationUpdate event data.
func GetProwlarrHealthIssue() (output ProwlarrHealthIssue, err error) {
	return output, get(EventHealthIssue, &output)
}

// GetProwlarrTest returns the ApplicationUpdate event data.
func GetProwlarrTest() (output ProwlarrTest, err error) {
	return output, get(EventTest, &output)
}
