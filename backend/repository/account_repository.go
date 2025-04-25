package repository

import (
	"errors"

	"github.com/nrmadi02/go_react_monorepo/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountRepositoryIface interface {
	Login(email, password string) (*models.Account, error)
	Register(email, password, name string) (*models.Account, error)
	Me(userID uint) (*models.Account, error)
	UpdateAccessToken(accountID uint, token string) error
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) UpdateAccessToken(accountID uint, token string) error {
	return r.DB.Model(&models.Account{}).Where("id = ?", accountID).Update("access_token", token).Error
}

func (r *AccountRepository) Login(email, password string) (*models.Account, error) {
	var account models.Account
	err := r.DB.Joins("User", r.DB.Where(&models.User{Email: email})).Where(&models.Account{}).First(&account).Error
	if err != nil {
		return nil, errors.New("password or email is incorrect")
	}
	if bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) != nil {
		return nil, errors.New("password or email is incorrect")
	}
	return &account, nil
}

func (r *AccountRepository) Register(email, password, name string) (*models.Account, error) {
	user := models.User{
		Email: email,
		Name:  name,
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := models.Account{
		UserID:   user.ID,
		Password: string(hashedPassword),
		User:     user,
	}
	if err := r.DB.Create(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) Me(userID uint) (*models.Account, error) {
	var account models.Account
	err := r.DB.Joins("User", r.DB.Where(&models.User{ID: userID})).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
