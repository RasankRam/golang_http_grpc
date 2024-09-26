package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mdobak/go-xerrors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"todo-list/internal/app/repository"
	"todo-list/internal/utils"
)

type AuthService struct {
	userRepo  *repository.UserRepo
	roleRepo  *repository.RoleRepo
	tokenRepo *repository.TokenRepo
	logger    *slog.Logger
}

func NewAuthService(userRepo *repository.UserRepo, roleRepo *repository.RoleRepo, tokenRepo *repository.TokenRepo, logger *slog.Logger) *AuthService {
	return &AuthService{userRepo: userRepo, roleRepo: roleRepo, tokenRepo: tokenRepo, logger: logger}
}

func (s *AuthService) Login(ctx context.Context, login string, pass string, ip string, userAgent string) (map[string]string, error) {
	dbUser, err := s.userRepo.UserByLogin(login)

	fmt.Println("dbUser", dbUser)

	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return nil, errors.New("user not found")
	}

	hashPass := dbUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return nil, errors.New("failed to process pass")
	}

	token, err := utils.GenerateAuthToken(login, dbUser.Role)
	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return nil, errors.New("failed to generate auth token")
	}

	_, err = s.tokenRepo.CreateToken(token, ip, dbUser.Id, userAgent)
	if err != nil {
		var outputErr error
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			outputErr = errors.New("you cannot log in twice")
		} else {
			outputErr = errors.New("failed to create token in database")
		}
		utils.LogErrorContext(s.logger, ctx, err)
		return nil, outputErr
	}

	// Return the tokens as JSON
	tokenResponse := map[string]string{
		"token": token,
	}

	return tokenResponse, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) (string, error) {
	deletedToken, err := s.tokenRepo.DeleteToken(token)

	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return "", errors.New("failed to invalidate token")
	}

	return deletedToken, err
}

func (s *AuthService) Register(ctx context.Context, login string, pass string, role string) (int, error) {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return 0, errors.New("invalid credentials")
	}

	roleEntity, err := s.roleRepo.RoleByNm(role)
	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return 0, errors.New("failed to fetch role")
	}
	if roleEntity == nil {
		err = xerrors.New(err)
		s.logger.ErrorContext(ctx, "", slog.Any("error", err))
		return 0, errors.New("failed to find role")
	}

	userId, err := s.userRepo.CreateUserRegister(login, string(hashedPassword), role)

	if err != nil {
		utils.LogErrorContext(s.logger, ctx, err)
		return 0, errors.New("internal error")
	}

	return userId, nil
}
