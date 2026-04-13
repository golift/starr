package starrconnect

import (
	"fmt"
	"net/http"
)

// LidarrHandler dispatches Lidarr webhook POSTs to per-event callbacks.
type LidarrHandler struct {
	OnGrab              func(*LidarrGrab) error
	OnDownload          func(*LidarrDownload) error
	OnDownloadFailure   func(*LidarrDownloadFailure) error
	OnImportFailure     func(*LidarrImportFailure) error
	OnRename            func(*LidarrRename) error
	OnRetag             func(*LidarrRetag) error
	OnArtistAdd         func(*ArtistAdd) error
	OnArtistDelete      func(*ArtistDelete) error
	OnAlbumDelete       func(*AlbumDelete) error
	OnHealth            func(*LidarrHealth) error
	OnHealthRestored    func(*LidarrHealth) error
	OnApplicationUpdate func(*LidarrApplicationUpdate) error
	OnTest              func(*LidarrGrab) error
	OnError             func(error)
}

// ServeHTTP implements http.Handler for Lidarr webhooks.
func (h *LidarrHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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

func (h *LidarrHandler) handleWebhook(req *http.Request) error {
	body, err := readRequestBody(req)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "bad request", err)
	}

	event, err := ParseLidarr(body)
	if err != nil {
		return webhookErr(http.StatusBadRequest, "invalid json", err)
	}

	return h.dispatchEvent(event)
}

func (h *LidarrHandler) dispatchEvent(event *LidarrEvent) error {
	switch event.EventType {
	case EventTest:
		return runWebhookCallback(h.OnTest, event.GetGrab)
	case EventGrab:
		return runWebhookCallback(h.OnGrab, event.GetGrab)
	case EventDownload:
		return runWebhookCallback(h.OnDownload, event.GetDownload)
	case EventDownloadFailure:
		return runWebhookCallback(h.OnDownloadFailure, event.GetDownloadFailure)
	case EventImportFailure:
		return runWebhookCallback(h.OnImportFailure, event.GetImportFailure)
	case EventRename:
		return runWebhookCallback(h.OnRename, event.GetRename)
	case EventRetag:
		return runWebhookCallback(h.OnRetag, event.GetRetag)
	case EventArtistAdd:
		return runWebhookCallback(h.OnArtistAdd, event.GetArtistAdd)
	case EventArtistDelete:
		return runWebhookCallback(h.OnArtistDelete, event.GetArtistDelete)
	case EventAlbumDelete:
		return runWebhookCallback(h.OnAlbumDelete, event.GetAlbumDelete)
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
