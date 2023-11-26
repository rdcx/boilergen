package db

import "lightban/api/model"

func (db *DB) GetProjects() ([]model.Project, error) {
	var l []model.Project
	err := db.DB.Preload("Tasks").Find(&l).Error

	if err != nil {
		return nil, err
	}

	return l, nil
}

func (db *DB) GetProjectsByUserID(id uint) ([]model.Project, error) {
	var l []model.Project
	err := db.DB.Where("user_id = ?", id).Find(&l).Error

	if err != nil {
		return nil, err
	}

	return l, nil
}

func (db *DB) GetProject(id uint) (*model.Project, error) {
	var l model.Project
	err := db.DB.Where("id = ?", id).Preload("Lists.Tasks").First(&l).Error

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (db *DB) DeleteProjectByID(id uint) error {
	return db.DB.Where("id = ?", id).Delete(&model.Project{}).Error
}

func (db *DB) UpdateProject(project *model.Project) error {
	return db.DB.Save(project).Error
}

func (db *DB) CreateProject(project *model.Project) error {
	return db.DB.Create(project).Error
}
