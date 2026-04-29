package products

import (
	"fmt"
	"goBackend/internal/products/dtos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func validateInput(c *gin.Context, request any) bool {
	if err := c.ShouldBindJSON(request); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return false
	}
	return true
}

func sendSuccess(c *gin.Context, statusCode int, message string, data any) {
	response := gin.H{
		"message":    message,
		"success":    true,
		"statusCode": statusCode,
	}
	if data != nil {
		response["result"] = data
	}
	c.JSON(statusCode, response)
}

func sendError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error":      message,
		"success":    false,
		"statusCode": statusCode,
	})
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts(c.Request.Context())
	if err != nil {
		sendError(c, http.StatusInternalServerError, "get all products is failed.")
		return
	}
	sendSuccess(c, http.StatusOK, "Success", products)
}

func (h *Handler) GetProductsByID(c *gin.Context) {
	queryID, ok := c.GetQuery("id")
	if !ok {
		sendError(c, http.StatusBadRequest, "{id} is reqired.")
		return
	}

	id, err := strconv.Atoi(queryID)
	if err != nil {
		sendError(c, http.StatusBadRequest, "id is not a number.")
		return
	}

	product, err := h.service.GetProductByID(c.Request.Context(), int64(id))
	if err != nil {
		errString := fmt.Sprintf("get a product id:{%d}  is failed.", id)
		sendError(c, http.StatusInternalServerError, errString)
		return
	}

	if product == nil {
		errString := fmt.Sprintf("Product id: %d is not found.", id)
		sendError(c, http.StatusNotFound, errString)
		return
	}

	sendSuccess(c, http.StatusOK, "Success", product)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var body dtos.ProductListCreateRequest
	if !validateInput(c, &body) {
		return
	}

	products, err := h.service.CreateProduct(c.Request.Context(), &body)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Created failed.")
		return
	}
	sendSuccess(c, http.StatusCreated, "Created successfuly.", products)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	queryID := c.Param("id")
	if queryID == "" {
		sendError(c, http.StatusBadRequest, "{id} is reqired.")
		return
	}

	id, err := strconv.Atoi(queryID)
	if err != nil {
		sendError(c, http.StatusBadRequest, "id is not a number.")
		return
	}

	rowsAffected, err := h.service.DeleteProduct(c.Request.Context(), int64(id))
	if err != nil {
		errString := fmt.Sprintf("Deleted product id:%d is failed:", id)
		sendError(c, http.StatusBadRequest, errString)
		return
	}

	if rowsAffected == 0 {
		errString := fmt.Sprintf("Rroduct id:%d is not exist", id)
		sendError(c, http.StatusNotFound, errString)
		return
	}

	sendSuccess(c, http.StatusOK, "Deleted successfuly.", nil)

}
