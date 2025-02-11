package initializers

import "github.com/roadleft/roadleft-msgc/models"

func SyncDatabse() {
	DB.AutoMigrate(&models.User{})
}
