package database

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func AddUserUpdate(data UserInfo) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"hardwareID": data.HardwareID,
	}
	_, err := UserCollection.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"timestamp": data.LoginTimestamp,
			"hardwareID": data.HardwareID,
		},
	}, opts)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}