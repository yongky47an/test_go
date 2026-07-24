package main

import (
		"bytes"
			"encoding/json"
				"net/http"
					"net/http/httptest"
						"testing"
							"todo-api/config"
								"todo-api/controllers"

									"github.com/gin-gonic/gin"
								)

								func setupRouter() *gin.Engine {
										gin.SetMode(gin.TestMode)
											r := gin.Default()
												taskCtrl := new(controllers.TaskController)

													r.POST("/tasks", taskCtrl.CreateTask)
														return r
													}

													func TestCreateTaskValidation(t *testing.T) {
															router := setupRouter()

																// Testing payload tanpa field "title" (harus return status 400 Bad Request)
																	invalidPayload := map[string]interface{}{
																				"description": "Tidak ada title",
																						"status":      "pending",
																								"due_date":    "2026-12-31",
																									}
																										body, _ := json.Marshal(invalidPayload)

																											req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
																												req.Header.Set("Content-Type", "application/json")

																													w := httptest.NewRecorder()
																														router.ServeHTTP(w, req)

																															if w.Code != http.StatusBadRequest {
																																		t.Errorf("Ekspektasi status 400, didapat %d", w.Code)
																																			}
																																		}
