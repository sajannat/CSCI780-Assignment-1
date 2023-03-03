package main

type Course struct{
	CourseName string
	CreditHours int
	Grade float32
}

type Student struct{
	Name string
	Id string
	Major string
	Age int
	Courses []Course
}

func NewStudent(name string, id string) *Student{
	return &Student{
		Name:    name,
		Id:      id,
		Major:   "",
		Age:     0,
		Courses: []Course{},
	}
}

func (s *Student) AddCourse(courseName string, creditHours int, grade float32){
	s.Courses = append(s.Courses, Course{CourseName: courseName, CreditHours: creditHours, Grade: grade})
}

func (s *Student) CalculateGPA() float32{
	var totalCreditHours, totalGradeCreditHours float32

	if len(s.Courses) == 0{
		return 0.0
	} else {
		for _, course := range s.Courses {
			totalCreditHours += float32(course.CreditHours)
			totalGradeCreditHours += float32(course.CreditHours) * float32(course.Grade)
		}

		return totalGradeCreditHours / totalCreditHours
	}
}
