package controllers

import (
		"net/http"
			"time"

				"todo-api/middleware"

					"github.com/gin-gonic/gin"
						"github.com/golang-jwt/jwt/v5"
					)

					type AuthController struct{}

					func (a *AuthController) Login(c *gin.Context) {
							var body struct {
										Username string `json:"username" binding:"required"`
												Password string `json:"password" binding:"required"`
													}

														if err := c.ShouldBindJSON(&body); err != nil {
																	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
																			return
																				}

																					// Mock Authentication (Dapat diganti DB check)
																						if body.Username == "admin" && body.Password == "password" {
																									token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
																													"username": body.Username,
																																"exp":      time.Now().Add(time.Hour * 24).Unix(),
																																		})

																																				tokenString, _ := token.SignedString(middleware.JWTSecret)
																																						c.JSON(http.StatusOK, gin.H{"token": tokenString})
																																								return
																																									}

																																										c.JSON(http.StatusUnauthorized, gin.H{"error": "Kredensial tidak valid"})
																																									}
