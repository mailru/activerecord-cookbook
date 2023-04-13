package serializer

import (
	"github.com/mailru/activerecord-cookbook/example/model/dictionary"
)

func ProductUnmarshal(val uint64, v *dictionary.Product) error {
	product, err := dictionary.GetProductByID(val)
	if err != nil {
		return err
	}

	*v = *product

	return nil
}

func ProductMarshal(data *dictionary.Product) (uint64, error) {
	return data.ID, nil
}
