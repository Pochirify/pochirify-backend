package v1

import (
	"net/http"

	"github.com/Pochirify/pochirify-backend/internal/usecase"
)

type WebhookHandler struct {
	WebhookApp usecase.App
}

func NewWebhookHandler(app usecase.App) *WebhookHandler {
	return &WebhookHandler{
		WebhookApp: app,
	}
}

func (h WebhookHandler) PayPayTransactionEventHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.WebhookApp.PayPayTransactionEvent(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
