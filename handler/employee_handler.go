package handler

import (
	"app/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type EmployeeHandler struct {
	DB *sql.DB
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	//get all data
	rows, err := h.DB.Query("SELECT * FROM employees")
	if err != nil {
		http.Error(w, "internal server error, failed get data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//collect data to struct
	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Phone, &e.CreatedAt, &e.UpdatedAt); err != nil {
			http.Error(w, "internal server error, failed collect data", http.StatusInternalServerError)
			return
		}
		employees = append(employees, e)
	}

	response := map[string]any{
		"message": "employee found",
		"data":    employees,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	var employee model.Employee
	//get data by id
	err = h.DB.QueryRow(`
	SELECT * 
	FROM employees 
	WHERE id=?`, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Phone,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("internal server error, failed get data: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	response := map[string]any{
		"message": "employee found",
		"data":    employee,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05") //get current datetime

	//input validation
	notValidated := []any{}
	if employee.Name == "" {
		data := map[string]string{
			"field":   "Name",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employee.Email == "" {
		data := map[string]string{
			"field":   "Email",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employee.Phone == "" {
		data := map[string]string{
			"field":   "Phone",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if len(notValidated) > 0 {
		response := map[string]any{
			"message": "error validation!",
			"errors":  notValidated,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//store data
	result, err := h.DB.Exec(`
	INSERT INTO employees 
	(name, email, phone, created_at, updated_at) VALUES 
	(?, ?, ?, ?, ?)`,
		employee.Name,
		employee.Email,
		employee.Phone,
		now,
		now,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, _ := result.LastInsertId()
	employee.ID = int(id)
	employee.CreatedAt = now
	employee.UpdatedAt = now
	response := map[string]any{
		"message": "employee stored successfully",
		"data":    employee,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05") //get current datetime

	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//input validation
	notValidated := []any{}
	if employee.Name == "" {
		data := map[string]string{
			"field":   "Name",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employee.Email == "" {
		data := map[string]string{
			"field":   "Email",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employee.Phone == "" {
		data := map[string]string{
			"field":   "Phone",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if len(notValidated) > 0 {
		response := map[string]any{
			"message": "error validation!",
			"errors":  notValidated,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//update data by id
	_, err = h.DB.Exec(`
	UPDATE employees 
	SET name=?, 
		email=?, 
		phone=?, 
		updated_at=? 
	WHERE id=?`,
		employee.Name,
		employee.Email,
		employee.Phone,
		now,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employee.ID = id
	employee.UpdatedAt = now
	response := map[string]any{
		"message": "employee updated successfully",
		"data":    employee,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee model.Employee

	//check data exist
	err = h.DB.QueryRow(`
	SELECT id,name 
	FROM employees 
	WHERE id=?`, id).Scan(
		&employee.ID,
		&employee.Name,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("internal server error, failed get data: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	//delete data by id
	_, err = h.DB.Exec("DELETE FROM employees WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"message": fmt.Sprintf("employee %s deleted successfully", employee.Name),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
