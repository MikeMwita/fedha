package usecase

import (
	"context"
	"errors"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/util"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrUserExists          = errors.New("user already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal serve error")
	ErrUnauthorizedUser    = errors.New("unauthorized user")
	ErrUnaauthorizedToken  = errors.New("unauthorized token")
)

type AuthUsecase struct {
	authService adapters.AuthService
	authRepo    adapters.AuthRepository
	cfg         *config.Config
	tracer      trace.Tracer
}

func (u *AuthUsecase) Register(ctx context.Context, user *entity.User) (*entity.UserWithToken, error) {
	ctx, span := u.tracer.Start(ctx, "auth.Register")
	defer span.End()
	existsUser, err := u.authRepo.FindByEmail(ctx, user)

	if existsUser != nil || err == nil {
		return nil, ErrUserExists
	}

	if err = user.PrepareCreate(); err != nil {
		return nil, ErrBadRequest
	}

	registerRequest := dto.RegisterRequest{
		Username: user.UserName,
		Password: user.Password,
	}
	createdUser, err := u.authRepo.Register(ctx, registerRequest)
	if err != nil {
		return nil, ErrInternalServerError
	}

	token, err := util.GenerateJWTToken(createdUser.User, u.cfg)
	if err != nil {
		return nil, ErrInternalServerError
	}

	return &entity.UserWithToken{
		User:  createdUser.User,
		Token: token,
	}, nil
}

func (u *AuthUsecase) Login(ctx context.Context, user *entity.User) (*entity.UserWithToken, error) {
	ctx, span := u.tracer.Start(ctx, "auth.Register")
	defer span.End()
	foundUser, err := u.authRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, err
	}

	if err = foundUser.ComparePasswords(user.Password); err != nil {
		return nil, ErrUnauthorizedUser
	}

	foundUser.SanitizePassword()

	token, err := util.GenerateJWTToken(foundUser, u.cfg)
	if err != nil {
		return nil, ErrUnaauthorizedToken
	}

	return &entity.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}

func (u *AuthUsecase) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	ctx, span := u.tracer.Start(ctx, "auth.Register")
	defer span.End()

	updateUserReq := &db.UpdateUserReq{
		UserName:     user.UserName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: user.PasswordHash,
	}
	updatedUserRes, err := u.authRepo.Update(ctx, updateUserReq)
	if err != nil {
		return nil, err
	}
	updatedUser := &entity.User{
		UserName:    updatedUserRes.UserName,
		Email:       updatedUserRes.Email,
		PhoneNumber: updatedUserRes.PhoneNumber,
	}

	return updatedUser, nil

}

func (u *AuthUsecase) Delete(ctx context.Context, userID uuid.UUID) error {
	ctx, span := u.tracer.Start(ctx, "auth.Register")
	defer span.End()
	if err := u.authRepo.Delete(ctx, userID); err != nil {
		return err
	}
	return nil
}

func (u *AuthUsecase) GetByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	ctx, span := u.tracer.Start(ctx, "auth.Register")
	defer span.End()
	regUserRes, err := u.authRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		UserName:    regUserRes.UserName,
		Email:       regUserRes.Email,
		PhoneNumber: regUserRes.PhoneNumber,
	}

	return user, nil
}

func NewAuthUsecase(cfg *config.Config, authService adapters.AuthService, authRepo adapters.AuthRepository, tracer trace.Tracer) adapters.AuthUseCase {
	return &AuthUsecase{
		cfg:         cfg,
		authService: authService,
		authRepo:    authRepo,
		tracer:      tracer,
	}
}
