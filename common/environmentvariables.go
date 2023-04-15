package common

import (
	"context"

	"regexp"

	"encoding/json"

	"github.com/go-redis/redis/v8"
)
var (
	JWTKey = "enter jwt key here"
	Rdb *redis.Client = nil
	Ctx = context.TODO()
)

var (
	MongoDbConnectionURL = "mongodb url here"
)

func IsValidAddress(v string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(v)
}

type ContractCacheResponse struct {
	IndexedData []byte `json:"indexed_data"`
	Fees float64 `json:"fees"`
	ImageUrl string `json:"image_url"`
}

func (i *ContractCacheResponse) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}