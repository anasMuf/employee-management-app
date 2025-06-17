package handler

import (
	"employee-management-app/model"
	"employee-management-app/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type EmployeeHandler struct {
	Repository repositories.EmployeeRepository
}

func NewEmployeeHandler(r repositories.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{Repository: r}
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	//get all data
	employees, err := h.Repository.GetAll()
	if err != nil {
		response := map[string]any{
			"message": err.Error(),
			"data":    employees,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]any{
		"message": "employees found",
		"data":    employees,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := map[string]any{
			"message": "Invalid employee ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//get data by id
	employee, err := h.Repository.GetByID(id)
	if err != nil {
		response := map[string]any{
			"message": err.Error(),
			"data":    employee,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
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
	var employeeReq model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employeeReq); err != nil {
		response := map[string]any{
			"message": "Invalid JSON",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//input validation
	notValidated := []any{}
	if employeeReq.Name == "" {
		data := map[string]string{
			"field":   "Name",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employeeReq.Email == "" {
		data := map[string]string{
			"field":   "Email",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employeeReq.Phone == "" {
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
	employee, err := h.Repository.Create(employeeReq)
	if err != nil {
		response := map[string]any{
			"message": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
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
		response := map[string]any{
			"message": "Invalid employee ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	var employeeReq model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employeeReq); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		response := map[string]any{
			"message": "Invalid JSON",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//input validation
	notValidated := []any{}
	if employeeReq.Name == "" {
		data := map[string]string{
			"field":   "Name",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employeeReq.Email == "" {
		data := map[string]string{
			"field":   "Email",
			"message": "must be required",
		}
		notValidated = append(notValidated, data)
	}
	if employeeReq.Phone == "" {
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
	employee, err := h.Repository.Update(id, employeeReq)
	if err != nil {
		response := map[string]any{
			"message": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
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
		response := map[string]any{
			"message": "Invalid employee ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//check data exist
	employee, err := h.Repository.Delete(id)
	if err != nil {
		response := map[string]any{
			"message": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := map[string]any{
		"message": fmt.Sprintf("employee %s deleted successfully", employee.Name),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
