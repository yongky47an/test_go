package config

import (
		"context"
			"fmt"
				"log"
					"os"

						"todo-api/models"

							"github.com/redis/go-redis/v9"
								"gorm.io/driver/postgres"
									"gorm.io/gorm"
								)

								var (
										DB  *gorm.DB
											RDB *redis.Client
												Ctx = context.Background()
											)

											func ConnectDB() {
													dsn := fmt.Sprintf(
																"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
																		getEnv("DB_HOST", "localhost"),
																				getEnv("DB_USER", "postgres"),
																						getEnv("DB_PASSWORD", "postgres"),
																								getEnv("DB_NAME", "todo_db"),
																										getEnv("DB_PORT", "5432"),
																											)

																												var err error
																													DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
																														if err != nil {
																																	log.Fatalf("Gagal terhubung ke Database: %v", err)
																																		}

																																			DB.AutoMigrate(&models.Task{}, &models.User{})
																																				log.Println("Database Postgres Berhasil Terhubung & Migrasi!")
																																			}

																																			func ConnectRedis() {
																																					RDB = redis.NewClient(&redis.Options{
																																								Addr: fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "6379")),
																																									})

																																										_, err := RDB.Ping(Ctx).Result()
																																											if err != nil {
																																														log.Printf("Peringatan: Redis tidak terhubung (%v). Menjalankan tanpa cache.\n", err)
																																																RDB = nil
																																																	} else {
																																																				log.Println("Redis Berhasil Terhubung!")
																																																					}
																																																				}

																																																				func getEnv(key, fallback string) string {
																																																						if value, ok := os.LookupEnv(key); ok {
																																																									return value
																																																										}
																																																											return fallback
																																																										}
