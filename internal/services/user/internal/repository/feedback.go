package repository

import (
	"github.com/pkg/errors"
	"user/internal/repository/model"
)

func (d *User) MakeFeedback(feedback *model.Feedback) error {
	tx := d.db.Begin()
	if err := tx.Table(feedback.Table()).Create(feedback).
		Error; err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	
	return errors.WithStack(tx.Commit().Error)
}
