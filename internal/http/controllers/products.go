package controllers

import (
	"net/http"
	"strconv"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/utils"
)

func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	createProductPayload := r.Context().Value(constants.CreateProductPayloadCtxKey{}).(dto.CreateProductPayload)

	product, err := c.productservice.CreateProduct(createProductPayload)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	utils.Dispatch200(w, "Product created successfully", product)

}

func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updateProductPayload := r.Context().Value(constants.UpdateProductPayloadCtxKey{}).(dto.UpdateProductPayload)
	productIdStr, err := utils.GetPathParam(r, "id")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
	}
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload, product Id must be an integer", err)
	}

	product, err := c.productservice.UpdateProduct(uint(productId), updateProductPayload)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	utils.Dispatch200(w, "Product updated successfully", product)
}

func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productIdStr, err := utils.GetPathParam(r, "id")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
	}
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload, product Id must be an integer", err)
	}
	err = c.productservice.DeleteProduct(uint(productId))
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	utils.Dispatch204(w)
}

func (c *Controller) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.productservice.GetAllProducts()
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	utils.Dispatch200(w, "Products retrieved successfully", products)
}

func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
	productIdStr, err := utils.GetPathParam(r, "id")
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload", err)
	}
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		utils.Dispatch400Error(w, "Invalid Payload, product Id must be an integer", err)
	}

	product, err := c.productservice.GetProductByID(uint(productId))
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	utils.Dispatch200(w, "Product retrieved successfully", product)
}
