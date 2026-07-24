package models

import (
		"time"

			"github.com/google/uuid"
				"gorm.io/gorm"
			)

			type Task struct {
					ID          string         `gorm:"primaryKey;type:uuid" json:"id"`
						Title       string         `gorm:"not null" json:"title" binding:"required"`
							Description string         `json:"description"`
								Status      string         `gorm:"type:varchar(20);default:'pending'" json:"status" binding:"required,oneof=pending completed"`
									DueDate     string         `gorm:"type:date" json:"due_date" binding:"required"`
										CreatedAt   time.Time      `json:"created_at"`
											UpdatedAt   time.Time      `json:"updated_at"`
												DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
											}

											func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
													if t.ID == "" {
																t.ID = uuid.New().String()
																	}
																		return
																	}

																	type User struct {
																			ID       string `gorm:"primaryKey;type:uuid"`
																				Username string `gorm:"unique;not null"`
																					Password string `gorm:"not null"`
																				}
