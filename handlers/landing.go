package handlers

import (
	"context"
	"net/http"

	myClerk "github.com/bourgeoisie-awards-2024/clerk"
	"github.com/bourgeoisie-awards-2024/pages"
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

type LandingPageHandler struct {
	clerkFrontendConfig *myClerk.FrontendConfig
}

func (h LandingPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, ok := clerk.SessionFromContext(r.Context())
	if ok {
		http.Redirect(w, r, "/welcome", http.StatusTemporaryRedirect)
		return
	}
	component := pages.LandingPage(h.clerkFrontendConfig)
	component.Render(context.Background(), w)
}

func NewLandingPageHandler(clerkFrontendConfig *myClerk.FrontendConfig) http.Handler {
	return LandingPageHandler{
		clerkFrontendConfig: clerkFrontendConfig,
	}
}
