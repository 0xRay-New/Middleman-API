package database

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mm-api/common"

	"github.com/andybalholm/brotli"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetIndexedContractHandler(c *fiber.Ctx) error {
	contract := c.Params("contract")
	if !common.IsValidAddress(contract) {
		return c.Status(401).JSON(fiber.Map{"success": false, "message": "invalid contract address"})
	}

	var decompressor io.Reader

	var d common.ContractCacheResponse
	
	if val, err := common.Rdb.Get(common.Ctx, contract).Result(); err == redis.Nil && err != nil {
		data, err := GetDataByContract(contract)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"success": false, "message": "Data not found in database", "error": err.Error()})
		}
		decompressor = brotli.NewReader(bytes.NewReader(data["indexed_data"].([]byte)))
		decompressed, _ := ioutil.ReadAll(decompressor)
		
		
		
		return c.Status(200).JSON(fiber.Map{"success": true, "data":fiber.Map{
			"indexed_data": string(decompressed),
			"fees": data["fees"].(float64),
			"image_url": data["image_url"].(string),
		}})
	} else if err != nil && err != redis.Nil{
		log.Println("Error getting data from redis,",err)
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Error getting data from redis", "error": err.Error()})
	} else {
		log.Println("found data for",contract,"in cache")
		json.Unmarshal([]byte(val), &d)
		data := d.IndexedData
		decompressor = brotli.NewReader(bytes.NewReader(data))
		decompressed, _ := ioutil.ReadAll(decompressor)
		return c.Status(200).JSON(fiber.Map{"success": true, "data": fiber.Map{
			"indexed_data": string(decompressed),
			"fees": d.Fees,
			"image_url": d.ImageUrl,
		}})
	}


	
}