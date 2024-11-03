package seeds

import "earn-expense/app/models"

func hashPassword(password string) string {
	hashedPassword, _ := models.Hash(password)
	return string(hashedPassword)
}

var users = []models.User{
	models.User{
		Name:     "Ruhul",
		Email:    "ruhul@gmail.com",
		Username: "ruhul",
		Password: hashPassword("123456"),
		Status:   1,
	},
	models.User{
		Name: "Emran	",
		Email:    "emrab@gmail.com",
		Username: "emran",
		Password: hashPassword("123456"),
		Status:   1,
	},
}
