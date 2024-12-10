package repositories

import (
	"compartamos-backend/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func (r *UserRepository) CreateUser(user models.User) (models.User, error) {

    user.Age = 19

    err := r.DB.Create(&user).Error
    return user, err
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
    var users []models.User
    err := r.DB.Find(&users).Error
    return users, err
}

func (r *UserRepository) GetUser(id string) (models.User, error) {
    var user models.User
    err := r.DB.First(&user, id).Error
    return user, err
}

func (r *UserRepository) UpdateUser(id string, input models.User) (models.User, error) {
    var user models.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return user, err
    }

    fmt.Println(input)
    
    r.DB.Model(&user).Updates(input)
    return user, nil
}

func (r *UserRepository) DeleteUser(id string) error {
    var user models.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return err
    }
    return r.DB.Delete(&user).Error
}

func calculateAge(birthDate time.Time) int {
    today := time.Now()
    age := today.Year() - birthDate.Year()
    
    if today.YearDay() < birthDate.YearDay() {
        age--
    }
    return age
}
