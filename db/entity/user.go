package entity

import (
	"time"

	"github.com/den19980107/go-fiber-gorm-starter/config"
	"github.com/den19980107/go-fiber-gorm-starter/log"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type JWTToken struct {
	Hash   string `json:"access_token"`
	Expire int64  `json:"expires_in"`
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"username,index,unique"`
	PasswordHash string `json:"-" gorm:"password"`
}

type UserRegisterDto struct {
	Username  string `json:"username" form:"username" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	CPassword string `json:"c_password" form:"c_password" validate:"required|eq_field:password"`
}

type UserLoginDto struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Create the user JWT token.
func (u *User) CreateJWTToken(secret string, expires ...int64) *JWTToken {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expire := config.App.JWT.Expire

	if len(expires) > 0 {
		expire = expires[0]
	}

	expiresIn := time.Now().Add(time.Duration(expire) * time.Second).Unix()

	claims["username"] = u.Username
	claims["exp"] = expiresIn

	tokenHash, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return &JWTToken{
		Hash:   tokenHash,
		Expire: expiresIn,
	}
}
