package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	taskID := r.URL.Query().Get("task_id")
	if taskID == "" {
		getDataTask, err := t.taskService.GetTasks(r.Context(), idLogin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getDataTask)
		return

	} else {
		taskIdConv, _ := strconv.Atoi(taskID)
		getDataTask, err := t.taskService.GetTaskByID(r.Context(), taskIdConv)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getDataTask)
		return
	}

	// TODO: answer here
}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	if task.Title == "" || task.Description == "" || strconv.Itoa(task.CategoryID) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	dataTask := entity.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      idLogin,
	}
	saveTask, err := t.taskService.StoreTask(r.Context(), &dataTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	var resp = map[string]interface{}{}
	resp["user_id"] = idLogin
	resp["task_id"] = saveTask.ID
	resp["message"] = "success create new task"
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	return
	// TODO: answer here

}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	taskID := r.URL.Query().Get("task_id")
	taskIDConv, _ := strconv.Atoi(taskID)

	err = t.taskService.DeleteTask(r.Context(), taskIDConv)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	var resp = map[string]interface{}{}
	resp["user_id"] = idLogin
	resp["task_id"] = taskIDConv
	resp["message"] = "success delete task"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return

	// TODO: answer here
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	dataTask := entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
	}
	updateDataTask, err := t.taskService.UpdateTask(r.Context(), &dataTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	var resp = map[string]interface{}{}
	resp["user_id"] = idLogin
	resp["task_id"] = updateDataTask.ID
	resp["message"] = "success update task"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	return
	// TODO: answer here
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     idLogin,
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": idLogin,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
