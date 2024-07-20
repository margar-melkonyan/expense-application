package service

import (
	"encoding/json"
	"errors"
	"expense-application/internal/model"
	"expense-application/internal/repository"
	"expense-application/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	cfg                     = config.MustLoad()
	accessTokenKey          []byte
	refreshTokenKey         []byte
	accessTokenDuration, _  = time.ParseDuration(cfg.Auth.AccessTokenTTL)
	refreshTokenDuration, _ = time.ParseDuration(config.MustLoad().Auth.RefreshTokenTTL)
)

type Token struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthService struct {
	repository repository.User
}

func NewAuthService(repository repository.User) *AuthService {
	return &AuthService{repository: repository}
}

func initTokenKeys() {
	accessTokenKey = []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
	refreshTokenKey = []byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))
}

func generationToken(userID uint, tokenKey []byte, duration time.Duration) (string, error) {
	sub, _ := json.Marshal(map[string]any{
		"id": userID,
	})

	payload := jwt.MapClaims{
		"sub": string(sub),
		"exp": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString(tokenKey)
}

func (s *AuthService) updateRefreshToken(refreshToken string, user *model.User) error {
	user.RefreshToken, _ = json.Marshal(map[string]string{
		"refresh_token": refreshToken,
	})

	return s.repository.Update(user)
}

func (s *AuthService) SignUp(user *model.User) (map[string]string, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err := bcrypt.CompareHashAndPassword(hash, []byte(user.PasswordConfirmation)); err != nil {
		return map[string]string{}, errors.New("password doesn't match")
	}

	userID, err := s.repository.Create(user)
	if err != nil {
		return map[string]string{}, errors.New("email not unique")
	}

	initTokenKeys()
	accessToken, err := generationToken(userID, accessTokenKey, accessTokenDuration)
	refreshToken, err := generationToken(userID, refreshTokenKey, refreshTokenDuration)

	if err != nil {
		return map[string]string{}, errors.New("error creating token")
	}

	if err = s.updateRefreshToken(refreshToken, user); err != nil {
		slog.Error(err.Error())
	}

	return map[string]string{
		"type":          "Bearer",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (s *AuthService) SignIn(user *model.User) (map[string]string, error) {
	currentUser, err := s.repository.GetByEmail(user.Email)
	if err != nil {
		return map[string]string{}, errors.New("this user doesn't exist")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(user.Password)); err != nil {
		return map[string]string{}, errors.New("password doesn't match")
	}

	initTokenKeys()
	accessToken, err := generationToken(currentUser.Id, accessTokenKey, accessTokenDuration)
	refreshToken, err := generationToken(currentUser.Id, refreshTokenKey, refreshTokenDuration)

	if err != nil {
		return map[string]string{}, errors.New("error creating token")
	}

	if err = s.updateRefreshToken(refreshToken, &currentUser); err != nil {
		slog.Error(err.Error())
	}

	return map[string]string{
		"type":          "Bearer",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil
}

func (s *AuthService) RefreshToken(ctx *gin.Context) {
	var currentUser model.User
	authorization := ctx.GetHeader("Authorization")
	initTokenKeys()

	if authorization == "" {
		ctx.JSON(http.StatusConflict, errors.New("authorization header missing"))
		return
	}

	refreshToken := strings.Split(authorization, "Bearer ")[1]
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return refreshTokenKey, nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, errors.New("error parsing refresh token"))
		return
	}

	if claims, ok := parsedToken.Claims.(jwt.Claims); ok && parsedToken.Valid {
		exp, err := claims.GetExpirationTime()
		if err != nil {
			slog.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{})
			return
		}

		sub, err := claims.GetSubject()
		if err != nil {
			slog.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{})
			return
		}

		err = json.Unmarshal([]byte(sub), &currentUser)
		if err != nil {
			slog.Error(err.Error())
		}
		currentUser, err = s.repository.Get(currentUser.Id)

		var currUserRT Token
		err = json.Unmarshal(currentUser.RefreshToken, &currUserRT)

		if err != nil {
			slog.Error(err.Error())
		}

		if currUserRT.RefreshToken != refreshToken {
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"message": "Token not valid!",
			})
			return
		}

		if err != nil {
			slog.Error(err.Error())
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{})
			return
		}

		if time.Now().Unix() > exp.Time.Unix() {
			if s.updateRefreshToken("", &currentUser) != nil {
				ctx.AbortWithStatusJSON(409, errors.New("couldn't update token"))
			}
			ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
				"message": "Token was expired!",
			})
			return
		}

		ctx.Set("user", currentUser)
	}

	accessToken, err := generationToken(currentUser.Id, accessTokenKey, accessTokenDuration)
	refreshToken, err = generationToken(currentUser.Id, refreshTokenKey, refreshTokenDuration)

	if err != nil {
		ctx.AbortWithStatusJSON(409, errors.New("error creating token"))
	}

	if s.updateRefreshToken(refreshToken, &currentUser) != nil {
		ctx.AbortWithStatusJSON(409, errors.New("couldn't update token"))
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"type":          "Bearer",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
