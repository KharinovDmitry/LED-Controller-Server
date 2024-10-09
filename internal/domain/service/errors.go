package service

import "fmt"

var (
	ErrUserAlreadyExist   = fmt.Errorf("пользователь с таким логином уже существует")
	ErrInvalidCredentials = fmt.Errorf("неверные логин или пароль")
	ErrTokenExpired       = fmt.Errorf("истекло действие токена")
)
