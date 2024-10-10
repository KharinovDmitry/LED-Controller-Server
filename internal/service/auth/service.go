package auth

import (
	"DynamicLED/config"
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/entity"
	"DynamicLED/internal/domain/repository"
	"DynamicLED/internal/domain/service"
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	jwtSecret string
	salt      string

	accessExpireTime  time.Duration
	refreshExpireTime time.Duration

	user  repository.User
	token repository.Token
}

func New(user repository.User, token repository.Token, config config.AuthConfig) *Service {
	return &Service{
		jwtSecret:         config.JwtSecret,
		salt:              config.Salt,
		refreshExpireTime: config.RefreshExpireTime,
		accessExpireTime:  config.AccessExpireTime,
		user:              user,
		token:             token,
	}
}

func (s *Service) Register(ctx context.Context, login, password string) error {
	_, err := s.user.GetUserByLogin(ctx, login)
	if err == nil {
		return service.ErrUserAlreadyExist
	}
	if !errors.Is(err, repository.ErrNotFound) {
		return fmt.Errorf("[ Auth Service ] register err: %w", err)
	}

	hash := s.hash(password)

	err = s.user.AddUser(ctx, entity.User{
		Login:    login,
		Password: hash,
	})
	if err != nil {
		return fmt.Errorf("[ Auth Service ] register err: %w", err)
	}

	return nil
}

func (s *Service) Login(ctx context.Context, login, password string) (access, refresh string, err error) {
	user, err := s.user.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", "", service.ErrInvalidCredentials
		}
		return "", "", fmt.Errorf("[ Auth Service ] login err: %w", err)
	}

	if user.Password != s.hash(password) {
		return "", "", service.ErrInvalidCredentials
	}

	tokenUUID := uuid.New()

	access, err = s.getToken(tokenUUID, s.accessExpireTime, user.Login, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] login err: %w", err)
	}

	refresh, err = s.getToken(tokenUUID, s.refreshExpireTime, user.Login, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] login err: %w", err)
	}

	if err := s.token.AddRefresh(ctx, login, refresh); err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] login err: %w", err)
	}

	return access, refresh, nil
}

func (s *Service) Refresh(ctx context.Context, oldAccess, oldRefresh string) (access, refresh string, err error) {
	accessClaims, err := s.ParseClaims(ctx, oldAccess)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] refresh err: %w", err)
	}

	refreshClaims, err := s.ParseClaims(ctx, oldRefresh)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] refresh err: %w", err)
	}

	curRefresh, err := s.token.GetRefresh(ctx, accessClaims.Login)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] refresh err: %w", err)
	}

	if accessClaims.UUID != refreshClaims.UUID ||
		curRefresh != oldRefresh {
		return "", "", service.ErrInvalidCredentials
	}

	if refreshClaims.Expire.After(time.Now().UTC()) {
		return "", "", service.ErrTokenExpired
	}

	access, err = s.getToken(refreshClaims.UUID, s.accessExpireTime, refreshClaims.Login, refreshClaims.Role)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] refresh err: %w", err)
	}

	refresh, err = s.getToken(refreshClaims.UUID, s.refreshExpireTime, refreshClaims.Login, refreshClaims.Role)
	if err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] refresh err: %w", err)
	}

	if err := s.token.AddRefresh(ctx, accessClaims.Login, refresh); err != nil {
		return "", "", fmt.Errorf("[ Auth Service ] refresh err: %w", err)
	}

	return access, refresh, nil
}

func (s *Service) ParseClaims(ctx context.Context, token string) (entity.Claims, error) {
	var claims entity.Claims
	_, err := jwt.ParseWithClaims(token, &claims, s.keyFunc)
	if err != nil {
		return entity.Claims{}, fmt.Errorf("[ Auth Service ] parse claims")
	}

	if claims.Expire.After(time.Now().UTC()) {
		return entity.Claims{}, service.ErrTokenExpired
	}

	return claims, nil
}

func (s *Service) hash(password string) string {
	return string(sha512.New().Sum([]byte(password + s.salt)))
}

func (s *Service) keyFunc(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexcepted sign method")
	}

	return s.jwtSecret, nil
}

func (s *Service) getToken(uuid uuid.UUID, expireTime time.Duration, login string, role constant.Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constant.TokenUUID:    uuid,
		constant.LoginClaims:  login,
		constant.RoleClaims:   role,
		constant.ExpireClaims: time.Now().UTC().Add(expireTime),
	})

	signedToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("[ Auth Service ] getToken err: %w", err)
	}

	return signedToken, nil
}
