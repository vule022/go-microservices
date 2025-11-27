package helper

import (
	"errors"
	"fmt"
	"go-microservices/internal/domain"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {
	if len(p) < 6 {
		return "", errors.New("password length should be at least 6 chars long")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		return "", errors.New("password hashing failed")
	}

	return string(hashP), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs missing")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"expiry":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("unablde to sign the token")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(hP string, pP string) error {
	if len(pP) < 6 {
		return errors.New("password length should be at least 6 chars long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(pP), []byte(hP))

	if err != nil {
		log.Printf("invavlid %v, %v", hP, pP)
		return errors.New("invalid password")
	}

	return nil
}

func (a Auth) VerifyToken(authHeader string) (domain.User, error) {
	tokenArr := strings.Split(authHeader, " ")

	if len(tokenArr) != 2 {
		return domain.User{}, errors.New("invalid authorization header format")
	}

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("missing Bearer prefix")
	}

	tokenStr := tokenArr[1]

	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, err
	}

	if !jwtToken.Valid {
		return domain.User{}, errors.New("invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if ok && jwtToken.Valid {
		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			return domain.User{}, errors.New("token expired")
		}
	}

	return domain.User{
		ID:       uint(claims["user_id"].(float64)),
		Email:    claims["email"].(string),
		UserType: claims["role"].(string),
	}, nil
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	user, err := a.VerifyToken(authHeader)

	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)

		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")

	return user.(domain.User)
}
