package starrshared

// DiskSpace is the /diskspace API resource shared by all Starr app clients.
type DiskSpace struct {
	ID         int    `json:"id"`
	Path       string `json:"path,omitempty"`
	Label      string `json:"label,omitempty"`
	FreeSpace  int64  `json:"freeSpace"`
	TotalSpace int64  `json:"totalSpace"`
}
