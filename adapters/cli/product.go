package cli

import (
	"fmt"
	"github.com/chjoaquim/go-hexagonal-arch/application"
)

func Run(service application.ProductServiceInterface, action, productName, productId, productStatus string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("PRODUCT_CREATED: ID %s\n Name %s\n Price %f\n",
			product.GetID(), product.GetName(), product.GetPrice())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("PRODUCT_ENABLED: ID %s\n Name %s\n Status %s\n",
			res.GetID(), res.GetName(), res.GetStatus())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("PRODUCT_DISABLED: ID %s\n Name %s\n Status %s\n",
			res.GetID(), res.GetName(), res.GetStatus())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("PRODUCT_FOUND: ID %s\n Name %s\n Status %s\n, Price %f\n",
			product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())
	}

	return result, nil
}
