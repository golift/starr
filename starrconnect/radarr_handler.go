package starrconnect

import (
	"fmt"
	"net/http"
)

// RadarrHandler dispatches Radarr webhook POSTs to per-event callbacks.
type RadarrHandler struct {
	OnGrab                      func(*RadarrGrab) error
	OnDownload                  func(*RadarrDownload) error
	OnRename                    func(*RadarrRename) error
	OnMovieAdded                func(*MovieAdded) error
	OnMovieDelete               func(*MovieDelete) error
	OnMovieFileDelete           func(*MovieFileDelete) error
	OnHealth                    func(*RadarrHealth) error
	OnHealthRestored            func(*RadarrHealth) error
	OnApplicationUpdate         func(*RadarrApplicationUpdate) error
	OnManualInteractionRequired func(*RadarrManualInteraction) error
	OnTest                      func(*RadarrGrab) error
	OnError                     func(error)
}

// ServeHTTP implements http.Handler for Radarr webhooks.
func (h *RadarrHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

func (h *RadarrHandler) handleWebhook(req *http.Request) error {
	body, err := readRequestBody(req)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "bad request", err)
	}

	event, err := ParseRadarr(body)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "invalid json", err)
	}

	return h.dispatchEvent(event)
}

func (h *RadarrHandler) dispatchEvent(event *RadarrEvent) error {
	switch event.EventType {
	case EventTest:
		return runWebhookCallback(h.OnTest, event.GetGrab)
	case EventGrab:
		return runWebhookCallback(h.OnGrab, event.GetGrab)
	case EventDownload:
		return runWebhookCallback(h.OnDownload, event.GetDownload)
	case EventRename:
		return runWebhookCallback(h.OnRename, event.GetRename)
	case EventMovieAdded:
		return runWebhookCallback(h.OnMovieAdded, event.GetMovieAdded)
	case EventMovieDelete:
		return runWebhookCallback(h.OnMovieDelete, event.GetMovieDelete)
	case EventMovieFileDelete:
		return runWebhookCallback(h.OnMovieFileDelete, event.GetMovieFileDelete)
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
