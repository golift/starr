package starrconnect

import (
	"fmt"
	"net/http"
)

// ProwlarrHandler dispatches Prowlarr webhook POSTs to per-event callbacks.
type ProwlarrHandler struct {
	OnGrab              func(*ProwlarrGrab) error
	OnTest              func(*ProwlarrTest) error
	OnDownload          func(*ProwlarrDownload) error
	OnRename            func(*ProwlarrRename) error
	OnHealth            func(*ProwlarrHealth) error
	OnHealthRestored    func(*ProwlarrHealth) error
	OnApplicationUpdate func(*ProwlarrApplicationUpdate) error
	OnError             func(error)
}

// ServeHTTP implements http.Handler for Prowlarr webhooks.
func (h *ProwlarrHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

func (h *ProwlarrHandler) handleWebhook(req *http.Request) error {
	body, err := readRequestBody(req)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "bad request", err)
	}

	event, err := ParseProwlarr(body)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "invalid json", err)
	}

	return h.dispatchEvent(event)
}

func (h *ProwlarrHandler) dispatchEvent(event *ProwlarrEvent) error {
	switch event.EventType {
	case EventTest:
		return runWebhookCallback(h.OnTest, event.GetTest)
	case EventGrab:
		return runWebhookCallback(h.OnGrab, event.GetGrab)
	case EventDownload:
		return runWebhookCallback(h.OnDownload, event.GetDownload)
	case EventRename:
		return runWebhookCallback(h.OnRename, event.GetRename)
	case EventHealth:
		return runWebhookCallback(h.OnHealth, event.GetHealth)
	case EventHealthRestored:
		return runWebhookCallback(h.OnHealthRestored, event.GetHealthRestored)
	case EventApplicationUpdate:
		return runWebhookCallback(h.OnApplicationUpdate, event.GetApplicationUpdate)
	default:
		return webhookErr(http.StatusInternalServerError, "handler error",
			fmt.Errorf("%w: %q", ErrUnknownEvent, event.EventType))
	}
}
