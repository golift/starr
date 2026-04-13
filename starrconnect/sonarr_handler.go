package starrconnect

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SonarrHandler dispatches Sonarr webhook POSTs to per-event callbacks.
type SonarrHandler struct {
	OnGrab                      func(*SonarrGrab) error
	OnDownload                  func(*SonarrDownload) error
	OnImportComplete            func(*SonarrImportComplete) error
	OnRename                    func(*SonarrRename) error
	OnSeriesAdd                 func(*SeriesAdd) error
	OnSeriesDelete              func(*SeriesDelete) error
	OnEpisodeFileDelete         func(*EpisodeFileDelete) error
	OnHealth                    func(*SonarrHealth) error
	OnHealthRestored            func(*SonarrHealth) error
	OnApplicationUpdate         func(*SonarrApplicationUpdate) error
	OnManualInteractionRequired func(*SonarrManualInteraction) error
	OnTest                      func(*SonarrGrab) error
	OnError                     func(error)
}

// ServeHTTP implements http.Handler for Sonarr webhooks.
func (h *SonarrHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resp, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.handleWebhook(req); err != nil {
		writeWebhookHTTPError(resp, h.OnError, err)
		return
	}

	resp.WriteHeader(http.StatusOK)
}

func (h *SonarrHandler) handleWebhook(req *http.Request) error {
	body, err := readRequestBody(req)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "bad request", err)
	}

	event, err := ParseSonarr(body)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "invalid json", err)
	}

	return h.dispatchEvent(event, body)
}

func (h *SonarrHandler) dispatchEvent(event *SonarrEvent, body []byte) error {
	switch event.EventType {
	case EventTest:
		return runWebhookCallback(h.OnTest, event.GetGrab)
	case EventGrab:
		return runWebhookCallback(h.OnGrab, event.GetGrab)
	case EventDownload:
		return h.dispatchSonarrDownload(event, body)
	case EventRename:
		return runWebhookCallback(h.OnRename, event.GetRename)
	case EventSeriesAdd:
		return runWebhookCallback(h.OnSeriesAdd, event.GetSeriesAdd)
	case EventSeriesDelete:
		return runWebhookCallback(h.OnSeriesDelete, event.GetSeriesDelete)
	case EventEpisodeFileDelete:
		return runWebhookCallback(h.OnEpisodeFileDelete, event.GetEpisodeFileDelete)
	case EventHealth:
		return runWebhookCallback(h.OnHealth, event.GetHealth)
	case EventHealthRestored:
		return runWebhookCallback(h.OnHealthRestored, event.GetHealthRestored)
	case EventApplicationUpdate:
		return runWebhookCallback(h.OnApplicationUpdate, event.GetApplicationUpdate)
	case EventManualInteractionRequired:
		return runWebhookCallback(h.OnManualInteractionRequired, event.GetManualInteraction)
	default:
		return webhookErr(http.StatusInternalServerError, "handler error",
			fmt.Errorf("%w: %q", ErrUnknownEvent, event.EventType))
	}
}

func (h *SonarrHandler) dispatchSonarrDownload(event *SonarrEvent, body []byte) error {
	if sonarrIsImportCompleteBody(body) {
		return runWebhookCallback(h.OnImportComplete, event.GetImportComplete)
	}

	return runWebhookCallback(h.OnDownload, event.GetDownload)
}

// sonarrIsImportCompleteBody reports whether the JSON is the batch import-complete shape (episodeFiles).
func sonarrIsImportCompleteBody(body []byte) bool {
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(body, &keys); err != nil {
		return false
	}

	raw, exists := keys["episodeFiles"]

	return exists && len(raw) > 0 && string(raw) != "null" && string(raw) != "[]"
}
