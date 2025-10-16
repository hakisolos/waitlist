package config

import (
	"fmt"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnDB() {
	var uri string = os.Getenv("MONGO_URI")

	err := mgm.SetDefaultConfig(nil, "waitlist", options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected successfully")
}
