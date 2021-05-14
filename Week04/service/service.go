package service

import (
	"errors"
	"strings"
)

// Service constants
const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")
	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

// Service Define a service interface
type Service interface {
	//用户登录
	Login(a, b string) (string, error)

	// HealthCheck check service health status
	HealthCheck() bool
}

//ArithmeticService implement Service interface
type LoginService struct {
}

func (s LoginService) Login(a, b string) (string, error) {
	// test for length overflow
	if len(a)+len(b) > StrMaxSize {
		return "", ErrMaxSize
	}
//数据库查找处理，To-Do
	return a + b, nil
}

// HealthCheck implement Service method
// 用于检查服务的健康状态，这里仅仅返回true。
func (s LoginService) HealthCheck() bool {
	return true
}

// ServiceMiddleware define service middleware
type ServiceMiddleware func(Service) Service

