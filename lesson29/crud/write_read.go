package crud

import (
	"gorm.io/gorm"
	"my_mod/model"
)

type UserDb struct {
	Db gorm.DB
}

func UserRepo(db gorm.DB) UserDb {
	return UserDb{Db: db}
}

func CreateS(db gorm.DB, user model.Users) {
	
	db.Create(&model.Users{First_name: "ali",
			LastName:   "vali",
			Email:      "erfgbdgf",
			Password:   "sfdgr",
			Age:        23,
			Field:      "efsdg",
			Gender:     "dfdger",
			IsEmployee: true})
}

func Read(db gorm.DB, user model.Users)  (*model.Users) {
	
	db.First(&user)
	return &user

}

func Update(db gorm.DB, user model.Users) {

	db.Model(&model.Users{}).Where("age=?", 23).Update("First_name", "toshmat")
}

func Delete(db gorm.DB, user model.Users) {
	
	db.Where("age = ?", "23").Delete(&model.Users{})

}