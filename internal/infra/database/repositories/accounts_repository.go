package repositories

import (
	"context"
	"mongoapi/internal/domain/accounts"
	"mongoapi/internal/main/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountsRepository struct {
	Connection *configs.MongoConfig
}

func (r *AccountsRepository) Create(account *accounts.Account) (string, error) {
	doc := bson.M{
		"name":       account.Name,
		"email":      account.Email,
		"password":   account.Password,
		"created_at": account.CreatedAt,
		"updated_at": account.UpdatedAt,
	}

	res, err := r.Connection.Collection().InsertOne(context.Background(), doc)

	if err != nil {
		return "", err
	}

	// r.Connection.Disconnect()

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (repo *AccountsRepository) FindOneById(id string) (*accounts.Account, error) {
	return nil, nil
}
