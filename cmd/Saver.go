package main

import (
	"encoding/json"
	"os"
)

func SaveModel(ml *model) {
	b, _ := json.Marshal(ml)

	err := os.WriteFile("sbcandlesM5_ml002.txt", b, 0644)
	if err != nil {
		panic(err)
	}

}
