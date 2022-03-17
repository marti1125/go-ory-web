package main

import (
	"fmt"
	ory "github.com/ory/client-go"
	"net/http"
	"os"
)

type App struct {
	ory     *ory.APIClient
	cookies string
	session *ory.Session
}

func main() {
	proxyPort := os.Getenv("PROXY_PORT")
	if proxyPort == "" {
		proxyPort = "4000"
	}
	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%s/.ory", proxyPort)}}

	app := &App{
		ory: ory.NewAPIClient(c),
	}

	mux := http.NewServeMux()
	mux.Handle("/", app.sessionMiddleware(app.dashboardHandler()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Application launched and running on http://127.0.0.1:%s\n", port)

	http.ListenAndServe(":"+port, mux)
}
