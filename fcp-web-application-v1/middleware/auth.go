package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here

		// Mengakses isi Cookie
		cookie, err := ctx.Request.Cookie("session_token")
		// Cek apakan session_token terdapat dalam Cookie
		if err != nil {
			// ketika berusaha mengirim data lewat postman namun belum login
			if ctx.GetHeader("Content-type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "error unauthorized user id"})
			} else {
				// ketika berusaha mengirim data dari browser dan belum login maka akan diarahkan menu login
				ctx.Redirect(http.StatusSeeOther, "login")
			}
			return
		}

		// mengambil data dalam cookie
		claims := &model.Claims{}
		// melakikan persing data kedalam jwt
		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		// melakukan pengecekan jika saat melakukan parsing terdapat error
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// melakukan validasi apakah token belum expired
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			return
		}
		// Menyimpan nilai userID dari claims kedalam contex dengan key id sehingga dapat digunakan pada hendler atau endpoint
		ctx.Set("email", claims.Email)
		// memanggil next untuk melanjutkan request ke hendler
		ctx.Next()
	})
}
