package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeeders(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	users := models.User{
		UUID:        uuid.New(),
		Name:        "Administrator",
		Username:    "admin",
		PhoneNumber: "08123456789",
		Email:       "admin@admin.com",
		Password:    string(password),
		RoleID:      constants.Admin,
	}

	err := db.FirstOrCreate(&users, models.User{Username: users.Username}).Error
	if err != nil {
		logrus.Errorf("failed to create user %s: %v", users.Email, err)
		panic(err)
	}
	logrus.Infof("user %s created successfully", users.Email)

}
