package middleware

import (
		"net/http"
			"strings"

				"github.com/gin-gonic/gin"
					"github.com/golang-jwt/jwt/v5"
				)

				var JWTSecret = []byte("super-secret-key")

				func JWTAuth() gin.HandlerFunc {
						return func(c *gin.Context) {
									authHeader := c.GetHeader("Authorization")
											if authHeader == "" {
															c.JSON(http.StatusUnauthorized, gin.H{"error": "Header Otorisasi Diperlukan"})
																		c.Abort()
																					return
																							}

																									tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
																											token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
																															return JWTSecret, nil
																																	})

																																			if err != nil || !token.Valid {
																																							c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau expired"})
																																										c.Abort()
																																													return
																																															}

																																																	c.Next()
																																																		}
																																																	}
