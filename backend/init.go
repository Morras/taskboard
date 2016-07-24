package backend

import (
	"log"
	"net/http"

	"github.com/morras/gitserver"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {

	config := gitserver.Config{
		FilePathToFrontend: "../frontend/",
		UrlPathToApiRoot:   "/api/task",
		UrlPathToLogin:     "/loggedin",
		UrlPathToLogout:    "/loggedout",
		LoginRedirectUrl:   "/",
		LogoutRedirectUrl:  "/foo", //TODO
		Audiences:          []string{"taskboard-1279"},
		SessionDuration:    24, //TODO
	}

	userStore := &userStore{}
	ctxProvider := &contextProvider{}
	logger := &GAELogger{}

	gitserver.Setup(TaskApi{}, config, userStore, ctxProvider, logger)

	log.Print(http.ListenAndServe(":8080", nil))
}

type contextProvider struct{}

func (*contextProvider) ContextFromRequest(req *http.Request) context.Context {
	return appengine.NewContext(req)
}
