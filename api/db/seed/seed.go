package seed

import (
	"encoding/json"
	"lightban/api/db"
	"lightban/api/model"
	"os"
)

func Run(db *db.DB) {
	// read the options from the file

	data, err := os.ReadFile("options.json")

	if err != nil {
		panic(err)
	}

	// parse the options
	var ops []model.Option

	if err := json.Unmarshal(data, &ops); err != nil {
		panic(err)
	}

	// insert the options into the database
	for _, op := range ops {
		if err := db.CreateOption(&op); err != nil {
			panic(err)
		}

	}
}
