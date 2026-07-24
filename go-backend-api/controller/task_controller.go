package controllers

import (
		"encoding/json"
			"fmt"
				"log"
					"math"
						"net/http"
							"strconv"
								"time"

									"todo-api/config"
										"todo-api/models"
											"todo-api/repository"

												"github.com/gin-gonic/gin"
											)

											type TaskController struct {
													repo repository.TaskRepository
												}

												// POST /tasks
												func (tc *TaskController) CreateTask(c *gin.Context) {
														var task models.Task
															if err := c.ShouldBindJSON(&task); err != nil {
																		log.Printf("[ERROR] Bad Request Create Task: %v", err)
																				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
																						return
																							}

																								if err := config.DB.Create(&task).Error; err != nil {
																											log.Printf("[ERROR] Database Create Task: %v", err)
																													c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat task"})
																															return
																																}

																																	// Invalidate Cache Redis
																																		if config.RDB != nil {
																																					config.RDB.FlushDB(config.Ctx)
																																						}

																																							c.JSON(http.StatusCreated, gin.H{
																																										"message": "Task created successfully",
																																												"task":    task,
																																													})
																																												}

																																												// GET /tasks
																																												func (tc *TaskController) GetTasks(c *gin.Context) {
																																														status := c.Query("status")
																																															search := c.Query("search")
																																																page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
																																																	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

																																																		if page < 1 {
																																																					page = 1
																																																						}
																																																							if limit < 1 {
																																																										limit = 10
																																																											}
																																																												offset := (page - 1) * limit

																																																													cacheKey := fmt.Sprintf("tasks:%s:%s:%d:%d", status, search, page, limit)

																																																														// Cek Redis Cache
																																																															if config.RDB != nil {
																																																																		cachedData, err := config.RDB.Get(config.Ctx, cacheKey).Result()
																																																																				if err == nil {
																																																																								var result gin.H
																																																																											json.Unmarshal([]byte(cachedData), &result)
																																																																														c.JSON(http.StatusOK, result)
																																																																																	return
																																																																																			}
																																																																																				}

																																																																																					// Panggil Concurrency Repository Fetch
																																																																																						tasks, totalTasks, err := tc.repo.GetAllConcurrent(status, search, limit, offset)
																																																																																							if err != nil {
																																																																																										log.Printf("[ERROR] Fetching tasks: %v", err)
																																																																																												c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil tasks"})
																																																																																														return
																																																																																															}

																																																																																																totalPages := int(math.Ceil(float64(totalTasks) / float64(limit)))

																																																																																																	responseData := gin.H{
																																																																																																				"tasks": tasks,
																																																																																																						"pagination": gin.H{
																																																																																																										"current_page": page,
																																																																																																													"total_pages":  totalPages,
																																																																																																																"total_tasks":  totalTasks,
																																																																																																																		},
																																																																																																																			}

																																																																																																																				// Simpan ke Redis Cache (TTL 5 Menit)
																																																																																																																					if config.RDB != nil {
																																																																																																																								bytes, _ := json.Marshal(responseData)
																																																																																																																										config.RDB.Set(config.Ctx, cacheKey, bytes, 5*time.Minute)
																																																																																																																											}

																																																																																																																												c.JSON(http.StatusOK, responseData)
																																																																																																																											}

																																																																																																																											// GET /tasks/:id
																																																																																																																											func (tc *TaskController) GetTaskByID(c *gin.Context) {
																																																																																																																													id := c.Param("id")
																																																																																																																														var task models.Task

																																																																																																																															if err := config.DB.First(&task, "id = ?", id).Error; err != nil {
																																																																																																																																		log.Printf("[ERROR] Task ID Not Found (%s): %v", id, err)
																																																																																																																																				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
																																																																																																																																						return
																																																																																																																																							}

																																																																																																																																								c.JSON(http.StatusOK, task)
																																																																																																																																							}

																																																																																																																																							// PUT /tasks/:id
																																																																																																																																							func (tc *TaskController) UpdateTask(c *gin.Context) {
																																																																																																																																									id := c.Param("id")
																																																																																																																																										var task models.Task

																																																																																																																																											if err := config.DB.First(&task, "id = ?", id).Error; err != nil {
																																																																																																																																														log.Printf("[ERROR] Task Update Not Found (%s): %v", id, err)
																																																																																																																																																c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
																																																																																																																																																		return
																																																																																																																																																			}

																																																																																																																																																				var input models.Task
																																																																																																																																																					if err := c.ShouldBindJSON(&input); err != nil {
																																																																																																																																																								log.Printf("[ERROR] Bad Request Update Task: %v", err)
																																																																																																																																																										c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
																																																																																																																																																												return
																																																																																																																																																													}

																																																																																																																																																														task.Title = input.Title
																																																																																																																																																															task.Description = input.Description
																																																																																																																																																																task.Status = input.Status
																																																																																																																																																																	task.DueDate = input.DueDate

																																																																																																																																																																		config.DB.Save(&task)

																																																																																																																																																																			// Invalidate Cache
																																																																																																																																																																				if config.RDB != nil {
																																																																																																																																																																							config.RDB.FlushDB(config.Ctx)
																																																																																																																																																																								}

																																																																																																																																																																									c.JSON(http.StatusOK, gin.H{
																																																																																																																																																																												"message": "Task updated successfully",
																																																																																																																																																																														"task":    task,
																																																																																																																																																																															})
																																																																																																																																																																														}

																																																																																																																																																																														// DELETE /tasks/:id
																																																																																																																																																																														func (tc *TaskController) DeleteTask(c *gin.Context) {
																																																																																																																																																																																id := c.Param("id")
																																																																																																																																																																																	if err := config.DB.Delete(&models.Task{}, "id = ?", id).Error; err != nil {
																																																																																																																																																																																				log.Printf("[ERROR] Delete Task Failed (%s): %v", id, err)
																																																																																																																																																																																						c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus task"})
																																																																																																																																																																																								return
																																																																																																																																																																																									}

																																																																																																																																																																																										// Invalidate Cache
																																																																																																																																																																																											if config.RDB != nil {
																																																																																																																																																																																														config.RDB.FlushDB(config.Ctx)
																																																																																																																																																																																															}

																																																																																																																																																																																																c.JSON(http.StatusOK, gin.H{
																																																																																																																																																																																																			"message": "Task deleted successfully",
																																																																																																																																																																																																				})
																																																																																																																																																																																																			}
