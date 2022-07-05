package middleware

import (
	"net/http"
	"strconv"
	"strings"
	jwt "todolist/auth"

	"github.com/labstack/echo/v4"
)

// get token function
func getToken(c echo.Context) string {
	authorization := c.Request().Header.Get("Authorization")
	tokenReq := strings.Replace(authorization, "Bearer ", "", -1)
	return tokenReq
}

// Check if the user is logged in
func CheckAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// continue your code here
		tokenReq := getToken(c)

		err := jwt.VerifyJWT(tokenReq)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		return next(c)
	}
}

// Check if the user has permission to access this api
func CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// continue your code here
		tokenReq := getToken(c)
		// get user info from jwt
		tokenUserID, err := jwt.GetUserInfoFromJWT(tokenReq)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		// get userID from user's request
		reqUserID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if tokenUserID != uint(reqUserID) {
			return c.JSON(http.StatusUnauthorized, "You do not have permission to access this api")
		}
		return next(c)
	}
}
