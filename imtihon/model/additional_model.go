package model

//Ma'lum bir foydalanuvchiga tegishli barcha kurslarni olish ucun ishlatiladigan struct
type GetCoursesbyUsers struct {
	Id     string
	Course Courses
}

// Ma'lum bir kursga tegishli barcha darslarni olish ucun ishlatiladigan struct
type GetLessonsbyCourses struct {
	Id     string
	Lesson Lessons
}

// Ma'lum bir kursga ro'yxatdan o'tgan barcha foydalanuvchilarni olish ucun ishlatiladigan struct
type GetEnrolledUsersbyCourses struct {
	Id   string
	User Users
}

// - Foydalanuvchilarni ismi, yoshi yoki email bo'yicha qidirish. Yosh oralig’i bo’yicha qidirish ucun ishlatiladigan struct
type SearchUser struct {
	Name   string
	Email  string
	OneAge int
	TwoAge int
}

// Ma'lum bir muddat davomida eng ko'p ro'yxatdan o'tilgan kurslarni olish ucunnnishlatilgan structlar 
type MostPopularCourses struct {
	StartDate string
	EndDate   string
}

type PopularCourses struct {
	CourseId        string
	CourseTitle     string
	EnrolmentsCount int
}

type MostPopularCoursesstruct struct {
	MostPopularCourse MostPopularCourses
	CopularCourses    []PopularCourses
}

type Result struct {
	Results []Users
}
