package handlers

import (
	"context"
	"net/http"

	myClerk "github.com/bourgeoisie-awards-2024/clerk"
	"github.com/bourgeoisie-awards-2024/pages"
)

type ThankYouPageHandler struct {
	clerkFrontendConfig *myClerk.FrontendConfig
}

func (h ThankYouPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	component := pages.ThankYouPage(h.clerkFrontendConfig)
	component.Render(context.Background(), w)
}

func NewThankYouPageHandler(clerkFrontendConfig *myClerk.FrontendConfig) http.Handler {
	return ThankYouPageHandler{
		clerkFrontendConfig: clerkFrontendConfig,
	}
}
