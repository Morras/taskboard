package taskboard

import "net/http"
import "errors"

type User struct {
    ID int64
    Email string
}

var ErrUnauthorizedUser = errors.New("Unauthorized User")

//TODO Obviously this needs to be implemented
func GetAuthorizedUser(req *http.Request) (User, error){
    return User{1, "test@example.com"}, nil
}