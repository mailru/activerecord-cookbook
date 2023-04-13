package dictionary

import "fmt"

type Product struct {
	ID           uint64
	Name         string
	Field        bool
	AnotherField int
}

var products = map[uint64]*Product{
	1:      {ID: 1, Name: "ProdName", Field: true, AnotherField: 22},
	2:      {ID: 2, Name: "ProdNam2", Field: false, AnotherField: 3},
	123456: {ID: 123456, Name: "ProdName123456", Field: true, AnotherField: 123},
}

func GetProductByID(id uint64) (*Product, error) {
	pr, ex := products[id]
	if !ex {
		return nil, fmt.Errorf("Product with id %d not found", id)
	}

	return pr, nil
}
