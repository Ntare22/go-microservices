package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID					 int	`json:"id"` // these are struct tags, used to add annotations to a field eg. changed the name of output field
	Name				 string `json:"name"`
	Description	 string `json:"description"`
	Price				 float32 `json:"price"`
	SKU					 string `json:"sku"`
	CreatedOn		 string `json:"-"`
	UpdatedOn		 string	`json:"-"`
	DeletedOn		 string	`json:"-"`
}

func GetProducts() Products {
	return productList
}

type Products []*Product

func (p*Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = Products {
	&Product{
		ID: 					1,
		Name: 				"Latte",
		Description: 	"Frothy milky coffee",
		Price:				2.45,
		SKU:					"abc323",
		CreatedOn: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
	&Product{
		ID: 					2,
		Name: 				"Espress",
		Description: 	"Short and strong coffee without milk",
		Price:				1.99,
		SKU:					"fjd323",
		CreatedOn: 		time.Now().UTC().String(),
		UpdatedOn: 		time.Now().UTC().String(),
	},
}