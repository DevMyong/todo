package user

import (
	"github.com/devmyong/todo/pkg/config"
	"github.com/devmyong/todo/pkg/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type RegisterService struct {
	cfg    *config.UserConfig
	logger *zap.Logger

	repo        Repository
	oAuthClient OAuthClient
	pub         EventPublisher
}

func NewRegisterService(cfg *config.UserConfig, logger *zap.Logger, repo Repository, oAuth OAuthClient, pub EventPublisher) *RegisterService {
	logger = logger.With(zap.String("service", "userRegister"))
	return &RegisterService{
		cfg:         cfg,
		logger:      logger,
		repo:        repo,
		oAuthClient: oAuth,
		pub:         pub,
	}
}

func (s *RegisterService) RegisterLocal(email, password, name string) (*User, error) {
	s.logger.Info("Registering a new user", zap.String("email", email))
	return nil, nil
	if existing, _ := s.repo.GetByEmail(email); existing != nil {
		return nil, ErrorUserEmailAlreadyExists
	}
	hash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:           primitive.NewObjectID(),
		Email:        email,
		Name:         name,
		Provider:     ProviderLocal,
		PasswordHash: hash,
	}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *RegisterService) RegisterOAuth(token string) (*User, error) {
	email, err := s.oAuthClient.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	if existing, _ := s.repo.GetByEmail(email); existing != nil {
		return nil, ErrorUserEmailAlreadyExists
	}

	u := &User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Provider: ProviderGoogle,
	}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}
