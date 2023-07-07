package controller

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/data/response"
	"golang-jwttoken/helper"
	"golang-jwttoken/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (controller *ProductController) CreateNewProduct(ctx *gin.Context) {
	createNewProductRequest := request.CreateProductRequest{}
	err := ctx.ShouldBindJSON(&createNewProductRequest)
	helper.ErrorPanic(err)

	controller.ProductService.CreateNewProduct(createNewProductRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created product!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProductController) GetAllProduct(ctx *gin.Context) {
	products, err := controller.ProductService.GetAllProduct()
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Ok",
			Message: "Bad request",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
	} else {
		webResponse := response.Response{
			Code:    200,
			Status:  "Ok",
			Message: "Succesfully fetch all products data!",
			Data:    products,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ProductController) GetProductById(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("productId"))
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Ok",
			Message: "Bad request, please correct the productId",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	product, err := controller.ProductService.GetProductById(productId)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Ok",
			Message: "Bad request",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Product by Id",
		Data:    product,
	}
	ctx.JSON(http.StatusBadRequest, webResponse)
}

func (controller *ProductController) GetProductByProductName(ctx *gin.Context) {
	productName := ctx.Param("productName")
	products, err := controller.ProductService.GetProductByProductName(productName)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Ok",
			Message: "Bad request",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all products by productName",
		Data:    products,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
