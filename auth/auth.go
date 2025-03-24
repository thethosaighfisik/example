package auth

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"auth_service/models"
)

var jwtSecret = []byte("my_secret_key")

type Claims struct {
	ID int `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func AuthenticateUser(db *sql.DB, email, password string) (string, error) {
	
	user, err := models.GetUserByEmail(db, email)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	claims := &Claims{
		ID: user.ID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24*time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("couldn't generate token")
	}

	return tokenString, nil
}
