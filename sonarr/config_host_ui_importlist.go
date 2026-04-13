package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const (
	bpConfigHost       = APIver + "/config/host"
	bpConfigUI         = APIver + "/config/ui"
	bpConfigImportList = APIver + "/config/importlist"
)

// HostConfig is the /api/v3/config/host resource.
type HostConfig struct {
	ID                        int    `json:"id,omitempty"`
	BindAddress               string `json:"bindAddress,omitempty"`
	Port                      int    `json:"port"`
	SSLPort                   int    `json:"sslPort"`
	EnableSSL                 bool   `json:"enableSsl"`
	LaunchBrowser             bool   `json:"launchBrowser"`
	AuthenticationMethod      string `json:"authenticationMethod,omitempty"`
	AuthenticationRequired    string `json:"authenticationRequired,omitempty"`
	AnalyticsEnabled          bool   `json:"analyticsEnabled"`
	Username                  string `json:"username,omitempty"`
	Password                  string `json:"password,omitempty"`
	PasswordConfirmation      string `json:"passwordConfirmation,omitempty"`
	LogLevel                  string `json:"logLevel,omitempty"`
	LogSizeLimit              int    `json:"logSizeLimit"`
	ConsoleLogLevel           string `json:"consoleLogLevel,omitempty"`
	Branch                    string `json:"branch,omitempty"`
	APIKey                    string `json:"apiKey,omitempty"`
	SSLCertPath               string `json:"sslCertPath,omitempty"`
	SSLCertPassword           string `json:"sslCertPassword,omitempty"`
	URLBase                   string `json:"urlBase,omitempty"`
	InstanceName              string `json:"instanceName,omitempty"`
	ApplicationURL            string `json:"applicationUrl,omitempty"`
	UpdateAutomatically       bool   `json:"updateAutomatically"`
	UpdateMechanism           string `json:"updateMechanism,omitempty"`
	UpdateScriptPath          string `json:"updateScriptPath,omitempty"`
	ProxyEnabled              bool   `json:"proxyEnabled"`
	ProxyType                 string `json:"proxyType,omitempty"`
	ProxyHostname             string `json:"proxyHostname,omitempty"`
	ProxyPort                 int    `json:"proxyPort"`
	ProxyUsername             string `json:"proxyUsername,omitempty"`
	ProxyPassword             string `json:"proxyPassword,omitempty"`
	ProxyBypassFilter         string `json:"proxyBypassFilter,omitempty"`
	ProxyBypassLocalAddresses bool   `json:"proxyBypassLocalAddresses"`
	CertificateValidation     string `json:"certificateValidation,omitempty"`
	BackupFolder              string `json:"backupFolder,omitempty"`
	BackupInterval            int    `json:"backupInterval"`
	BackupRetention           int    `json:"backupRetention"`
	TrustCgnatIPAddresses     bool   `json:"trustCgnatIpAddresses"`
}

// UIConfig is the /api/v3/config/ui resource.
type UIConfig struct {
	ID                       int    `json:"id,omitempty"`
	FirstDayOfWeek           int    `json:"firstDayOfWeek"`
	CalendarWeekColumnHeader string `json:"calendarWeekColumnHeader,omitempty"`
	ShortDateFormat          string `json:"shortDateFormat,omitempty"`
	LongDateFormat           string `json:"longDateFormat,omitempty"`
	TimeFormat               string `json:"timeFormat,omitempty"`
	ShowRelativeDates        bool   `json:"showRelativeDates"`
	EnableColorImpairedMode  bool   `json:"enableColorImpairedMode"`
	Theme                    string `json:"theme,omitempty"`
	UILanguage               int    `json:"uiLanguage"`
}

// ImportListConfig is the /api/v3/config/importlist resource.
type ImportListConfig struct {
	ID            int    `json:"id,omitempty"`
	ListSyncLevel string `json:"listSyncLevel,omitempty"`
	ListSyncTag   int    `json:"listSyncTag"`
}

// GetHostConfig returns the host configuration.
func (s *Sonarr) GetHostConfig() (*HostConfig, error) {
	return s.GetHostConfigContext(context.Background())
}

