package datasources

import (
	"capsuler/internal/domain/user"
	"capsuler/internal/domain/user/model"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type SqlUserRepository struct {
	Db *gorm.DB
}

func NewSqlUserRepository(lc fx.Lifecycle, db *gorm.DB) user.Repository {
	return &SqlUserRepository{Db: db}
}

func (d *SqlUserRepository) Save(user *model.User) error {
	return d.Db.Save(user).Error
}
func (d *SqlUserRepository) GetById(id string) (*model.User, error) {
	var user model.User
	if err := d.Db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (d *SqlUserRepository) Remove(id string) error {
	return d.Db.Delete(&model.User{}, id).Error
}
func (d *SqlUserRepository) Count() int {
	return 0
}
func (d *SqlUserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := d.Db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
