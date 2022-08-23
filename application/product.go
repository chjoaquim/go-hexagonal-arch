package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type (
	ProductInterface interface {
		IsValid() (bool, error)
		Enable() error
		Disable() error
		GetID() string
		GetName() string
		GetStatus() string
		GetPrice() float64
	}

	ProductServiceInterface interface {
		Get(id string) (ProductInterface, error)
		Create(name string, price float64) (ProductInterface, error)
		Enable(productInterface ProductInterface) (ProductInterface, error)
		Disable(productInterface ProductInterface) (ProductInterface, error)
	}

	ProductReader interface {
		Get(id string) (ProductInterface, error)
	}

	ProductWriter interface {
		Save(product ProductInterface) (ProductInterface, error)
	}

	ProductPersistenceInterface interface {
		ProductReader
		ProductWriter
	}

	Product struct {
		ID     string  `valid:"uuidv4"`
		Name   string  `valid:"required"`
		Status string  `valid:"required"`
		Price  float64 `valid:"float,optional"`
	}
)

const (
	ENABLED  = "enabled"
	DISABLED = "disabled"
)

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater than zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("price must be greater than zero to enable product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("price must be zero to enable product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
