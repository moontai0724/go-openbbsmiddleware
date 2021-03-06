package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func setRequest(path string, params interface{}, jwt string, headers map[string]string) *http.Request {
	jsonStr, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(jsonStr))

	req.Header.Set("Host", "localhost:5678")
	req.Header.Set("X-Forwarded-For", "127.0.0.1:5679")
	if jwt != "" {
		req.Header.Set("Authorization", "bearer "+jwt)
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	return req
}

func getJwt(router *gin.Engine, userID string, passwd string) string {
	jwt, _ := createToken(userID)

	return jwt
}

func createToken(userID string) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: pttbbsapi.JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		return "", err
	}

	cl := &pttbbsapi.JwtClaim{
		UUserID: bbs.UUserID(userID),
		Expire:  jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	raw, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		return "", err
	}

	return raw, nil
}
