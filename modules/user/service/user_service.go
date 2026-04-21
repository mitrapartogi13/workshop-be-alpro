package service

import (
    "github.com/Mobilizes/materi-be-alpro/database/entities"
    "github.com/Mobilizes/materi-be-alpro/modules/user/dto"
    "github.com/Mobilizes/materi-be-alpro/modules/user/repository"
    "github.com/Mobilizes/materi-be-alpro/pkg/helpers"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*entities.User, error) {
    hashedPassword, err := helpers.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := &entities.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
    }

    err = s.repo.Create(user)
    return user, err
}

// ambil data dari repository berdasarkan id
func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    return user, nil
}
