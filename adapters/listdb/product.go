package listdb

import (
	"github.com/chjoaquim/go-hexagonal-arch/application"
)

type ListDb struct {
	list []application.Product
}

func NewListDb() *ListDb {
	return &ListDb{
		list: make([]application.Product, 0),
	}
}

func (l *ListDb) Get(id string) (application.ProductInterface, error) {
	for _, v := range l.list {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return &application.Product{}, nil
}

func (l *ListDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	l.list = append(l.list, application.Product{
		ID:     product.GetID(),
		Name:   product.GetName(),
		Status: product.GetStatus(),
		Price:  product.GetPrice(),
	})

	return product, nil
}
