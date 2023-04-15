package database

import (
	"log"
	"mm-api/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func GetDataByContract(contract string) (map[string]interface{}, error) {
	filter := bson.M {
		"address": contract,
	}

	resp := make(map[string]interface{})

	
	var data bson.M
	if err := IndexedCollection.FindOne(ctx, filter).Decode(&data); err != nil {
		return resp, err
	}

	

	resp["indexed_data"] = data["indexed_data"].(primitive.Binary).Data
	resp["fees"] = data["fees"].(float64)
	resp["image_url"] = data["image_url"].(string)

	obj := &common.ContractCacheResponse{
		IndexedData: resp["indexed_data"].([]byte),
		Fees: resp["fees"].(float64),
		ImageUrl: resp["image_url"].(string),
	}
	d, err := obj.MarshalBinary()
	if err != nil {
		log.Println("error marshalling data into binary")
		return resp, err
	}
	err = common.Rdb.Set(ctx, contract, d, 30 * time.Minute).Err()
	if err != nil {
		log.Println("Error setting redis cache for contract: ", contract, err)
		return resp, err
	}
	
	return resp, nil
}