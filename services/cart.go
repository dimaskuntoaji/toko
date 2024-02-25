package services

import (
	"errors"
	"toko1/helper"
	"toko1/entity"
	"toko1/repository"
)

type CartService interface {
	CreateCart(input entity.CartInput, userID uint) (entity.CartPostResponse, error)
	GetCarts(userID uint) ([]entity.UserCartResponse, error)
	DeleteCart(productID uint) error
}

type cartService struct {
	cartRepository    repository.CartRepository
	productRepository repository.ProductRepository
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
}

func NewCartService(
	cartRepository repository.CartRepository,
	productRepository repository.ProductRepository,
	userRepository repository.UserRepository,
	categoryRepository repository.CategoryRepository,
) *cartService {
	return &cartService{
		cartRepository,
		productRepository,
		userRepository,
		categoryRepository,
	}
}

//MEMASUKAN PRODUCT KE KERANJANG

// CreateCart godoc
// @Summary      Add to cart
// @Description  Add product to shopping cart
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request body entity.CartInput true "Payload Body [RAW]"
// @Success      200 {object} entity.CartPostResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /cart [post]
// @Security BearerAuth
func (s *cartService) CreateCart(input entity.CartInput, userID uint) (entity.CartPostResponse, error) {
	var (
		cartResponse entity.CartPostResponse
		cart         entity.Cart
	)

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return cartResponse, err
	}
	product, err := s.productRepository.GetDataByID(uint(input.ProductID))
	if err != nil {
		return cartResponse, err
	}
	category, err := s.categoryRepository.GetDataByID(product.CategoryID)
	if err != nil {
		return cartResponse, err
	}

	if *product.Stock < input.Quantity {
		return cartResponse, errors.New("product is not available")
	}

	category, err = s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return cartResponse, err
	}

	stock := *product.Stock - input.Quantity
	*product.Stock = stock
	product, err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return cartResponse, err
	}

	cart.UserID = user.ID
	cart.ProductID = product.ID
	cart.Quantity = input.Quantity

	cart, err = s.cartRepository.CreateCart(cart)
	cartResponse = entity.CartPostResponse{
		Quantity:     cart.Quantity,
		ProductTitle: product.Title,
	}

	return cartResponse, helper.ReturnIfError(err)
}


//MELIHAT PRODUK DI KERANJANG

// GetCarts godoc
// @Summary      Cart
// @Description  List of products that have been added to the shopping cart
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200 {object} entity.UserCartResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /cart [get]
// @Security BearerAuth
func (s *cartService) GetCarts(userID uint) ([]entity.UserCartResponse, error) {
	carts, err := s.cartRepository.GetCarts(userID)
	var cartResponses []entity.UserCartResponse
	for _, cart := range carts {
		cartResponse := entity.UserCartResponse{
			ID:        cart.ID,
			ProductID: cart.ProductID,
			UserID:    cart.UserID,
			Quantity:  cart.Quantity,
			Product: entity.ProductResponse{
				ID:         cart.Product.ID,
				Title:      cart.Product.Title,
				Price:      cart.Product.Price,
				Stock:      *cart.Product.Stock,
				CategoryID: cart.Product.CategoryID,
				CreatedAt:  cart.Product.CreatedAt,
			},
		}
		cartResponses = append(cartResponses, cartResponse)
	}

	return cartResponses, helper.ReturnIfError(err)
}


//#HAPUS PRODUCT DARI KERANJANG

// GetCarts godoc
// @Summary      Delete cart
// @Description  Delete product list in shopping cart
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /cart/1 [delete]
// @Security BearerAuth
func (s *cartService) DeleteCart(cartID uint) error {

	cart, err := s.cartRepository.GetDataByID(cartID)
	if err != nil {
		return err
	}

	product, err := s.productRepository.GetDataByID(uint(cart.ProductID))
	if err != nil {
		return err
	}

	stock := *product.Stock + cart.Quantity
	*product.Stock = stock
	product, err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return err
	}

	err = s.cartRepository.DeleteCart(cart)

	return helper.ReturnIfError(err)
}
