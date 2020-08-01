package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "golang.org/x/text/message"
	"ims-be/dto"
	"ims-be/entity"
	_ "ims-be/entity"
	"ims-be/response"
	"log"
	"net/http"
	"strconv"
)

type IBranchController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	DeActiveProduct(w http.ResponseWriter, r *http.Request)
	ActiveProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

type branchDto struct {
	dto.IBranchDto
}
type brR response.BranchResponse
type brDR response.BranchDetailResponse
type brUR response.UpdateBranch
type lbrR response.ListBranchResponse
type rjE entity.ReponseJson

// @tags product-apis
//@Summary Create a new product
//@Description Create a new product
//@Accept json
//@Param product body dto.ProductPayload true "product information"
//@Success 200 {object} controller.ProductResponse
//@Router /product [post]

func (brDto branchDto) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var data dto.BranchData
	response := &brR{Message: "Data invalidate, please enter name of product"}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		response = &brR{
			Message: err.Error(),
			Success: false,
		}
	} else {
		if data.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := brDto.CreateBranchDto(&data)
			if err != nil {
				response = &brR{
					Message: err.Error(),
					Success: false,
				}
			} else {
				response = &brR{
					Message: "success",
					Success: true,
				}
			}
		}
	}
	render.JSON(w, r, response)
}

//@tags product-apis
//@Summary Get a product
//@Description Get a product
//@Accept json
//@Param Id path integer true "product id"
//@Success 200 {object} controller.ProductDetailResponse
//@Router /product/{Id} [get]

func (brDto branchDto) GetProductById(w http.ResponseWriter, r *http.Request) {
	brIdString := chi.URLParam(r, "Id")
	branchId, _ := strconv.Atoi(brIdString)
	response := &brDR{}
	log.Print(branchId)

	if branchId < 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = &brDR{
			Data:    nil,
			Message: "Invalidate, please enter unsigned integer in branch",
			Success: false,
		}
		render.JSON(w, r, response)
		return
	}

	product, err := brDto.GetBranchByIdDto(branchId)
	if err != nil {
		response = &brDR{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		response = &brDR{
			Data:    product,
			Message: "Success",
			Success: true,
		}
	}
	render.JSON(w, r, response)
}

//@tags product-apis
//@Summary Get list of products
//@Description Get list of products
//@Accept json
//@Param page query int false "page"
//@Param pageSize query int false "pageSize"
//@Success 200 {object} controller.ListProductResponse
//@Router /products [get]

func (brDto branchDto) GetProducts(w http.ResponseWriter, r *http.Request) {
	var page int
	var pageSize int
	var err error

	page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	pageSize, err = strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 5
	}
	data, meta, err := brDto.GetBranchesDtos(page, pageSize)
	response := &lbrR{}
	if err != nil {
		response = &lbrR{
			Data:    nil,
			Meta:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		response = &lbrR{
			Data:    data,
			Meta:    meta,
			Message: "Success",
			Success: true,
		}
	}
	render.JSON(w, r, response)
}

//@tags product-apis
//@Summary DeActive a product
//@Description DeActive a product
//@Accept json
//@Param Id path integer true "product id"
//@Success 200 {object} controller.ProductDetailResponse
//@Router /product/deactive/{Id} [put]
//func (brDto branchDto) DeActiveProduct(w http.ResponseWriter, r *http.Request) {
//	productIdString := chi.URLParam(r, "Id")
//	productId, _ := strconv.Atoi(productIdString)
//	response := &brUR{}
//	if productId < 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		response = &brUR{
//			Data:    nil,
//			Message: "Invalidate, please enter unsigned integer",
//			Success: false,
//		}
//		render.JSON(w, r, response)
//		return
//	}
//
//	err := p.D(productId)
//	if err != nil {
//		response = &brUR{
//			Message: err.Error(),
//			Success: false,
//		}
//	} else {
//		response = &brUR{
//			Message: "Success",
//			Success: true,
//		}
//	}
//	render.JSON(w, r, response)
//}

//@tags product-apis
//@Summary Active a product
//@Description Active a product
//@Accept json
//@Param Id path integer true "product id"
//@Success 200 {object} controller.ProductDetailRespon	se
//@Router /product/active/{Id} [put]
//func (brDto branchDto) ActiveProduct(w http.ResponseWriter, r *http.Request) {
//	productIdString := chi.URLParam(r, "Id")
//	productId, _ := strconv.Atoi(productIdString)
//	response := &ProductDetailResponse{}
//	if productId < 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		response = &ProductDetailResponse{
//			Data:    nil,
//			Message: "Invalidate, please enter unsigned integer",
//			Success: false,
//		}
//		render.JSON(w, r, response)
//		return
//	}
//
//	err := p.productService.ActiveProduct(productId)
//	if err != nil {
//		response = &ProductDetailResponse{
//			Message: err.Error(),
//			Success: false,
//		}
//	} else {
//		response = &ProductDetailResponse{
//			Message: "Success",
//			Success: true,
//		}
//	}
//	render.JSON(w, r, response)
//}

//@tags product-apis
//@Summary Update a product
//@Description Update a product
//@Accept json
//@Param Id path integer true "product id"
//@Param product body controller.UpdateProduct true "product information"
//@Success 200 {object} controller.ProductDetailResponse
//@Router /product/update/{Id} [put]

func (brDto branchDto) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	productIdString := chi.URLParam(r, "Id")
	productId, _ := strconv.Atoi(productIdString)

	var UpdateBr = &brUR{}
	err := json.NewDecoder(r.Body).Decode(&UpdateBr)

	response := &brR{}

	if productId < 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = &brR{
			Message: "Invalidate, please enter unsigned integer",
			Success: false,
		}
		render.JSON(w, r, response)
		return
	}

	err = brDto.UpdateBranchDto(productId, UpdateBr.Name)
	if err != nil {
		response = &brR{
			Message: err.Error(),
			Success: false,
		}
	} else {
		response = &brR{
			Message: "Success",
			Success: true,
		}
	}
	render.JSON(w, r, response)
}

func NewProductController() IBranchController {
	productService, _ := dto.New()
	return &productController{
		productService: productService,
	}
}
