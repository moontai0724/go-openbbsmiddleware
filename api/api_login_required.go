package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginRequiredJSON(theFunc LoginRequiredApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func LoginRequiredQuery(theFunc LoginRequiredApiFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err)
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func loginRequiredProcess(theFunc LoginRequiredApiFunc, params interface{}, c *gin.Context) {

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 403, ErrInvalidRemoteAddr)
		return
	}

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin)
		return
	}

	userID, err := verifyJwt(c)
	if err != nil {
		processResult(c, nil, 401, err)
	}

	result, statusCode, err := theFunc(remoteAddr, userID, params, c)
	processResult(c, result, statusCode, err)
}
