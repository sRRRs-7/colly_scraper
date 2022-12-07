package db

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadJson(path string) interface{} {
	byteArr, _ := os.ReadFile(path)
	var jsonObj interface{}
	json.Unmarshal(byteArr, &jsonObj)
	return jsonObj
}

func JsonMapping(jsonObj interface{}) interface{} {
	products := []Product{}

	cnt := jsonObj.(map[string]interface{})["product_count"].(interface{})[0].(int)
	for i := 0; i < int(cnt); i++ {
		id := jsonObj.(map[string]interface{})["product"].([]interface{})[i].(map[string]interface{})["id"].(float64)
		title := jsonObj.(map[string]interface{})["product"].([]interface{})[i].(map[string]interface{})["title"].(string)
		url := jsonObj.(map[string]interface{})["product"].([]interface{})[i].(map[string]interface{})["url"].(string)
		price := jsonObj.(map[string]interface{})["product"].([]interface{})[i].(map[string]interface{})["price"].(string)
		star := jsonObj.(map[string]interface{})["product"].([]interface{})[i].(map[string]interface{})["star"].(string)
		image := jsonObj.(map[string]interface{})["product"].([]interface{})[i].(map[string]interface{})["image"].(string)

		product := Product{
			ID:          int(id),
			ProductName: title,
			URL:         url,
			Price:       price,
			Star:        star,
			Image:       image,
		}
		products = append(products, product)
	}

	return products
}

func JsonInitData() {
	jsonObj := LoadJson("/Users/srrrs/App/APP/service/scraper/db/scrape.json")
	products := JsonMapping(jsonObj)
	fmt.Println(products)
}
