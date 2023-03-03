package main

import (
	"fmt"
	"strconv"
)

var programMenu string = `
1. Add a student
2. Find a student by id
3. Find a student by name
4. List students by course
5. Add a course to a student
6. Print the number of transactions executed
7. Exit 
`

type UserIO struct{}

func (u *UserIO) Input(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func (u *UserIO) Output(message string) {
	fmt.Printf(message)
}

type Program struct {
	UserIO               *UserIO
	database             *Database
	lastStudent          *Student
	transactionsExecuted int
}

func NewProgram(UserIO *UserIO, database *Database) *Program {
	return &Program{
		UserIO:               UserIO,
		database:             database,
		lastStudent:          nil,
		transactionsExecuted: 0,
	}
}

func (p *Program) run() {
	for {
		p.UserIO.Output(programMenu)
		prompt := p.UserIO.Input("Enter a number: ")

		if prompt == "1" {
			name := p.UserIO.Input("Enter Name: ")
			id := p.UserIO.Input("Enter ID: ")
			major := p.UserIO.Input("Enter Major: ")
			age := p.UserIO.Input("Enter Age: ")

			student := p.database.AddStudent(name, id)
			student.Major = major
			student.Age, _ = strconv.Atoi(age)
			p.lastStudent = student
			p.transactionsExecuted++
			p.UserIO.Output(fmt.Sprintf("Student %s added\n", name))
		} else if prompt == "2" {
			id := p.UserIO.Input("Enter ID: ")
			student := p.database.FindStudentById(id)

			if student == nil {
				p.UserIO.Output("Student not found")
			} else {
				message := fmt.Sprintf("Student\n\tName: %v\n\tID: %v\n\tMajor: %v\n\tAge: %v\n", student.Name, id, student.Major, student.Age)
				p.UserIO.Output(message)
			}

			p.lastStudent = student
			p.transactionsExecuted++
		} else if prompt == "3" {
			name := p.UserIO.Input("Enter Name: ")
			student := p.database.FindStudentByName(name)

			if student == nil {
				p.UserIO.Output("Student not found")
			} else {
				message := fmt.Sprintf("Student\n\tName: %v\n\tID: %v\n\tMajor: %v\n\tAge: %v\n", student.Name, student.Id, student.Major, student.Age)
				p.UserIO.Output(message)
			}

			p.lastStudent = student
			p.transactionsExecuted++
		} else if prompt == "4" {
			courseName := p.UserIO.Input("Enter Course Name: ")
			students := p.database.FindStudentByCourse(courseName)

			var studentList string

			if len(students) == 0 {
				studentList = "No student found"
			} else {
				for i, student := range students {
					if i == 0 {
						studentList += student.Name
					} else {
						studentList += "," + student.Name
					}
				}
				studentList = "[" + studentList + "]"
			}

			p.UserIO.Output(studentList)
			p.transactionsExecuted++
		} else if prompt == "5" {
			if p.lastStudent == nil {
				p.UserIO.Output("Search or add student before adding a course")
			} else {
				courseName := p.UserIO.Input("Course name of the student: ")
				creditHours, _ := strconv.Atoi(p.UserIO.Input("Credit hours of the student: "))
				grade, err := strconv.ParseFloat(p.UserIO.Input("Grade of the student: "), 32)
				if err != nil {

				}

				student := p.lastStudent
				student.AddCourse(courseName, creditHours, float32(grade))
				p.transactionsExecuted++
				p.UserIO.Output(fmt.Sprintf("Course %v added to student %v", courseName, student.Name))
			}
		} else if prompt == "6" {
			p.UserIO.Output(fmt.Sprintf("Transactions: %v", p.transactionsExecuted))
			p.transactionsExecuted++
		} else {
			break
		}
	}
}

func main() {
	userio := UserIO{}
	database := Database{}
	lastStudent := Student{}
	transactionsExecuted := 0
	program := Program{&userio, &database, &lastStudent, transactionsExecuted}
	program.run()
}
