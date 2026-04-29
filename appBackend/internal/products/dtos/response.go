package dtos

type ProductResponse struct {
    ID          int64  `json:"id"`
    ProductCode string `json:"product_code"`

}

type ProductListResponse struct {
    Data       []*ProductResponse `json:"data"`
}