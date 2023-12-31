package AuthImpl

import (
	"context"
	"strings"

	Auth "FM/src/auth"
	"FM/src/auth/models"

	"FM/src/entities"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	*gorm.DB
}

func NewAuthRepositoryImpl(DB *gorm.DB) Auth.AuthRepository {
	return &AuthRepositoryImpl{
		DB: DB,
	}
}

func (authRepository *AuthRepositoryImpl) SignInWithGoogle(ctx context.Context, req models.Payload) (entities.User, error) {
	var user entities.User

	email := req.Email
	isExist := authRepository.DB.WithContext(ctx).Where("email = ?", email).Find(&user)

	if isExist.Error != nil {
		return entities.User{}, isExist.Error
	}

	if isExist.RowsAffected == 0 {
		var roles string = entities.STAFF
		if !strings.HasSuffix(email, "@fpt.edu.vn") {
			roles = entities.TEACHER

		}

		newUser := entities.User{
			Email:    email,
			Url:      req.Picture,
			Name:     req.Name,
			Position: req.Position,
			Role:     roles,
		}

		result := authRepository.DB.WithContext(ctx).Create(&newUser)
		if result.Error != nil {
			return entities.User{}, result.Error
		}
		user = newUser
	} else {
		user.Url = req.Picture
		user.Name = req.Name
		result := authRepository.DB.WithContext(ctx).Save(&user)
		if result.Error != nil {
			return entities.User{}, result.Error
		}
	}
	return user, nil
}
