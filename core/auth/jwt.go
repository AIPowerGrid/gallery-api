package auth

import (
	"errors"
	"gallery/api"
	"gallery/core"
	"time"

	"gallery/db"
	"gallery/models"
	"gallery/settings"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

// JWTAuthenticationBackend hi
type JWTAuthenticationBackend struct {
	// privateKey *rsa.PrivateKey
	// PublicKey  *rsa.PublicKey
	Secret string
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var (
	appSettings = settings.Get()
	log         = core.GetLogger()
)

var authBackendInstance *JWTAuthenticationBackend = nil

func GenPassword(password string) (string, error) {

	p, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(p), err
}

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			Secret: appSettings.JWTSecret,
			// Secret: "mysupersecret101",
			// privateKey: getPrivateKey(),
			// PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

// GenerateToken returns the signed string we send to user;
func (backend *JWTAuthenticationBackend) GenerateToken(email, id string) (string, error) {
	// token := jwt.New(jwt.SigningMethodRS512)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(appSettings.JWTExpirationDelta)).Unix(),
		"iat": time.Now().Unix(),
		"sub": email,
		"id":  id,
	}
	tokenString, err := token.SignedString([]byte(appSettings.JWTSecret))
	if err != nil {
		// panic(err)
		return "", err
	}
	return tokenString, nil
}

// Authenticate  find user , compare password, return true / false and user object
func (backend *JWTAuthenticationBackend) Authenticate(payload api.LoginPayload) (models.User, error) {

	user, err := db.UserByEmail(payload.Email)
	if err != nil {
		log.Error(err)
		return models.User{}, err
	}
	pwMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) == nil
	match := user.Email == payload.Email && pwMatch
	if match {
		return user, nil
	}
	return models.User{}, errors.New("password mismatch")
	// return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil;
}

func (backend *JWTAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

// Logout just is fake for now lol
func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	return nil

}

/*func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	redisConn := redis.Connect()
	return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims.(jwt.MapClaims)["exp"]))
}

func (backend *JWTAuthenticationBackend) IsInBlacklist(token string) bool {
	redisConn := redis.Connect()
	redisToken, _ := redisConn.GetValue(token)

	if redisToken == nil {
		return false
	}

	return true
}*/

/*func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}*/

/*func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}*/
