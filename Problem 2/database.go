package main

type Database struct{
	students []*Student
}

func NewDatabase() *Database {
	return &Database{students: []*Student{}}
}

func (db *Database) AddStudent(name string, id string) *Student{
	student := NewStudent(name, id)
	db.students = append(db.students, student)
	return student
}

func (db *Database) FindStudentById(id string) *Student{
	for _, student := range db.students {
		if student.Id == id {
			return student
		}
	}
	return nil
}

func (db *Database) FindStudentByName(name string) *Student {
	for _, student := range db.students {
		if student.Name == name {
			return student
		}
	}
	return nil
}

func (db *Database) FindStudentByCourse(courseName string) []*Student {
	var studentsByCourse []*Student
	for _, student := range db.students {
		for _,course := range student.Courses {
			if course.CourseName == courseName {
				studentsByCourse = append(studentsByCourse, student)
				break
			}
		}
	}
	return studentsByCourse
}

func (db *Database) NumStudent() int {
	return len(db.students)
}