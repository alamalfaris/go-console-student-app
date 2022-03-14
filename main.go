package main

import (
	"bufio"
	"context"
	"fmt"
	"golang-student-app/database"
	"golang-student-app/entity"
	"golang-student-app/repository"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to student app")

	ShowMenu()
}

func ShowMenu() {
	fmt.Println("Menu")
	fmt.Println("================")
	fmt.Println("1. Show All Student")
	fmt.Println("2. Search Student")
	fmt.Println("3. Add Student")
	fmt.Println("4. Edit Student")
	fmt.Println("5. Delete Student")
	fmt.Println("6. Keluar")
	fmt.Println("")

	var strMenu string
	fmt.Print("Select menu: ")
	fmt.Scanln(&strMenu)

	switch strMenu {
	case "1":
		ShowStudent()
	case "2":
		SearchStudent()
	case "3":
		AddStudent()
	case "4":
		EditStudent()
	case "5":
		DeleteStudent()
	default:
		fmt.Println("App closed")
		time.Sleep(2 * time.Second)
	}
}

func BackToMenu(menu string) {
	switch menu {
	case "y":
		ShowMenu()
	default:
		fmt.Println("App closed")
		time.Sleep(2 * time.Second)
	}
}

func ShowStudent() {
	studentRepo := repository.NewStudentRepository(database.GetConnection())
	ctx := context.Background()
	fmt.Println("Show All Student")

	results, err := studentRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, i := range results {
		fmt.Println("")
		fmt.Println("Id:", i.Id)
		fmt.Println("Name:", strings.TrimSpace(i.Name))
		fmt.Println("Address:", strings.TrimSpace(i.Address))
		fmt.Println("Class:", strings.TrimSpace(i.Class))
	}

	fmt.Print("Go back to menu? [y/n]: ")
	var menu string
	fmt.Scanln(&menu)
	BackToMenu(menu)
}

func SearchStudent() {
	studentRepo := repository.NewStudentRepository(database.GetConnection())

	var strId string
	fmt.Println("Search Student")
	fmt.Print("Input student ID: ")
	fmt.Scanln(&strId)
	idInt, _ := strconv.Atoi(strId)
	id := int32(idInt)

	ctx := context.Background()

	result, err := studentRepo.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Id:", result.Id)
	fmt.Println("Name:", strings.TrimSpace(result.Name))
	fmt.Println("Address:", strings.TrimSpace(result.Address))
	fmt.Println("Class:", strings.TrimSpace(result.Class))

	fmt.Print("Go back to menu? [y/n]: ")
	var menu string
	fmt.Scanln(&menu)
	BackToMenu(menu)
}

func AddStudent() {
	studentRepo := repository.NewStudentRepository(database.GetConnection())
	ctx := context.Background()
	fmt.Println("Add Student")

	student := entity.Student{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Name: ")
	student.Name, _ = reader.ReadString('\n')
	fmt.Print("Input City: ")
	student.Address, _ = reader.ReadString('\n')
	fmt.Print("Input Class: ")
	student.Class, _ = reader.ReadString('\n')
	student.Name = strings.TrimSpace(student.Name)
	student.Address = strings.TrimSpace(student.Address)
	student.Class = strings.TrimSpace(student.Class)

	result, err := studentRepo.Insert(ctx, student)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert student with id:", result.Id)
	fmt.Println("Name:", result.Name)
	fmt.Println("Address:", result.Address)
	fmt.Println("Class:", result.Class)

	fmt.Print("Go back to menu? [y/n]: ")
	var menu string
	fmt.Scanln(&menu)
	BackToMenu(menu)
}

func EditStudent() {
	studentRepo := repository.NewStudentRepository(database.GetConnection())
	ctx := context.Background()
	fmt.Println("Edit Student")

	student := entity.Student{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Student Id you want edit: ")
	strStudentId, _ := reader.ReadString('\n')
	studentId, _ := strconv.Atoi(strings.TrimSpace(strStudentId))
	fmt.Print("Input Name: ")
	student.Name, _ = reader.ReadString('\n')
	fmt.Print("Input City: ")
	student.Address, _ = reader.ReadString('\n')
	fmt.Print("Input Class: ")
	student.Class, _ = reader.ReadString('\n')

	student.Id = int32(studentId)
	student.Name = strings.TrimSpace(student.Name)
	student.Address = strings.TrimSpace(student.Address)
	student.Class = strings.TrimSpace(student.Class)

	result, err := studentRepo.Update(ctx, student)
	if err != nil {
		panic(err)
	}
	fmt.Println(result, "data was updated successfully")

	fmt.Print("Go back to menu? [y/n]: ")
	var menu string
	fmt.Scanln(&menu)
	BackToMenu(menu)
}

func DeleteStudent() {
	studentRepo := repository.NewStudentRepository(database.GetConnection())
	ctx := context.Background()
	fmt.Println("Delete Student")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Student Id you want edit: ")
	strId, _ := reader.ReadString('\n')
	studentId, _ := strconv.Atoi(strings.TrimSpace(strId))
	id := int32(studentId)

	result, err := studentRepo.Delete(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result, "data was deleted successfully")

	fmt.Print("Go back to menu? [y/n]: ")
	var menu string
	fmt.Scanln(&menu)
	BackToMenu(menu)
}
