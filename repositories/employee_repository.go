package repositories

import (
	"database/sql"
	"employee-management-app/model"
	"errors"
	"time"
)

type EmployeeRepository interface {
	GetAll() ([]model.Employee, error)
	GetByID(id int) (model.Employee, error)
	Create(employee model.Employee) (model.Employee, error)
	Update(id int, employee model.Employee) (model.Employee, error)
	Delete(id int) (model.Employee, error)
}

type employeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	//get all data
	rows, err := r.DB.Query("SELECT * FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//collect data to struct
	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Phone, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, err
}

func (r *employeeRepository) GetByID(id int) (model.Employee, error) {
	var employee model.Employee
	//get data by id
	err := r.DB.QueryRow(`
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
	return employee, err
}

func (r *employeeRepository) Create(employee model.Employee) (model.Employee, error) {
	now := time.Now().Format("2006-01-02 15:04:05") //get current datetime

	//store data
	result, err := r.DB.Exec(`
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
		return model.Employee{}, err
	}
	id, _ := result.LastInsertId()
	employee.ID = int(id)
	employee.CreatedAt = now
	employee.UpdatedAt = now
	return employee, err
}

func (r *employeeRepository) Update(id int, employeeReq model.Employee) (model.Employee, error) {
	now := time.Now().Format("2006-01-02 15:04:05") //get current datetime

	//update data by id
	result, err := r.DB.Exec(`
	UPDATE employees 
	SET name=?, 
		email=?, 
		phone=?, 
		updated_at=? 
	WHERE id=?`,
		employeeReq.Name,
		employeeReq.Email,
		employeeReq.Phone,
		now,
		id,
	)
	if err != nil {
		return model.Employee{}, err
	}

	// cek apakah ada row yang ter-update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Employee{}, err
	}
	if rowsAffected == 0 {
		return model.Employee{}, errors.New("employee not found")
	}

	// ambil data lengkap setelah update untuk memastikan konsistensi
	var updatedEmployee model.Employee
	err = r.DB.QueryRow(`
		SELECT id, name, email, phone, created_at, updated_at 
		FROM employees 
		WHERE id=?`, id).Scan(
		&updatedEmployee.ID,
		&updatedEmployee.Name,
		&updatedEmployee.Email,
		&updatedEmployee.Phone,
		&updatedEmployee.CreatedAt,
		&updatedEmployee.UpdatedAt,
	)
	if err != nil {
		return model.Employee{}, err
	}

	return updatedEmployee, err
}

func (r *employeeRepository) Delete(id int) (model.Employee, error) {
	var employee model.Employee

	//check data exist
	err := r.DB.QueryRow(`
	SELECT id,name 
	FROM employees 
	WHERE id=?`, id).Scan(
		&employee.ID,
		&employee.Name,
	)
	if err == sql.ErrNoRows {
		return model.Employee{}, err
	} else if err != nil {
		return model.Employee{}, err
	}

	//delete data by id
	_, err = r.DB.Exec("DELETE FROM employees WHERE id=?", id)
	return employee, err
}
