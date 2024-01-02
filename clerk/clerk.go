package clerk

var SessionCookie = "__session"

type FrontendConfig struct {
	PublishableKey string `json:"publishableKey"`
	FrontendAPI    string `json:"frontendAPI"`
	Version        string `json:"version"`
}
