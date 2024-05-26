package handler

import (
	"github.com/yekuanyshev/aiplus/internal/service"
)

type Manager struct {
	City     *city
	Employee *employee
}

func NewManager(service service.Manager) *Manager {
	return &Manager{
		City:     NewCity(service),
		Employee: NewEmployee(service),
	}
}
