package internal

type Item struct {
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Price int `json:"price" bson:"price"`
}