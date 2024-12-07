package auth_user_service

import (
	"errors"
	"fmt"

	auth_user_model "backend/src/models/user"
	jwt_utils "backend/src/utils/jwt"
	logging_utils "backend/src/utils/logging"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser はユーザーを認証してJWTトークンを返す
func (s *authUserService) ServiceAuthenticateUser(email, password string) (string, error) {
	start := logging_utils.LogStart()

	// ユーザーを取得
	user, err := s.repo.RepositoryGetUserByEmail(email)
	if err != nil {
		logging_utils.LogError(start, err)
		return "", errors.New("user not found")
	}

	logging_utils.LogInfo("ServiceAuthenticateUser user:", user)

	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logging_utils.LogError(start, err)
		return "", errors.New("invalid credentials")
	}

	logging_utils.LogInfo("ServiceAuthenticateUser password is valid")

	// JWTトークンを生成
	token, err := jwt_utils.GenerateToken(user.ID)
	if err != nil {
		logging_utils.LogError(start, err)
		return "", err
	}

	logging_utils.LogInfo("ServiceAuthenticateUser end token:", token)
	return token, nil
}

// ServiceRegisterUser はユーザーを登録します
func (s *authUserService) ServiceRegisterUser(username, email, password string) (string, error) {
	start := logging_utils.LogStart()

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logging_utils.LogError(start, err)
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	// Eメールからユーザーを取得
	_, err = s.repo.RepositoryGetUserByEmail(email)
	if err == nil {
		logging_utils.LogError(start, errors.New("user already exists"))
		return "", errors.New("user already exists")
	}

	user := &auth_user_model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// ユーザーを作成
	if err := s.repo.RepositoryCreateUser(user); err != nil {
		logging_utils.LogError(start, err)
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	// JWTトークンを生成
	token, err := jwt_utils.GenerateToken(user.ID)
	if err != nil {
		logging_utils.LogError(start, err)
		return "", err
	}

	logging_utils.LogInfo("ServiceRegisterUser end token:", token)
	return token, nil
}

// ServiceGetUserByID はユーザー情報を取得します
func (s *authUserService) ServiceGetUserByID(userID string) (*auth_user_model.User, error) {
	start := logging_utils.LogStart()

	// ユーザーを取得
	user, err := s.repo.RepositoryGetUserByID(userID)
	if err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogInfo("ServiceGetUserByID user:", user)
	return user, nil
}
