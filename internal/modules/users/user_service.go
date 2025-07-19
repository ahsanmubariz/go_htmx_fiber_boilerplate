package users

import "gorm.io/gorm"

type UsersService struct {
	db *gorm.DB
}

func NewUsersService(db *gorm.DB) *UsersService {
	return &UsersService{db: db}
}

func (s *UsersService) GetAllUsers() ([]User, error) {
	var users []User
	result := s.db.Order("created_at desc").Find(&users)
	return users, result.Error
}

func (s *UsersService) CreateUser(user *User) error {
	result := s.db.Create(user)
	return result.Error
}

func (s *UsersService) DeleteUser(id uint) error {
	result := s.db.Delete(&User{}, id)
	return result.Error
}
