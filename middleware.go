package main

import (
	"log"
	"net/http"
)

func (app *App) sessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("handling middleware request\n")
		var cookies string
		for _, cookie := range request.Cookies() {
			cookies += cookie.String() + ";"
		}
		session, _, err := app.ory.V0alpha2Api.ToSession(request.Context()).Cookie(cookies).Execute()
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			http.Redirect(writer, request, "/.ory/api/kratos/public/self-service/login/browser", http.StatusSeeOther)
		}
		app.cookies = cookies
		app.session = session
		next.ServeHTTP(writer, request)
		return
	}
}
