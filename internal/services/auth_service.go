package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"auth-service-backend/internal/interfaces"
	"auth-service-backend/internal/models"
	"auth-service-backend/internal/utils"
)

type AuthService struct {
	repoCode   interfaces.CodeRepository
	repoUser   interfaces.UserRepository
	repoRole   interfaces.RoleRepository
	mailModule interfaces.MailModule
	jwtModule  interfaces.JWTModule
}

func NewAuthService(
	repoCode interfaces.CodeRepository,
	repoUser interfaces.UserRepository,
	repoRole interfaces.RoleRepository,
	mailModule interfaces.MailModule,
	jwtModule interfaces.JWTModule,
) *AuthService {
	return &AuthService{
		repoCode:   repoCode,
		repoUser:   repoUser,
		repoRole:   repoRole,
		mailModule: mailModule,
		jwtModule:  jwtModule,
	}
}

func (s *AuthService) SendMail(ctx context.Context, dto *models.CreateMailDto) error {
	code := int16(rand.Intn(9000) + 1000)
	err := s.repoCode.Create(ctx, &models.Code{Code: &code, Email: dto.Email})

	if err != nil {
		return err
	}

	subject := "Your code: " + fmt.Sprint(code)
	body := "<h1>Hello!</h1><p>You have received the authorization code</p>"

	err = s.mailModule.SendEmail(dto.Email, subject, body)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Register(ctx context.Context, dto *models.CreateUserDto) (string, error) {
	email, err := s.repoCode.GetByCode(ctx, dto.Code)
	if err != nil {
		return "", err
	}

	if email != dto.Email {
		return "", errors.New("registration email is not valid")
	}

	err = s.repoCode.Delete(ctx, dto.Code)
	if err != nil {
		return "", err
	}

	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:    dto.Email,
		FullName: dto.FullName,
		Password: hashedPassword,
	}

	userRow, err := s.repoUser.Create(ctx, &user)
	if err != nil {
		return "", err
	}

	role, err := s.repoRole.GetRoleByName(ctx, "admin")
	if err != nil {
		return "", err
	}

	err = s.repoRole.CreateUserRole(ctx, &models.UserRole{
		User_ID: userRow.ID,
		Role_ID: role.ID,
	})
	if err != nil {
		return "", err
	}

	token, err := s.jwtModule.GenerateToken(userRow.ID.String(), role.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Login(ctx context.Context, dto *models.LoginDto) (string, error) {
	user, err := s.repoUser.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}

	ok := utils.CheckPasswordHash(dto.Password, user.Password)
	if !ok {
		return "", errors.New("password don't match")
	}

	role := "user"
	if user.RoleName != "" {
		role = user.RoleName
	}

	token, err := s.jwtModule.GenerateToken(user.ID.String(), role)
	if err != nil {
		return "", err
	}

	return token, nil
}
