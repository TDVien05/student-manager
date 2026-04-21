package main

import (
	"fmt"
	"log"
	"vientruongdoan/student-manager/config"
	"vientruongdoan/student-manager/models"
	"vientruongdoan/student-manager/repository"
	"vientruongdoan/student-manager/service"
	"vientruongdoan/student-manager/utils"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.GetNewStudentRepository(db)
	service := service.GetNewStudentService(repo)
	for {
		choice := ShowMenu()

		switch choice {
		case 1:
			fmt.Println("===ADDING STUDENT===")
			fullName := util.GetString("Enter student name: ")
			address := util.GetString("Enter student address: ")
			age := util.GetInt("Enter student age: ")
			email := util.GetString("Enter student email: ")
			_, err := service.AddNewStudentService(fullName, address, age, email)
			if err == nil {
				fmt.Println("Add student successfully!!!")
			}
		case 6:
			fmt.Println("===STUDENT LIST===")
			list, err := service.GetAllStudentService()
			if err != nil {
				fmt.Println("Cannot read student list!!!")
			} else {
				ShowListStudent(list)
			}
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func ShowListStudent(list []models.Student) {
	for _, s := range list {
		fmt.Printf("ID: %d | Name: %s | Address: %s | Age: %d | Email: %s\n",
			s.Id, s.FullName, s.Address, s.Age, s.Email)
	}
}

func ShowMenu() int {
	fmt.Println("=== STUDENT MANAGER ===")
	fmt.Println("1. Add new student")
	fmt.Println("2. Update student info")
	fmt.Println("3. Search student by full name")
	fmt.Println("4. Search student by ID")
	fmt.Println("5. Delete a student")
	fmt.Println("6. Show list")
	fmt.Println("7. Exit")

	choice := util.GetInt("Enter your choice: ")

	return choice
}