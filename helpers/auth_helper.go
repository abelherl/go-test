package helpers

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
}

func GetAuthHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func GetUserIDFromAuth(c *gin.Context) (string, error) {
	tokenStr := GetAuthHeader(c)
	if tokenStr == "" {
		return "", jwt.NewValidationError("No token provided", jwt.ValidationErrorUnverifiable)
	}

	token, err := ValidateJWT(tokenStr)
	if err != nil || !token.Valid {
		return "", jwt.NewValidationError("Invalid token", jwt.ValidationErrorMalformed)
	}

	userID, ok := token.Claims.(jwt.MapClaims)["sub"].(float64)
	if !ok {
		return "", jwt.NewValidationError("Invalid token claims", jwt.ValidationErrorClaimsInvalid)
	}

	return strconv.Itoa(int(userID)), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
