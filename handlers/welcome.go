package handlers

import (
	"context"
	"net/http"

	myClerk "github.com/bourgeoisie-awards-2024/clerk"
	"github.com/bourgeoisie-awards-2024/pages"
)

type WelcomePageHandler struct {
	clerkFrontendConfig *myClerk.FrontendConfig
}

func (h WelcomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	component := pages.WelcomePage(h.clerkFrontendConfig)
	component.Render(context.Background(), w)
}

func NewWelcomePageHandler(clerkFrontendConfig *myClerk.FrontendConfig) http.Handler {
	return WelcomePageHandler{
		clerkFrontendConfig: clerkFrontendConfig,
	}
}
