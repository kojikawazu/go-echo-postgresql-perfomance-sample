package auth_handler

import (
	"net/http"

	jwt_utils "backend/src/utils/jwt"
	logging_utils "backend/src/utils/logging"

	"github.com/labstack/echo/v4"
)

// サインインハンドラー
func (h *AuthUserHandler) SignIn(c echo.Context) error {
	start := logging_utils.LogStart()

	// リクエストボディ
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// リクエストバインド
	if err := c.Bind(&req); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// ユーザー認証
	token, err := h.service.ServiceAuthenticateUser(req.Email, req.Password)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	// クッキーにトークンを設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	// レスポンス
	logging_utils.LogInfo("AuthUserHandler SignIn end.")
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, map[string]string{"message": "sign-in successful"})
}

// サインアップハンドラー
func (h *AuthUserHandler) SignUp(c echo.Context) error {
	start := logging_utils.LogStart()

	// リクエストバインド
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	// リクエストバインド
	if err := c.Bind(&req); err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// ユーザー登録
	token, err := h.service.ServiceRegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// クッキーにトークンを設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	// 完了
	logging_utils.LogInfo("AuthUserHandler SignUp end.")
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, map[string]string{"message": "sign-up successful"})
}

// サインアウトハンドラー
func (h *AuthUserHandler) SignOut(c echo.Context) error {
	start := logging_utils.LogStart()

	// トークンクッキーを削除
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.MaxAge = -1 // クッキーを即座に無効化
	c.SetCookie(cookie)

	logging_utils.LogInfo("AuthUserHandler SignOut end.")
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, map[string]string{"message": "sign-out successful"})
}

// ユーザー情報取得ハンドラー
func (h *AuthUserHandler) GetAuthenticatedUser(c echo.Context) error {
	start := logging_utils.LogStart()

	// クッキーからトークンを取得
	cookie, err := c.Cookie("token")
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
	}

	// トークンを検証
	claims, err := jwt_utils.ValidateToken(cookie.Value)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
	}

	// トークンから取得したユーザーID
	userID := claims.UserID
	logging_utils.LogInfo("AuthUserHandler GetAuthenticatedUser userID:", userID)

	// ユーザー情報を取得
	user, err := h.service.ServiceGetUserByID(userID)
	if err != nil {
		logging_utils.LogError(start, err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	// パスワードなどのセンシティブな情報を削除
	user.Password = ""

	logging_utils.LogInfo("AuthUserHandler GetAuthenticatedUser end.")
	logging_utils.LogEnd(start)
	return c.JSON(http.StatusOK, user)
}
