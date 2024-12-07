package user_repository

import (
	auth_user_model "backend/src/models/user"
	logging_utils "backend/src/utils/logging"
)

// RepositoryGetUserByEmail はメールアドレスをキーにユーザーを取得
func (r *userRepository) RepositoryGetUserByEmail(email string) (*auth_user_model.User, error) {
	start := logging_utils.LogStart()

	var user auth_user_model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogInfo("RepositoryGetUserByEmail user:", user)
	return &user, nil
}

// RepositoryGetUserByID はユーザーを取得
func (r *userRepository) RepositoryGetUserByID(userID string) (*auth_user_model.User, error) {
	start := logging_utils.LogStart()

	var user auth_user_model.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		logging_utils.LogError(start, err)
		return nil, err
	}

	logging_utils.LogInfo("RepositoryGetUserByID user:", user)
	return &user, nil
}

// RepositoryCreateUser はユーザーを作成
func (r *userRepository) RepositoryCreateUser(user *auth_user_model.User) error {
	start := logging_utils.LogStart()

	// トランザクション開始
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ユーザーを作成
	if err := tx.Create(user).Error; err != nil {
		logging_utils.LogError(start, err)
		tx.Rollback()
		return err
	}

	// トランザクション終了
	if err := tx.Commit().Error; err != nil {
		logging_utils.LogError(start, err)
		return err
	}

	logging_utils.LogInfo("RepositoryCreateUser end.")
	return nil
}
