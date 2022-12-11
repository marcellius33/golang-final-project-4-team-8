package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"testing"
	"tokobelanja/models"
)

var (
	user = &models.User{
		ID:    1,
		Email: "testing@gmail.com",
		Role:  "member",
	}
)

func TestSuccessGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		t.Errorf("Terjadi kesalahan : %s", err.Error())
		return

	}
	t.Logf("Berhasil generate token : %s", signedToken)
}

func TestFailedGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString("")

	if err != nil {
		t.Errorf("Terjadi kesalahan : %s", err.Error())
		return

	}
	t.Logf("Berhasil generate token : %s", signedToken)
}
