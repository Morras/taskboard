package backend

import (
	"net/http"

	"github.com/morras/gitserver"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {

	config := gitserver.Config{
		URLPathToLogin:     "/loggedin",
		URLPathToLogout:    "/loggedout",
		LoginRedirectURL:   "/foo",
		LogoutRedirectURL:  "/login.html?mode=select", //TODO
		NewUserRedirectURL: "/bar",
		Audiences:          []string{"taskboard-1279"},
		SessionDuration:    24, //TODO
	}

	userStore := &userStore{}
	ctxProvider := &contextProvider{}
	logger := &GAELogger{}

	serveMux := http.NewServeMux()

	//Safe as it can only serve files from within the frontend directory
	//At least according to the source but the doc does not mention this
	fileHandler := http.FileServer(http.Dir("../frontend/"))

	serveMux.Handle("/api/task", TaskApi{})

	serveMux.Handle("/", fileHandler)

	gitserver.Setup(serveMux, config, userStore, ctxProvider, logger)

	http.DefaultServeMux = serveMux
}

type contextProvider struct{}

func (*contextProvider) ContextFromRequest(req *http.Request) context.Context {
	return appengine.NewContext(req)
}
