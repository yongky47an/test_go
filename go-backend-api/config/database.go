package config

import (
		"fmt"
			"log"
				"os"

					"todo-api/models"

						"gorm.io/driver/postgres"
							"gorm.io/gorm"
						)

						var DB *gorm.DB

						func ConnectDB() {
								dsn := fmt.Sprintf(
											"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
													getEnv("DB_HOST", "192.168.1.109"),
															getEnv("DB_USER", "postgres"),
																	getEnv("DB_PASSWORD", "Admin123!"),
																			getEnv("DB_NAME", "db_go_test"),
																					getEnv("DB_PORT", "5432"),
																						)

																							var err error
																								DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
																									if err != nil {
																												log.Fatalf("Gagal terhubung ke Database: %v", err)
																													}

																														// Auto Migration Schema
																															DB.AutoMigrate(&models.Task{}, &models.User{})
																																log.Println("Database PostgreSQL (db_go_test) Berhasil Terhubung & Migrasi!")
																															}

																															func getEnv(key, fallback string) string {
																																	if value, ok := os.LookupEnv(key); ok {
																																				return value
																																					}
																																						return fallback
																																					}
