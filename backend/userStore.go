package backend

import (
	"golang.org/x/net/context"

	"errors"
	"github.com/morras/gitserver"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	userEntityKind = "user"
)

type user struct {
	Email       string
	UserID      string
	GitSessions []gitserver.Session
}

type userStore struct {
}

func (us *userStore) UpdateUser(ctx context.Context, gitUser *gitserver.User) error {
	log.Infof(ctx, "Update user called with %v", gitUser)

	if gitUser.ID == "" {
		log.Warningf(ctx, "Error UserID is empty")
		return errors.New("ID must not be empty")
	}

	//Get existing user from the datastore
	var u user

	key := datastore.NewKey(ctx, userEntityKind, gitUser.ID, 0, nil)
	if err := datastore.Get(ctx, key, &u); err != nil && err != datastore.ErrNoSuchEntity {
		log.Warningf(ctx, "Error getting user %s (%s) from the user store: %v", gitUser.ID, key, err)
		return err
	}

	//Update the user and save it
	u.Email = gitUser.Email
	u.UserID = gitUser.ID
	u.GitSessions = gitUser.Sessions

	_, err := datastore.Put(ctx, key, &u)
	if err != nil {
		log.Warningf(ctx, "Error Saving user %s (%s) to the user store: %v", gitUser.ID, key, err)
		return err
	}

	return nil
}

func (us *userStore) LookupUser(ctx context.Context, id string) (*gitserver.User, error) {
	log.Infof(ctx, "Lookup user called with %v", id)

	q := datastore.NewQuery(userEntityKind)

	q = q.Filter("UserID = ", id)

	var u []user

	_, err := q.GetAll(ctx, &u)
	if err == datastore.ErrNoSuchEntity || len(u) == 0 {
		return nil, gitserver.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	//if u[0].UserID != id || u[0].User.ID != u[0].UserID {
	//	log.Errorf(ctx, "ID: %s. UserID: %s, User.ID: %s", id, u[0].UserID, u[0].User.ID)
	//}/

	log.Infof(ctx, "Retrieved users: %v", u)

	return &gitserver.User{
		ID:       u[0].UserID,
		Email:    u[0].Email,
		Sessions: u[0].GitSessions,
	}, nil
}
