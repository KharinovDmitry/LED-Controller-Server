package service

import "fmt"

var (
	ErrUserAlreadyExist   = fmt.Errorf("пользователь с таким логином уже существует")
	ErrInvalidCredentials = fmt.Errorf("неверные логин или пароль")
	ErrTokenExpired       = fmt.Errorf("истекло действие токена")

	ErrPanelNotFound      = fmt.Errorf("панель не найдена, проверьте ключ или перезагрузите панель")
	ErrPanelNotRegistered = fmt.Errorf("панель не зарегестрирована")

	ErrCacheUpdate = fmt.Errorf("cache update error")
)
