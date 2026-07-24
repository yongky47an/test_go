package repository

import (
		"sync"
			"todo-api/config"
				"todo-api/models"
			)

			type TaskRepository struct{}

			func (r *TaskRepository) GetAllConcurrent(status, search string, limit, offset int) ([]models.Task, int64, error) {
					var tasks []models.Task
						var total int64
							var errTasks, errCount error

								var wg sync.WaitGroup
									wg.Add(2)

										// Goroutine 1: Fetch List Data
											go func() {
														defer wg.Done()
																query := config.DB.Model(&models.Task{})
																		if status != "" {
																						query = query.Where("status = ?", status)
																								}
																										if search != "" {
																														query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
																																}
																																		errTasks = query.Limit(limit).Offset(offset).Find(&tasks).Error
																																			}()

																																				// Goroutine 2: Count Total Data
																																					go func() {
																																								defer wg.Done()
																																										query := config.DB.Model(&models.Task{})
																																												if status != "" {
																																																query = query.Where("status = ?", status)
																																																		}
																																																				if search != "" {
																																																								query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
																																																										}
																																																												errCount = query.Count(&total).Error
																																																													}()

																																																														wg.Wait()

																																																															if errTasks != nil {
																																																																		return nil, 0, errTasks
																																																																			}
																																																																				if errCount != nil {
																																																																							return nil, 0, errCount
																																																																								}

																																																																									return tasks, total, nil
																																																																								}
