package main

import (
	"encoding/json"
	"os"
)

func SaveModel(ml *model, suffix string) {
	b, _ := json.Marshal(ml)

	err := os.WriteFile("sbcandlesM5_ml005"+suffix+".txt", b, 0644)
	if err != nil {
		panic(err)
	}

}
