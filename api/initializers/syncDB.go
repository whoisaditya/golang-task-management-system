package initializers

import "github.com/whoisaditya/golang-task-management-system/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.Group{})
	DB.AutoMigrate(&models.Task{})
}