// GetHostConfigContext returns the host configuration.
func (s *Sonarr) GetHostConfigContext(ctx context.Context) (*HostConfig, error) {
	var output HostConfig

	req := starr.Request{URI: bpConfigHost}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetHostConfigByID returns the host configuration for the given id.
func (s *Sonarr) GetHostConfigByID(id int) (*HostConfig, error) {
	return s.GetHostConfigByIDContext(context.Background(), id)
}

// GetHostConfigByIDContext returns the host configuration for the given id.
func (s *Sonarr) GetHostConfigByIDContext(ctx context.Context, id int) (*HostConfig, error) {
	var output HostConfig

	req := starr.Request{URI: path.Join(bpConfigHost, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateHostConfig updates the host configuration.
func (s *Sonarr) UpdateHostConfig(in *HostConfig) (*HostConfig, error) {
	return s.UpdateHostConfigContext(context.Background(), in)
}

// UpdateHostConfigContext updates the host configuration.
func (s *Sonarr) UpdateHostConfigContext(ctx context.Context, input *HostConfig) (*HostConfig, error) {
	var output HostConfig

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpConfigHost, err)
	}

	req := starr.Request{URI: path.Join(bpConfigHost, starr.Str(input.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// GetUIConfig returns the UI configuration.
func (s *Sonarr) GetUIConfig() (*UIConfig, error) {
	return s.GetUIConfigContext(context.Background())
}

// GetUIConfigContext returns the UI configuration.
func (s *Sonarr) GetUIConfigContext(ctx context.Context) (*UIConfig, error) {
	var output UIConfig

	req := starr.Request{URI: bpConfigUI}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetUIConfigByID returns the UI configuration for the given id.
func (s *Sonarr) GetUIConfigByID(id int) (*UIConfig, error) {
	return s.GetUIConfigByIDContext(context.Background(), id)
}

// GetUIConfigByIDContext returns the UI configuration for the given id.
func (s *Sonarr) GetUIConfigByIDContext(ctx context.Context, id int) (*UIConfig, error) {
	var output UIConfig

	req := starr.Request{URI: path.Join(bpConfigUI, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateUIConfig updates the UI configuration.
func (s *Sonarr) UpdateUIConfig(in *UIConfig) (*UIConfig, error) {
	return s.UpdateUIConfigContext(context.Background(), in)
}

// UpdateUIConfigContext updates the UI configuration.
func (s *Sonarr) UpdateUIConfigContext(ctx context.Context, input *UIConfig) (*UIConfig, error) {
	var output UIConfig

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpConfigUI, err)
	}

	req := starr.Request{URI: path.Join(bpConfigUI, starr.Str(input.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// GetImportListConfig returns the import list global configuration.
func (s *Sonarr) GetImportListConfig() (*ImportListConfig, error) {
	return s.GetImportListConfigContext(context.Background())
}

// GetImportListConfigContext returns the import list global configuration.
func (s *Sonarr) GetImportListConfigContext(ctx context.Context) (*ImportListConfig, error) {
	var output ImportListConfig

	req := starr.Request{URI: bpConfigImportList}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetImportListConfigByID returns the import list global configuration for the given id.
func (s *Sonarr) GetImportListConfigByID(id int) (*ImportListConfig, error) {
	return s.GetImportListConfigByIDContext(context.Background(), id)
}

// GetImportListConfigByIDContext returns the import list global configuration for the given id.
func (s *Sonarr) GetImportListConfigByIDContext(ctx context.Context, id int) (*ImportListConfig, error) {
	var output ImportListConfig

	req := starr.Request{URI: path.Join(bpConfigImportList, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateImportListConfig updates the import list global configuration.
func (s *Sonarr) UpdateImportListConfig(input *ImportListConfig) (*ImportListConfig, error) {
	return s.UpdateImportListConfigContext(context.Background(), input)
}

// UpdateImportListConfigContext updates the import list global configuration.
func (s *Sonarr) UpdateImportListConfigContext(
	ctx context.Context, input *ImportListConfig,
) (*ImportListConfig, error) {
	var output ImportListConfig

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpConfigImportList, err)
	}

	req := starr.Request{URI: path.Join(bpConfigImportList, starr.Str(input.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
