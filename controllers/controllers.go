package controllers

import (
	"encoding/json"
	"kamudrikah/to-do-api/models"
	"kamudrikah/to-do-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasks(response http.ResponseWriter, request *http.Request) {
	var httpError = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	jsonResponse := services.GetTasksFromDB()

	if jsonResponse == nil {
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func InsertTask(response http.ResponseWriter, request *http.Request) {
	var httpError = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's meaa.",
	}
	var taskDetails models.Task
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&taskDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if taskDetails.Title == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isInserted := services.InsertTaskInDB(taskDetails)
			if isInserted {
				GetTasks(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func UpdateTask(response http.ResponseWriter, request *http.Request) {
	var httpError = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var taskDetails models.Task
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&taskDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if taskDetails.Title == "" {
			httpError.Message = "Title can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if taskDetails.ID == 0 {
			httpError.Message = "Task ID can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isUpdated := services.UpdateTaskInDB(taskDetails)
			if isUpdated {
				GetTasks(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func DeleteTask(response http.ResponseWriter, request *http.Request) {
	var httpError = models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
		// Code: http.StatusInternalServerError, Message: request.,
	}
	taskID := mux.Vars(request)["id"]
	if taskID == "" {
		httpError.Message = "Task id can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		isdeleted := services.DeleteTaskFromDB(taskID)
		if isdeleted {
			GetTasks(response, request)
		} else {
			returnErrorResponse(response, request, httpError)
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage models.ErrorResponse) {
	httpResponse := &models.ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
