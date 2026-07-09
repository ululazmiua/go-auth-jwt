package middleware

import (
	"GO-AUTH-JWT/exception"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(app *fiber.App) {
	app.Use("/user", func(ctx fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")                             // ? ctx.Get(key string) string, digunakan untuk mengambil nilai header dengan key tertentu, jika header tidak ditemukan maka akan mengembalikan string kosong
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") { // ? strings.HasPrefix(str string, prefix string) bool, digunakan untuk memeriksa apakah string awal memiliki prefix tertentu
			panic(exception.NewUnauthorizedError("Akses ditolak, token tidak ditemukan!"))

		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")             // ? strings.TrimPrefix(str string, prefix string) string, digunakan untuk menghapus prefix dari string
		token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) { // ? jwt.Parse(token string, keyFunc func(t *jwt.Token) (interface{}, error)) (*Token, error), digunakan untuk memparse token JWT, jika token valid maka akan mengembalikan token dan nil, jika tidak maka akan mengembalikan nil dan error
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // ? token.Claims.(type) => digunakan untuk mengambil claims dari token, jika claims bertipe jwt.MapClaims maka akan mengembalikan claims, jika tidak maka akan mengembalikan nil, token.Valid => digunakan untuk memeriksa apakah token valid atau tidak, jika token valid maka akan mengembalikan true, jika tidak maka akan mengembalikan false
			ctx.Locals("userId", claims["sub"]) // ? ctx.Locals(key string, value any) => digunakan untuk menyimpan data lokal pada context, data ini bisa diakses oleh middleware atau handler lainnya
			return ctx.Next()	// ? ctx.Next() => digunakan untuk menjalankan handler berikutnya
		} else {
			panic(exception.NewUnauthorizedError("Token tidak valid"))
		}

	})
}
