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

		case 2:
			fmt.Println("===UPDATE STUDENT===")
			id := util.GetInt("Enter student id: ")
			student, err := service.GetStudentByIdService(uint(id))
			if err != nil {
				fmt.Println("Error: ", err.Error())
				break
			}
			newName := util.UpdateString("Enter new name: ", student.FullName)
			newAdd := util.UpdateString("Enter new address: ", student.Address)
			newAge := util.UpdateInt("Enter new age: ", int(student.Age))
			newEmail := util.UpdateString("Enter new email: ", student.Email)

			err = service.UpdateStudentService(id, newName, newAdd, newAge, newEmail)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Update student information successfully!!!")
			}
			
		case 3:
			fmt.Println("===STUDENT===")
			fullName := util.GetString("Enter student full name: ")
			s, err := service.GetStudentsByFullNameService(fullName)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				break
			}
			ShowListStudent(s)
		case 4:
			fmt.Println("===STUDENT===")
			id := util.GetInt("Enter student id: ")
			s, err := service.GetStudentByIdService(uint(id))
			if err != nil {
				fmt.Println("Error: ", err.Error())
				break
			}
			fmt.Printf("ID: %d | Name: %s | Address: %s | Age: %d | Email: %s\n",
			s.Id, s.FullName, s.Address, s.Age, s.Email)


		case 5:
			id := util.GetInt("Enter student id to delete: ")

			err := service.DeleteStudentByIdService(uint(id))
			if err != nil {
				fmt.Println("Error:", err.Error())
				break
			}

			fmt.Println("Student deleted successfully.")	

		case 6:
			fmt.Println("===STUDENT LIST===")
			list, err := service.GetAllStudentService()
			if err != nil {
				fmt.Println("Cannot read student list!!!")
			} else {
				ShowListStudent(list)
			}
		case 7: 
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