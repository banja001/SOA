package service

import (
	"errors"
	"os"
	"stakeholders/dto"
	"stakeholders/model"
	"stakeholders/repo"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
	UserRepository *repo.UserRepository
}

func (service *AuthenticationService) Login(credentials *dto.Credentials) (*dto.AuthenticationTokens, error) {
	user, err := service.UserRepository.FindActiveByUsername(credentials.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	tokens, err := service.GenerateAccessToken(user.ID, user.Username, user.Role, user.PersonId)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (service *AuthenticationService) GenerateAccessToken(userId int, username string, role model.UserRole, personId int) (*dto.AuthenticationTokens, error) {
	var secretKey = getEnv("JWT_KEY", "explorer_secret_key")
	var issuer = getEnv("JWT_ISSUER", "explorer")
	var audience = getEnv("JWT_AUDIENCE", "explorer-front.com")

	claims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       strconv.FormatInt(int64(userId), 10),
		"username": username,
		"personId": personId,
		"role":     role.String(),
		"exp":      time.Now().Add(time.Minute * 60 * 24).Unix(),
		"iss":      issuer,
		"aud":      audience,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	var returnToken = dto.AuthenticationTokens{
		Id:          userId,
		AccessToken: signedToken,
	}

	return &returnToken, nil
}

func getEnv(key, fallbackValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallbackValue
	}
	return value
}
