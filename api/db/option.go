package db

import "lightban/api/model"

func (db *DB) CreateOption(op *model.Option) error {
	return db.Create(op).Error
}
