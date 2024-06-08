package crud

import (
	"may_mod/model"

	"gorm.io/gorm"
)

type ProbuctDb struct {
	Db gorm.DB
}

func P_UserRepo(db gorm.DB) UserDb {
	return UserDb{Db: db}
}

func (u *ProbuctDb) GetAllStudents() []model.Product {
	product  := []model.Product{}
	return product
}

func P_CreateS(db gorm.DB, product model.Product) {

	db.Create(&model.Product{
		Name:   "olma",
		Year: 2020,
	})
}

func P_Read(db gorm.DB, product model.Product) *model.Product {

	db.First(&product)
	return &product

}

func P_Update(db gorm.DB, product model.Product) {

	db.Model(&model.Product{}).Where("year=?", 2020).Update("name", "toshmat")
}

func P_Delete(db gorm.DB, user model.Product) {

	db.Where("year = ?", "2023").Delete(&model.Product{})

}
