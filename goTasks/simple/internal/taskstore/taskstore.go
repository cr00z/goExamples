package taskstore

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

// CreateTask создаёт новую задачу в хранилище.
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
		Tags: make([]string, len(tags)),
		Due:  due,
	}
	copy(task.Tags, tags)
	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

// GetTask получает задачу из хранилища по ID. Если ID не существует -
// будет возвращена ошибка.
func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	var err error = nil
	t, inTS := ts.tasks[id]
	if !inTS {
		t, err = Task{}, fmt.Errorf("task with id=%d not found", id)
	}
	return t, err
}

// DeleteTask удаляет задачу с заданным ID. Если ID не существует -
// будет возвращена ошибка.
func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	var err error = nil
	if _, inTS := ts.tasks[id]; !inTS {
		err = fmt.Errorf("task with id=%d not found", id)
	} else {
		delete(ts.tasks, id)
	}
	return err
}

// DeleteAllTasks удаляет из хранилища все задачи.
func (ts *TaskStore) DeleteAllTasks() error {
	ts.Lock()
	defer ts.Unlock()

	ts.tasks = make(map[int]Task)
	return nil
}

// GetAllTasks возвращает из хранилища все задачи в произвольном порядке.
func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	result := make([]Task, 0, len(ts.tasks))
	for _, t := range ts.tasks {
		result = append(result, t)
	}
	return result
}

// GetTasksByTag возвращает, в произвольном порядке, все задачи
// с заданным тегом.
func (ts *TaskStore) GetTasksByTag(tag string) []Task {
	ts.Lock()
	defer ts.Unlock()

	result := make([]Task, 0)
	for _, task := range ts.tasks {
		for _, taskTag := range task.Tags {
			if taskTag == tag {
				result = append(result, task)
				break
			}
		}
	}
	return result
}

// GetTasksByDueDate возвращает, в произвольном порядке, все задачи, которые
// запланированы на указанную дату.
func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task {
	ts.Lock()
	defer ts.Unlock()

	result := make([]Task, 0)
	for _, task := range ts.tasks {
		y, m, d := task.Due.Date()
		if y == year && m == month && d == day {
			result = append(result, task)
		}
	}
	return result
}
