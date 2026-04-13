package starrshared

// Health is the /health API resource shared by all Starr app clients.
type Health struct {
	ID      int    `json:"id"`
	Source  string `json:"source,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	WikiURL string `json:"wikiUrl,omitempty"`
}
