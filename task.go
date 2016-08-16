package taskboard

import (
	"log"
	"time"

	"errors"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const TaskEntityKind = "task"

var ErrorTaskNotFound = errors.New("Task not found in datastore")
var ErrorUnauthorized = errors.New("User does not have permission to access this task")
var ErrorUpdatingUnknownEntity = errors.New("Unable to update entity because the id does not exists in the datastore")

type Task struct { //TODO add key to id when storing
	ID          int64
	UserID      int64
	Tekst       string
	Workload    int
	WorkDone    int
	Recurring   bool
	Unit        string
	MustDo      bool
	PeriodStart time.Time
}

func PutTasks(tasks []Task, user User, ctx context.Context) ([]Task, error) {

	nrNewTasks := 0
	//Make sure the user owns all the tasks before
	//trying to persist any of them
	for _, t := range tasks {
		if t.UserID != user.UserID {
			return nil, ErrorUnauthorized
		}
		if t.ID != 0 {
			key := datastore.NewKey(ctx, TaskEntityKind, "", t.ID, nil)

			var t Task
			if err := datastore.Get(ctx, key, &t); err != nil {
				if err == datastore.ErrNoSuchEntity {
					return nil, ErrorUpdatingUnknownEntity
				}
				return nil, err
			}
			if t.UserID != user.UserID {
				return nil, ErrorUnauthorized
			}
		} else {
			nrNewTasks++
		}
	}

	var storedTasks []Task
	var keys []*datastore.Key
	intID, _, err := datastore.AllocateIDs(ctx, TaskEntityKind, nil, nrNewTasks)
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		var key *datastore.Key
		if t.ID == 0 {
			key = datastore.NewKey(ctx, TaskEntityKind, "", intID, nil)
			t.ID = key.IntID()
		} else {
			key = datastore.NewKey(ctx, TaskEntityKind, "", t.ID, nil)
		}
		storedTasks = append(storedTasks, t)
		keys = append(keys, key)
		intID++
	}

	if _, err = datastore.PutMulti(ctx, keys, storedTasks); err != nil {
		return nil, err
	}

	return storedTasks, nil
}

func DeleteTaskById(id int64, user User, ctx context.Context) error {
	key := datastore.NewKey(ctx, TaskEntityKind, "", id, nil)

	var t Task
	if err := datastore.Get(ctx, key, &t); err != nil {
		return err
	}

	if t.UserID != user.UserID {
		return ErrorUnauthorized
	}

	if err := datastore.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

func GetAllTasksForUser(u User, ctx context.Context) ([]Task, error) {
	log.Printf("Getting all tasks for user %v", u)
	return queryTask(ctx, filter{"UserID = ", u.UserID})
}

func GetTasksInPeriodForUser(u User, t time.Time, ctx context.Context) ([]Task, error) {
	log.Printf("Getting all tasks for user %v with time %v", u, t)
	periodStart := PeriodStartByTime(t)
	log.Printf("periodStart %v", periodStart)
	return queryTask(ctx, filter{"UserID = ", u.UserID}, filter{"PeriodStart = ", periodStart})
}

type filter struct {
	query string
	value interface{}
}

func queryTask(ctx context.Context, filters ...filter) ([]Task, error) {
	log.Printf("Querying using filters %v", filters)
	q := datastore.NewQuery(TaskEntityKind)

	for _, f := range filters {
		q = q.Filter(f.query, f.value)
	}

	log.Printf("query %v", q)

	var tasks []Task

	if _, err := q.GetAll(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
