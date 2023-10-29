package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"simple-ecommerce/configs"
	"simple-ecommerce/responses"
	"strings"
)

func JWTMiddleware(c *fiber.Ctx) error {
	resp := responses.Response{}
	resp.Success = false

	auth_header := c.Get("Authorization")
	if !strings.HasPrefix(auth_header, "Bearer") {
		resp.Message = "No Token is Provided"
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	tokenString := strings.TrimPrefix(auth_header, "Bearer ")

	isSuccess, claims, err := ValidateJWTToken(tokenString)
	if !isSuccess || err != nil {
		resp.Message = err.Error()
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}
	log.Info(claims)
	c.Locals("claims", claims)
	return c.Next()
}

// EasyToken is an Struct to encapsulate username and expires as parameter
type EasyToken struct {
	// Username is the name of the user
	UserID   int64
	Username string
	Role     string
	// Expires is a NumericDate with expiration date
	Expires *jwt.NumericDate
}

var (
	mySigningKey []byte
	issuer       string
)

type MyCustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

//temporary disable for easy setup as example code, not production
/*
var (
	verifyKey    *rsa.PublicKey
	mySigningKey *rsa.PrivateKey
)
*/

func InitJWT() {

	//temporary disable for easy setup as example code, not production
	/*
		verifyBytes, err := ioutil.ReadFile(configs.GetConfigString("jwt.public_key_path"))
		if err != nil {
			log.Fatal(err)
		}

		verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			log.Fatal(err)
		}

		signBytes, err := ioutil.ReadFile(configs.GetConfigString("jwt.private_key_path"))

		if err != nil {
			log.Fatal(err)
		}
	*/

	mySigningKey = []byte(configs.GetConfigString("jwt.signature_key"))
	issuer = configs.GetConfigString("jwt.issuer")
}

func GetJWTSigningKey() []byte {
	return mySigningKey
}

// GetToken is a function that exposes the method to get a simple token for jwt
func (e EasyToken) GetJWTToken() (string, error) {

	// Create the Claims
	claims := MyCustomClaims{
		e.UserID,
		e.Username,
		e.Role,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: e.Expires,
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Info(mySigningKey)
		log.Info(err)
	}

	return tokenString, err
}

// ValidateToken get token strings and return if is valid or not
func ValidateJWTToken(tokenString string) (bool, MyCustomClaims, error) {

	if tokenString == "" {
		return false, MyCustomClaims{}, errors.New("token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if token == nil {
		log.Println(err)
		return false, MyCustomClaims{}, errors.New("Token Can't Be Parsed")
	}

	if token.Valid {
		jsonData, err := json.Marshal(token.Claims)
		if err != nil {
			log.Info(err)
			return false, MyCustomClaims{}, errors.New("Cannot marshal claims")
		}
		var claims MyCustomClaims
		err = json.Unmarshal(jsonData, &claims)
		if err != nil {
			log.Info(err)
			return false, MyCustomClaims{}, errors.New("Cannot UnMarshal claims")
		}
		return true, claims, nil
	} else {
		//"Couldn't handle this token:"
		return false, MyCustomClaims{}, errors.New("Invalid Token or Token is Expired")
	}
}
