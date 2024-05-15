package listing 

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error 
	GetById(c *fiber.Ctx) error 
	Update(c *fiber.Ctx) error 
	Delete(c *fiber.Ctx) error
}

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) GetAll(c *fiber.Ctx) error  {
	product := controller.ProductService.GetAll(r.Context())
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) Create(c *fiber.Ctx) error {
	productCreateRequest := web.ProductsRequest{}
	helper.ReadFromRequestBody(r, &productCreateRequest)

	product := controller.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) GetById(c *fiber.Ctx) error {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	product := controller.ProductService.GetById(r.Context(), id)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) Update(c *fiber.Ctx) error {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(r, &productUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	product := controller.ProductService.Update(r.Context(), productUpdateRequest)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
		Data: product,
		Message: "SUCCES UPDATE",
	}

	helper.WriteToRequestBody(w, webResponse)
}

func (controller *ProductControllerImpl) Delete(c *fiber.Ctx) error {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(r.Context(), id)
	webResponse := web.WebResponse {
		Code: 201,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, webResponse)
}
