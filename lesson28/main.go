package main

import (
	"fmt"

	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	users, err := st.GetAllStudents()
	if err != nil {
		panic(err)
	}

	user, _ := st.GetByID()

	fmt.Println("kursda oqiydigan studentlar royhati:\n", users)

	err = postgres.Create(db, model.User{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nyangi student qoshildi:\n")
	}
	err = postgres.Read(db, model.User{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("barcha studentlar royhati:\n")

	}
	err = postgres.Update(db, model.User{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("student malumotlari ozgartirildi:\n")
	}

	err = postgres.Delete(db)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("tanlangan student oqishdan chetlashtirildi:\n")

	}
	fmt.Println(user)
	cr := postgres.NewCourseRepo(db)

	courses, err := cr.GetAllCourse()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("barcha kurslar roydati\n:")
		fmt.Println(courses)
	}

	course, _ := st.GetByID()

	//fmt.Println(course)

	err = postgres.C_Create(db, model.Course{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("yangi curse qoshildi\n:")
	}

	err = postgres.C_Read(db, model.Course{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("barcha kurslar royhati:\n")
	}

	err = postgres.C_Update(db, model.Course{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("kurs malumotlari ozgartirilmoqda\n:")
	}

	err = postgres.C_Delete(db)
	if err != nil {
		fmt.Println("salom")
		panic(err)
	} else {
		fmt.Println("keraksiz kurs ochirildi:\n")
	}
	fmt.Println(course)

}
