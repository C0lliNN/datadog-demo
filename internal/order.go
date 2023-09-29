package internal

import "encoding/json"

type Order struct {
	ID string `json:"id" bson:"_id"`
	Items []Item `json:"items" bson:"items"`
	CreatedAt int64
}

func (o Order) TotalPrice() int {
	total := 0
	for _, item := range o.Items {
		total += item.Price
	}
	return total
}

func (o *Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	return json.Marshal(&struct {
		TotalPrice int `json:"total_price"`
		*Alias
	}{
		TotalPrice: o.TotalPrice(),
		Alias: (*Alias)(o),
	})
}