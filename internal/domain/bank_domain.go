package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type for working with bank services
type BankService struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Details       string             `json:"details" bson:"details"`
	NumberOfUsers int64              `json:"numberOfUsers" bson:"numberOfUsers"`
	Status        string             `json:"status" bson:"status"`
	Type          string             `json:"type" bson:"type"`
	Icon          string             `json:"icon" bson:"icon"`
}

// type for working with the company request
type BankRequest struct {
	Name          string `json:"name" bson:"name"`
	Details       string `json:"details" bson:"details"`
	NumberOfUsers int64  `json:"numberOfUsers" bson:"numberOfUsers"`
	Status        string `json:"status" bson:"status"`
	Type          string `json:"type" bson:"type"`
	Icon          string `json:"icon" bson:"icon"`
}
// interface for working with bank services Repo
type BankRepository interface{
	GetBankById(c context.Context,id string)(*BankService,error)
	UpdateBank(c context.Context,id string,bank *BankRequest)(*BankService,error)
	DeleteBank(c context.Context,id string)error
	GetBanks(c context.Context,page int64,size int)([]*BankService,error)
	PostBank(c context.Context,bank *BankRequest)(*BankService,error)
	SearchByName(c context.Context,term string)(*BankService,error)

}

