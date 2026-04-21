package repository

import (
	"database/sql"
	"fmt"
	"vientruongdoan/student-manager/models"
)

type StudentRepository struct {
	DB *sql.DB
}

func GetNewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{DB : db}
}

func (studentRepo *StudentRepository) GetAllStudents() ([]models.Student, error) {
	var students []models.Student

	rows, error := studentRepo.DB.Query("SELECT * FROM students")
	
	if error != nil {
		return nil, fmt.Errorf("Can not get list of student: %w", error)
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if error := rows.Scan(&student.Id, &student.FullName, &student.Address, &student.Age, &student.Email); error != nil {
			return nil, fmt.Errorf("Cannot read data: %w", error)
		}

		students = append(students, student)
	}

	if error := rows.Err(); error != nil {
		return nil, fmt.Errorf("Error while fetching student: %w", error)
	}

	return students, nil
} 

func (studentRepo *StudentRepository) CreateStudent(student models.Student) (uint, error) {
	query := `
		INSERT INTO students (full_name, address, age, email)
		VALUES (?, ?, ?, ?)
	`

	result, err := studentRepo.DB.Exec(
		query,
		student.FullName,
		student.Address,
		student.Age,
		student.Email,
	)
	if err != nil {
		return 0, fmt.Errorf("cannot insert student: %w", err)
	}

	// Get auto-generated ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("cannot get inserted id: %w", err)
	}

	return uint(id), nil
}