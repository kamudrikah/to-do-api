package services

import (
	"encoding/json"
	"fmt"
	"kamudrikah/to-do-api/models"
)

func GetTasksFromDB() []byte {
	var tasks []models.Task
	rows, err := db.Query("SELECT * FROM to_dos ORDER BY id DESC")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.ID, &task.Title, &task.Checked)
		tasks = append(tasks, task)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(tasks)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

func InsertTaskInDB(taskDetails models.Task) bool {
	stmt, err := db.Prepare("INSERT INTO to_dos SET title=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(taskDetails.Title)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func UpdateTaskInDB(taskDetails models.Task) bool {
	stmt, err := db.Prepare("UPDATE to_dos SET title=?,checked=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(taskDetails.Title, taskDetails.Checked, taskDetails.ID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func DeleteTaskFromDB(taskID string) bool {
	stmt, err := db.Prepare("DELETE FROM to_dos WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(taskID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}
