package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Mengambil nilai session token dari cookie
	sessionToken, err := ctx.Cookie("session_token")
	if err != nil {
		if ctx.GetHeader("Content-Type") == "application/json" {
			ctx.AbortWithStatusJSON(401, model.ErrorResponse{Error: err.Error()})
			return
		} else {
			ctx.Redirect(http.StatusSeeOther, "/user/login") // Ganti dengan URL halaman login yang sesuai
		}
		return
	}

	// Verifikasi token JWT
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(sessionToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(model.JwtKey), nil // Ganti dengan secret key yang sesuai
	})
	if err != nil || !token.Valid {
		ctx.AbortWithStatus(400)
		return
	}

	// Menyimpan user ID di dalam konteks
	ctx.Set("email", claims.Email)

	// Lanjut ke middleware atau handler berikutnya
	ctx.Next()
	})
}
