package branch

import "gorm.io/gorm"

type Repository interface {
	Create(branch Branch) (Branch, error)
	FindByID(ID string) (Branch, error)
	Update(branch Branch) (Branch, error)
	//GetAll() ([]Branch, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(branch Branch) (Branch, error) {
	err := r.db.Create(&branch).Error
	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (r *repository) FindByID(ID string) (Branch, error) {
	var branch Branch

	err := r.db.Where("id = ?", ID).Find(&branch).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}

func (r *repository) Update(branch Branch) (Branch, error) {
	err := r.db.Save(&branch).Error

	if err != nil {
		return branch, err
	}

	return branch, nil
}
