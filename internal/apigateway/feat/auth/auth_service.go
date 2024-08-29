package auth

import (
	"banking/internal/apigateway/feat/users"
	"banking/internal/apigateway/shared"
	. "banking/internal/apigateway/shared/errors"
	"banking/pkg"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func SignIn(email string, password string) (*TokenDto, *shared.HttpError) {
	user, err := users.QueryUserByEmail(email)
	if err != nil {
		return nil, Unauthorized(Params{Detail: err.Error()})
	}

	isMatch, err := pkg.Argon2iCheckPassword(password, user.Hash)
	if !isMatch || err != nil {
		return nil, Unauthorized(Params{Detail: err.Error()})
	}

	return createTokens(users.ToUserDto(user)), nil
}

// region - Public functions

func createTokens(user *users.UserDto) *TokenDto {
	now := time.Now()

	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub":      user.Id,
		"username": user.Username,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Hour * 1).Unix(),
	}).SignedString(pkg.Certs.AccessTokenPrivate)

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.Id,
		"iat": now.Unix(),
		"exp": now.Add(time.Hour * 24).Unix(),
	}).SignedString(pkg.Certs.RefreshTokenPrivate)

	return &TokenDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

// endregion
