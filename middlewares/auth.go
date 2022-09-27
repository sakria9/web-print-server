package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const UserKey = "user"

var secret = []byte("secret")

func AuthCookieStore() gin.HandlerFunc {
	store := cookie.NewStore([]byte("secret"))
	return sessions.Sessions("sessionid", store)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(UserKey)
		if user == nil {
			c.AbortWithStatus(401)
		} else {
			c.Next()
		}
	}
}
