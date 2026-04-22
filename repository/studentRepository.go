package repository

import (
	"database/sql"
	"fmt"
	"vientruongdoan/student-manager/models"
)

type StudentRepository struct {
	DB *sql.DB
}

/* 
	This function to get one instance of StudentRepository 
	Then you can declare it in main function and use one instance for the whole program
	--> The reason why use pointer
*/
func GetNewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

// To bind this function belongs to StudentRepository struct
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

func (studentRepo *StudentRepository) UpdateStudent(student models.Student) error {
	query := `
		UPDATE students
		SET full_name = ?,
		    address = ?,
		    email = ?,
		    age = ?
		WHERE id = ?
	`

	result, err := studentRepo.DB.Exec(
		query,
		student.FullName,
		student.Address,
		student.Email,
		student.Age,
		student.Id,
	)
	if err != nil {
		return fmt.Errorf("cannot update student: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no student found with id %d", student.Id)
	}

	return nil
}

func (studentRepo *StudentRepository) GetStudentById(id uint) (*models.Student, error) {
	query := `
		SELECT id, full_name, address, age, email
		FROM students
		WHERE id = ?
	`

	var student models.Student

	err := studentRepo.DB.QueryRow(query, id).Scan(
		&student.Id,
		&student.FullName,
		&student.Address,
		&student.Age,
		&student.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("student with id %d not found", id)
		}
		return nil, fmt.Errorf("cannot get student by id: %w", err)
	}

	return &student, nil
}

func (studentRepo *StudentRepository) GetStudentsByFullName(fullName string) ([]models.Student, error) {
	query := `
		SELECT id, full_name, address, age, email
		FROM students
		WHERE full_name LIKE ?
	`

	search := "%" + fullName + "%"

	rows, err := studentRepo.DB.Query(query, search)
	if err != nil {
		return nil, fmt.Errorf("cannot query students by full name: %w", err)
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var student models.Student

		err := rows.Scan(
			&student.Id,
			&student.FullName,
			&student.Address,
			&student.Age,
			&student.Email,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot scan student row: %w", err)
		}

		students = append(students, student)
	}

	// check iteration error
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	// optional: return empty slice instead of nil
	if len(students) == 0 {
		return []models.Student{}, nil
	}

	return students, nil
}

func (studentRepo *StudentRepository) DeleteStudentById(id uint) error {
	query := `
		DELETE FROM students
		WHERE id = ?
	`

	result, err := studentRepo.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("cannot delete student: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot fetch affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("student with id %d not found", id)
	}

	return nil
}

func (studentRepo *StudentRepository) GetOneRecordStudentByFullName(fullname string) ([]models.Student, error) {
	query := `
		SELECT id, full_name, email, age, address
		FROM students
		WHERE full_name LIKE ?
	`

	search := "%" + fullname + "%"

	var student models.Student
	var result []models.Student

	err := studentRepo.DB.QueryRow(query, search).Scan(
		&student.Id,
		&student.FullName,
		&student.Address,
		&student.Age,
		&student.Email,
	)

	if err != nil {
		return nil, fmt.Errorf("Cannot get student data: %w", err)
	}

	result = append(result, student)

	return result, nil
}
