package jwt_utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// jwtKey はJWTの秘密鍵
var jwtKey = []byte("your_secret_key")

// Claims はJWTのクレーム
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// トークンを生成
func GenerateToken(userID string) (string, error) {
	fmt.Println("GenerateToken start...")

	// 1日後に期限切れ
	expirationTime := time.Now().Add(24 * time.Hour)
	// クレームを作成
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("GenerateToken end.")
	return token.SignedString(jwtKey)
}

// トークンを検証
func ValidateToken(tokenStr string) (*Claims, error) {
	fmt.Println("ValidateToken start...")

	// クレームを作成
	claims := &Claims{}
	// トークンを検証
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// アルゴリズムのチェック
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		fmt.Println("ValidateToken parse token.")
		return jwtKey, nil
	})

	// トークンの検証結果をチェック
	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			switch {
			case validationErr.Errors&jwt.ValidationErrorExpired != 0:
				return nil, fmt.Errorf("token expired")
			case validationErr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				return nil, fmt.Errorf("invalid token signature")
			default:
				return nil, fmt.Errorf("token validation failed: %v", err)
			}
		}
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		fmt.Println("ValidateToken token is invalid.")
		return nil, fmt.Errorf("token is invalid")
	}

	// 検証が成功した場合
	fmt.Println("ValidateToken end.")
	return claims, nil
}
