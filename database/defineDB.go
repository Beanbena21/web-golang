package database

type Users struct {
	UserName     string `bson: UserName, not is empty`
	UserPassword string `bson: UserPassword, not is empty`
	Email        string `bson: UserEmail, not is empty`
}

//get collection
var Collection = client.Database("fet_kltn").Collection("users")
