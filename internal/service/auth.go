package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lavatee/messenger/internal/repository"
)

type AuthService struct {
	repo repository.Auth
}

const (
	salt            = "3c453453d"
	signingKey      = "c38jxmk"
	accessTokenTT   = 1 * time.Hour
	refreshTokenTT  = 7 * 24 * time.Hour
	accessTokenTTL  = 15 * time.Second
	refreshTokenTTL = 30 * time.Second
)

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) SignUp(username string, name string, password string) (int, error) {
	PasswordHash := s.HashPassword(password)
	return s.repo.SignUp(username, name, PasswordHash)
}
func (s *AuthService) HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (s *AuthService) SignIn(username string, password string) (int, string, string, string, error) {
	user, err := s.repo.SignIn(username, s.HashPassword(password))
	if err != nil {
		return 0, "", "", "", err
	}
	accessToken, err := s.newToken(jwt.MapClaims{"user_id": user.Id, "exp": time.Now().Add(accessTokenTTL).Unix()})
	if err != nil {
		return 0, "", "", "", err
	}
	refreshToken, er := s.newToken(jwt.MapClaims{"user_id": user.Id, "exp": time.Now().Add(refreshTokenTTL).Unix()})
	if er != nil {
		return 0, "", "", "", er
	}
	return user.Id, user.Name, accessToken, refreshToken, nil
}
func (s *AuthService) newToken(tokenStruct jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenStruct)
	tokenS, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenS, nil
}
func (s *AuthService) Refresh(token string) (string, string, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", "", err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return "", "", errors.New("token is invalid")
		}
		accessToken, err := s.newToken(jwt.MapClaims{"user_id": claims["user_id"], "exp": time.Now().Add(accessTokenTTL).Unix()})
		if err != nil {
			return "", "", err
		}
		refreshToken, er := s.newToken(jwt.MapClaims{"user_id": claims["user_id"], "exp": time.Now().Add(refreshTokenTTL).Unix()})
		if er != nil {
			return "", "", er
		}
		return accessToken, refreshToken, nil
	}
	return "", "", errors.New("invalid token")
}
func (s *AuthService) GetUserById(id int) (string, string, error) {
	return s.repo.GetUserById(id)
}
func (s *AuthService) PutUser(username string, name string, id int) error {
	return s.repo.PutUser(username, name, id)
}
