package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JSON array that contains multiple objects
const jsonArray string = `[
	{
		"id": 1,
		"name": "oranges",
		"kind": "fruit",
		"amount": "3000kg",
		"origin": {
			"city": "Tacoma",
			"state": "Washington",
			"country": "USA",
			"suppliers": [ "Best Produce Co.", "FreshCo", "Walmart" ]
		}
	},
	{
		"id": 2,
		"name": "strawberries",
		"kind": "fruit",
		"amount": "1000kg",
		"origin": {
			"city": "Watsonville",
			"state": "California",
			"country": "USA",
			"suppliers": [ "Berry Berry Co.", "Greenery Co.", "Walgreens" ]
		}
	},
	{
		"id": 3,
		"name": "broccoli",
		"kind": "vegetable",
		"amount": "300kg",
		"origin": {
			"city": "Guadalajara",
			"state": "Jalisco",
			"country": "Mexico",
			"suppliers": [ "Delish Co." ]
		}
	}
]`

func handleJSONObject(object interface{}, key, indentation string) {
	switch t := object.(type) {
	case string:
		fmt.Println(indentation+key+": ", t) // t has type string
	case bool:
		fmt.Println(indentation+key+": ", t) // t has type bool
	case float64:
		fmt.Println(indentation+key+": ", t) // t has type float64
	case map[string]interface{}:
		fmt.Println(indentation + key + ":")
		for k, v := range t {
			handleJSONObject(v, k, indentation+"\t")
		}
	case []interface{}:
		fmt.Println(indentation + key + ":")
		for index, v := range t {
			handleJSONObject(v, "["+strconv.Itoa(index)+"]", indentation+"\t")
		}
	}
}

func main() {
	// Empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal JSON to the interface.
	json.Unmarshal([]byte(jsonArray), &results)

	// For each array object:
	for _, result := range results {
		// But if you don't know the field types, you can use type switching to determine (safe):
		// Keep in mind that, since this is a map, the order is not guaranteed.
		fmt.Println("\nType Switching: ")
		for k := range result {
			handleJSONObject(result[k], k, "\t")
		}

		fmt.Println("------------------------------")
	}
}
