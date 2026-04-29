package dtos

type ProductListCreateRequest []ProductCreateRequest

type ProductCreateRequest struct {
	ProductCode string `json:"product_code"` 
}
