package main

import (
	"fmt"
)

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var students []Student

func updateStudent() {
	var studentID string
	fmt.Print("Enter student ID to update: ")
	fmt.Scanln(&studentID)

	index := findStudentIndex(studentID)
	if index == -1 {
		fmt.Println("Student not found.")
		return
	}

	var updatedStudent Student
	fmt.Print("Enter updated student name: ")
	fmt.Scanln(&updatedStudent.Name)

	students[index] = updatedStudent
	fmt.Println("Student updated successfully!")
}

func deleteStudent() {
	var studentID string
	fmt.Print("Enter student ID to delete: ")
	fmt.Scanln(&studentID)

	index := findStudentIndex(studentID)
	if index == -1 {
		fmt.Println("Student not found.")
		return
	}

	students = append(students[:index], students[index+1:]...)
	fmt.Println("Student deleted successfully!")
}

func listStudents() {
	if len(students) == 0 {
		fmt.Println("No students in the list.")
		return
	}

	fmt.Println("List of Students:")
	for _, student := range students {
		fmt.Printf("ID: %s, Name: %s, \n", student.ID, student.Name)
	}
}

func findStudentIndex(studentID string) int {
	for i, student := range students {
		if student.ID == studentID {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Print("----------------- Add Student -----------------\n")

	var newStudent Student
	fmt.Print("Enter student ID: ")
	fmt.Scanln(&newStudent.ID)
	fmt.Print("Enter student name: ")
	fmt.Scanln(&newStudent.Name)

	students = append(students, newStudent)
	fmt.Println("Student added successfully!")

	fmt.Print("----------------- List Student -----------------\n")
	listStudents()

	fmt.Print("----------------- Update Student -----------------\n")
	updateStudent()

	fmt.Print("----------------- List Student -----------------\n")
	listStudents()

	fmt.Print("----------------- Delete Student -----------------\n")
	deleteStudent()
}
