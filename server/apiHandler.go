package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/morras/taskboard"
	"google.golang.org/appengine"
)

type TaskApi struct {
}

func (api TaskApi) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	method := req.Method

	user, err := taskboard.GetAuthorizedUser(req)
	if err != nil {
		if err == taskboard.ErrUnauthorizedUser {
			res.WriteHeader(http.StatusUnauthorized) //TODO RFC says I MUST return a www-Authentication header
		} else {
			res.WriteHeader(http.StatusInternalServerError)
		}
	}

	switch method {
	case "GET":
		getTasks(res, req, user)
		return
	case "PUT":
		putTask(res, req, user)
		return
	case "DELETE":
		deleteTask(res, req, user)
		return
	default:
		res.Header().Set("Allow", "GET, PUT, DELETE")
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func getTasks(res http.ResponseWriter, req *http.Request, user taskboard.User) {

	ctx := appengine.NewContext(req)

	queryValues := req.URL.Query()

	dateInPeriod := queryValues.Get("dateInPeriod")

	log.Printf("Query values %v", queryValues)
	log.Printf("dateInPeriod %v", dateInPeriod)

	var t []taskboard.Task
	var err error

	if dateInPeriod == "" {
		log.Printf("Empty period")
		t, err = taskboard.GetAllTasksForUser(user, ctx)
	} else {
		log.Printf("Non empty period")
		format := "2006-01-02"
		time, err := time.ParseInLocation(format, dateInPeriod, taskboard.GetDefaultLocation())
		if checkAndWriteError(err, res) {
			return
		}
		t, err = taskboard.GetTasksInPeriodForUser(user, time.UTC(), ctx)
	}
	log.Printf("t %v, err %v", t, err)
	writeTasksResponse(t, res, err)
}

func writeTasksResponse(tasks []taskboard.Task, res http.ResponseWriter, err error) {
	if checkAndWriteError(err, res) {
		return
	}

	json, err := json.Marshal(tasks)
	if checkAndWriteError(err, res) {
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

//Returns true if there was an error and execution should stop,
//(meaning that an error has been reported and nothing else should be written)
func checkAndWriteError(err error, res http.ResponseWriter) bool {
    if err == taskboard.ErrorTaskNotFound {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return true
	} else if err == taskboard.ErrorUnauthorized {
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte(err.Error()))
		return true
	} else if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return true
	}
	return false
}

func putTask(res http.ResponseWriter, req *http.Request, user taskboard.User) {
	ctx := appengine.NewContext(req)

	var tasks []taskboard.Task

	dec := json.NewDecoder(req.Body)
	for {
		var t taskboard.Task
		if err := dec.Decode(&t); err == io.EOF {
			break
		} else if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}
        //Should probably look into location here, but I might also give up on it
        t.PeriodStart = taskboard.PeriodStartByTime(t.PeriodStart) 	
        if t.UserID == 0 {
            t.UserID = user.ID;
        }	
        tasks = append(tasks, t)
	}

	storedTasks, err := taskboard.PutTasks(tasks, user, ctx)
	if checkAndWriteError(err, res) {
		return
	}

	json, err := json.Marshal(storedTasks)
	if checkAndWriteError(err, res) {
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

func deleteTask(res http.ResponseWriter, req *http.Request, user taskboard.User) {

	ctx := appengine.NewContext(req)

	queryValues := req.URL.Query()

	taskID := queryValues.Get("TaskID")

	if taskID == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Delete must have query parameter TaskID"))
		return
	}

	id, err := strconv.ParseInt(taskID, 10, 64)
	if checkAndWriteError(err, res) {
		return
	}

	err = taskboard.DeleteTaskById(id, user, ctx)
	if checkAndWriteError(err, res) {
		return
	}
}
