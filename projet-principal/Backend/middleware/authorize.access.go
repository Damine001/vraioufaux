package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// authorize determines if current user has been authorized to take an action on an object.
func Authorize(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Get current user/subject
		sub, existed := ctx.Get("role_name")
		if !existed {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "User hasn't logged in yet"})
			return
		}

		// Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"message": "Failed to load policy from DB"})
			return
		}

		// Casbin enforces policy
		auth, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"message": "Error occurred when authorizing user"})
			return
		}

		if !auth {
			ctx.AbortWithStatusJSON(403, gin.H{"message": "You are not authorized"})
			return
		}
		ctx.Next()
	}
}
