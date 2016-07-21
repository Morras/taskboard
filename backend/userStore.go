package backend

import (
    
    "golang.org/x/net/context"
    
	"github.com/morras/gitserver"
    "google.golang.org/appengine/log"
)

type userStore struct {
}

func (u userStore) UpdateUser(ctx context.Context, user *gitserver.User) (*gitserver.User, error) {
    log.Infof(ctx, "Update user called with %v", user)
	return nil, nil
}

func (u userStore) LookupUser(ctx context.Context, id string) (*gitserver.User, error) {
	log.Infof(ctx, "Lookup user called with %v", id)
	return nil, nil
}

func (u userStore) Sessions(ctx context.Context, id string) []gitserver.Session {
	log.Infof(ctx, "sessions called with  %v", id)
	return nil
}
