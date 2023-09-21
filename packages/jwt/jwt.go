package jwt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// nuboJwt "github.com/nubo/jwt"
)

type JwtCtrl struct{}

// /// แกะ token
// func (j jwtCtrl) VerifyToken(rawToken string) string {
// 	var payload interface{}
// 	var careId string
// 	// rawToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJjYXJlSWQiOiJQb1RBWHRRQldnIn0.998AX3mvnYYfUQf6AZi12M1AmHnVUyQu5u9PlWOpHtM"

// 	token, ok := nuboJwt.ParseAndVerify(rawToken, "blueposh")
// 	if !ok {
// 		log.Fatal("Invalid token")
// 	} else {
// 		payload = token.ClaimSet["careId"]
// 		careId := fmt.Sprintf("%v", payload)
// 		// fmt.Println("Type", token.Header.Type)
// 		// fmt.Println("Algorithm", token.Header.Algorithm)
// 		// fmt.Println("Claim Set", token.ClaimSet)

// 		return careId
// 	}

// 	return careId
// }

// func (j jwtCtrl) CreateToken(careId string) (string, error) {
// 	var err error

// 	secret := "blueposh"
// 	atClaims := jwt.MapClaims{}
// 	// atClaims["careId"] = "PoRcipekGZ"
// 	atClaims["careId"] = careId

// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	token, err := at.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }

// jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	// atClaims := jwt.MapClaims{}
	// atClaims["authorized"] = true
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

func (j JwtCtrl) ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "jdnfksdmfksd" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// check token signing method etc
		return hmacSecret, nil
	})
	fmt.Println("Parse ] ]", err)
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
