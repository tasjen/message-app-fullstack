package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /login/{provider}", app.oauthLoginHandler)
	mux.HandleFunc("GET /callback/{provider}", app.oauthCallbackHandler)
	return mux
}
