package db

import "github.com/darth-licht/go_learn/shopping/models"

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
