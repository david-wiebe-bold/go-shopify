package goshopify

import (
	"fmt"
	"time"
)

const cartsBasePath = "carts"
const cartsResourceName = "carts"

// CartService is an interface for interfacing with the carts endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/cart
type CartService interface {
	Get(int64, interface{}) (*Cart, error)
	Create(Cart) (*Cart, error)
	Update(Cart) (*Cart, error)
	Delete(int64) error
}

// CartServiceOp handles communication with the product related methods of
// the Shopify API.
type CartServiceOp struct {
	client *Client
}

// Cart represents a Shopify cart
type Cart struct {
	ID               int        `json:"id,omitempty"`
	SessionToken     string     `json:"session_token,omitempty"`
	Identifier       string     `json:"identifier,omitempty"`
	CurrencyTemplate string     `json:"currency_template,omitempty"`
	Notes            string     `json:"notes,omitempty"`
	TotalWeight      int        `json:"total_weight,omitempty"`
	ItemCount        int        `json:"item_count,omitempty"`
	RequiresShipping bool       `json:"requires_shipping,omitempty"`
	CartCompleted    bool       `json:"cart_completed,omitempty"`
	IncludesTax      bool       `json:"includes_tax,omitempty"`
	ShippingIsTaxed  bool       `json:"shipping_is_taxed,omitempty"`
	NativeCartID     string     `json:"native_cart_id,omitempty"`
	NativeCartType   string     `json:"native_cart_type,omitempty"`
	ShopID           int        `json:"shop_id,omitempty"`
	PublicToken      string     `json:"public_token,omitempty"`
	Attributes       string     `json:"attributes,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// CartResource represents the result from the carts/X.json endpoint
type CartResource struct {
	Cart *Cart `json:"cart"`
}

// CartsResource represents the result from the carts.json endpoint
type CartsResource struct {
	Carts []Cart `json:"carts"`
}

// CartTagsResource ~epresents the result from the carts/tags.json endpoint
type CartTagsResource struct {
	Tags []string `json:"tags"`
}

// Get cart
func (s *CartServiceOp) Get(cartID int64, options interface{}) (*Cart, error) {
	path := fmt.Sprintf("%s/%v.json", cartsBasePath, cartID)
	resource := new(CartResource)
	err := s.client.Get(path, resource, options)
	return resource.Cart, err
}

// Create a new cart
func (s *CartServiceOp) Create(cart Cart) (*Cart, error) {
	path := fmt.Sprintf("%s.json", cartsBasePath)
	wrappedData := CartResource{Cart: &cart}
	resource := new(CartResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Cart, err
}

// Update an existing cart
func (s *CartServiceOp) Update(cart Cart) (*Cart, error) {
	path := fmt.Sprintf("%s/%d.json", cartsBasePath, cart.ID)
	wrappedData := CartResource{Cart: &cart}
	resource := new(CartResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Cart, err
}

// Delete an existing cart
func (s *CartServiceOp) Delete(cartID int64) error {
	path := fmt.Sprintf("%s/%d.json", cartsBasePath, cartID)
	return s.client.Delete(path)
}
