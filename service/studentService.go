package service

import (
	"fmt"
	"vientruongdoan/student-manager/models"
	"vientruongdoan/student-manager/repository"
)

type StudentService struct {
	studentRepository *repository.StudentRepository
}

func GetNewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{studentRepository: repo}
}

func (s *StudentService) Ping() string {
	return "Ping for student service..."
}

func (s *StudentService) GetAllStudentService() ([]models.Student, error) {
	return s.studentRepository.GetAllStudents()
}

func (s *StudentService) AddNewStudentService(
	fullName string, address string, age int, email string) (bool, error) {
	student := models.Student{
		FullName: fullName,
		Address:  address,
		Age:      uint(age),
		Email:    email,
	}

	_, error := s.studentRepository.CreateStudent(student)
	if error != nil {
		return false, fmt.Errorf("Cannot create student: %w", error)
	}

	return true, nil
}
