package main

import (
		"todo-api/config"
			"todo-api/controllers"
				"todo-api/middleware"

					"github.com/gin-gonic/gin"
				)

				func main() {
						// Initialize Database & Cache
							config.ConnectDB()
								config.ConnectRedis()

									r := gin.Default()

										authCtrl := new(controllers.AuthController)
											taskCtrl := new(controllers.TaskController)

												// Auth Route
													r.POST("/login", authCtrl.Login)

														// Task Routes (Protected with JWT)
															protected := r.Group("/")
																protected.Use(middleware.JWTAuth())
																	{
																				protected.POST("/tasks", taskCtrl.CreateTask)
																						protected.GET("/tasks", taskCtrl.GetTasks)
																								protected.GET("/tasks/:id", taskCtrl.GetTaskByID)
																										protected.PUT("/tasks/:id", taskCtrl.UpdateTask)
																												protected.DELETE("/tasks/:id", taskCtrl.DeleteTask)
																													}

																														r.Run(":8080")
																													}
