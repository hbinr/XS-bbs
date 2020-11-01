package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"xs.bbs/internal/pkg/constant/key"
)

type MyClaims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 24

// GenToken 生成token
func GenToken(userID int64) (string, error) {

	// 创建一个我们自己的声明
	c := &MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "XS-bbs",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名，注意转换为字节切片，并获得完整的编码后的字符串token
	return token.SignedString([]byte(key.MySecret))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var (
		token *jwt.Token
	)
	// 这行分配内存地址的代码一定要写，否则在赋值时会提示invalid vlaue
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key.MySecret), nil
	})
	if err != nil {
		return
	}
	if token.Valid { // 校验token
		return
	}
	return nil, errors.New("invalid token")
}
