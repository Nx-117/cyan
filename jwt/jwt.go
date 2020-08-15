package jwt

/**
jwt相关
*/

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

/**
签发人
*/
var Issuer = "Cyan"

/**
令牌有效时间
*/
var ExpireTime = 3 * time.Hour

/**
密钥
*/
var jwtSecret = []byte("")

/**
生成token的实体类
参数可以自定义
*/
type Claims struct {
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
	ID       interface{} `json:"id"`
	jwt.StandardClaims
}

/**
创建token
*/
func GenerateToken(username, password, ID interface{}) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(ExpireTime)

	claims := Claims{
		username,
		password,
		ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 失效时间
			Issuer:    Issuer,            // 签发人

		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

/**
解析token
*/
func ParseToken(token string) (*Claims, bool) {
	tokenClaims, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, true
		}
	}

	return nil, false
}
