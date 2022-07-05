package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// Decoding JWT to get payload, not verifying JWT
func Decode(JWTToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(JWTToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err.Error() == jwt.ErrInvalidKeyType.Error() {
		return token, nil
	}
	return nil, err
}

// Generate HS256 JWT token
func GenerateHS256JWT(secret string, payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}
	for key, val := range payload {
		claims[key] = val
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	signedToken, err := token.SignedString([]byte(secret))
	return signedToken, err
}

// Create jwt token
func Create(userId uint, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() //Token hết hạn sau 12 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

// Verify JWT func
func VerifyJWT(tokenString string) error {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		return ve
	} else {
		return ve
	}
}

// get user id from jwt
func GetUserInfoFromJWT(JWTToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(JWTToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.UserID, err
	}
	return 0, err
}
