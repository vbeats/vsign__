package applestore

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
	"vsign/logger"
)

type Jwt struct {
	Iss string `json:"iss" binding:"required"`
	Kid string `json:"kid" binding:"required"`
	P8  string `json:"p8" binding:"required"`
}

func (j Jwt) GenerateToken() (string, error) {

	if j.Iss == "" || j.Kid == "" {
		logger.Error("iss or kid is null")
		return "", errors.New("iss or kid is null")
	}

	payload := jwt.MapClaims{
		"aud": "appstoreconnect-v1",
		"exp": time.Now().Add(time.Duration(20) * time.Minute).Unix(),
		"iat": time.Now().Unix(),
		"iss": j.Iss,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodES256, payload)

	claims.Header["kid"] = j.Kid

	// Sign and get the complete encoded token as a string using the p8 private key
	privateKey := j.P8

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		logger.Error("p8 private key decode error.........")
		return "", errors.New("p8 file error")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		logger.Error("p8 private key parse error..... %s", err.Error())
		return "", errors.New("p8 file error")
	}

	if _, ok := key.(*ecdsa.PrivateKey); !ok {
		logger.Error("p8 private key type error.........")
		return "", errors.New("p8 file error")
	}

	tokenStr, err := claims.SignedString(key)
	if err != nil {
		logger.Error("jwt生成异常: %s", err)
		return "", err
	}
	return tokenStr, nil
}

func HandleJwt(s *gin.Engine) {
	s.POST("/asca/jwt", func(c *gin.Context) {
		param := Jwt{}
		if err := c.ShouldBind(&param); err != nil {
			logger.Error("参数错误：%s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		} else {
			token, err := Jwt{
				Iss: param.Iss,
				Kid: param.Kid,
				P8:  param.P8,
			}.GenerateToken()
			if err != nil {
				logger.Error("生成token失败：%s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "success",
				"data":    token,
			})
		}
	})
}
