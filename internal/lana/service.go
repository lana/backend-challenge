package lana

import (
	"context"
	"errors"
	"patriciabonaldy/lana/internal/models"
	"patriciabonaldy/lana/internal/platform/storage"

	"github.com/google/uuid"
)

var (
	ErrIDIsRequired = errors.New("basket id is required")
	promotion       = map[string]func(item models.Item) models.Item{
		models.Tshirt: models.Discount3OrMore,
		models.Pen:    models.DiscountBuyingTwoGetOneFree,
	}
)

// Service is the default lanaService interface
// implementation returned by lana.NewService.
type Service struct {
	repository storage.Repository
}

type ProductRequest struct {
	BasketID    string `json:"basket_id" binding:"required"`
	ProductCode string `json:"product_code" binding:"required"`
}

type BasketRequest struct {
	BasketID string `json:"basket_id" binding:"required"`
}

// NewService returns the default Service interface implementation.
func NewService(repository storage.Repository) Service {
	return Service{repository: repository}
}

// CreateBasket create a basket.
// it will return a new basket if this is ok.
// otherwise will return error
func (s Service) CreateBasket(ctx context.Context) (models.Basket, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return models.Basket{}, err
	}

	basket, err := s.repository.CreateBasket(ctx, id.String())
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

// GetBasket return a basket.
// require a basket id
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) GetBasket(ctx context.Context, id string) (models.Basket, error) {
	basket, err := s.repository.FindBasketByID(ctx, id)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

// RemoveBasket remove a basket.
// it will remove basket if this is ok.
// otherwise will return error
func (s Service) RemoveBasket(ctx context.Context, id string) error {
	err := s.repository.RemoveBasket(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) createItem(basket models.Basket, productCode string) (models.Item, error) {
	if basket.Close {
		return models.Item{}, models.ErrBasketIsClosed
	}

	product, ok := models.ProductMap[productCode]
	if !ok {
		return models.Item{}, models.ErrProductNotFound
	}

	item, ok := basket.Items[product.Code]
	if !ok {
		_item := models.Item{
			Product: product,
		}
		_item.WithOutDiscount()

		return _item, nil
	}

	return item, nil
}

// AddProduct add a new product into basket.
// require a basket id and product code
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) AddProduct(ctx context.Context, request ProductRequest) (models.Basket, error) {
	var basket, err = s.repository.FindBasketByID(ctx, request.BasketID)
	if err != nil {
		return models.Basket{}, err
	}

	if basket.Close {
		return models.Basket{}, models.ErrBasketIsClosed
	}

	item, err := s.repository.GetItem(ctx, request.BasketID, request.ProductCode)
	if err != nil {
		item, err = s.createItem(basket, request.ProductCode)
		if err != nil {
			return models.Basket{}, err
		}
	}

	item.Quantity++
	item.WithOutDiscount()
	code := item.Product.Code
	basket.Items[code] = item
	basket.CalculateTotal()

	basket, err = s.repository.UpdateBasket(ctx, basket)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

// RemoveProduct remove product inside basket.
// require a basket id and product code
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) RemoveProduct(ctx context.Context, request ProductRequest) (models.Basket, error) {
	_, ok := models.ProductMap[request.ProductCode]
	if !ok {
		return models.Basket{}, models.ErrProductNotFound
	}

	var basket, err = s.repository.FindBasketByID(ctx, request.BasketID)
	if err != nil {
		return models.Basket{}, err
	}

	if basket.Close {
		return models.Basket{}, models.ErrBasketIsClosed
	}

	basket, err = s.repository.RemoveProduct(ctx, request.BasketID, request.ProductCode)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

// CheckoutBasket close a basket.
// require a basket id
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) CheckoutBasket(ctx context.Context, request BasketRequest) (models.Basket, error) {
	var basket models.Basket
	basket, err := s.repository.FindBasketByID(ctx, request.BasketID)
	if err != nil {
		return models.Basket{}, err
	}

	for _, item := range basket.Items {
		calculatePromotion, ok := promotion[item.Product.Code]
		if ok {
			item = calculatePromotion(item)
		}

		basket.Items[item.Product.Code] = item
	}

	basket.CalculateTotal()
	basket.Close = true
	basket, err = s.repository.UpdateBasket(ctx, basket)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}
