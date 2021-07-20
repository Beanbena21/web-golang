package env

import (
	"os"
)

func MongoEnv(key string) string {
	os.Setenv(key, "mongodb+srv://bean:bean213@cluster0.4uxui.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	return os.Getenv(key)
}

var DB = os.Getenv("dqw")
