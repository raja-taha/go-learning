package services

import "gin-api/models"

var users = []models.User{
	{
		ID:    1,
		Name:  "Taha",
		Email: "taha@example.com",
	},
}

func GetUsers() []models.User {
	return users
}

func CreateUser(user models.User) models.User {
	user.ID = len(users) + 1

	users = append(users, user)

	return user
}
