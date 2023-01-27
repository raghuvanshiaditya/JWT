package initializers

import (
	"example.com/m/models"
)

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
