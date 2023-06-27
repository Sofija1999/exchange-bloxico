package product

import (
	"context"
	"net/http"

	//"github.com/Bloxico/exchange-gateway/sofija/core/domain"
	"github.com/Bloxico/exchange-gateway/sofija/core/ports"
	"github.com/emicklei/go-restful/v3"
	"errors"
	//"github.com/gorilla/mux"
	//"strconv"
	//"log"
)


type EgwProductHttpHandler struct {
	productSvc ports.EgwProductUsecase
}

func NewEgwProductHandler(productSvc ports.EgwProductUsecase, wsCont *restful.Container) *EgwProductHttpHandler {
	httpHandler := &EgwProductHttpHandler{
		productSvc: productSvc,
	}

	ws := new(restful.WebService)

	ws.Path("/product").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/insert").To(httpHandler.InsertProduct))
	ws.Route(ws.PUT("/update/{id}").To(httpHandler.UpdateProduct))
	//ws.Route(ws.DELETE("/delete/{id}").To(httpHandler.DeleteProduct))

	wsCont.Add(ws)

	return httpHandler
}

// Performs insert product
func (e *EgwProductHttpHandler) InsertProduct(req *restful.Request, resp *restful.Response) {
	var reqData InsertRequestData
	req.ReadEntity(&reqData)

	var egwProduct *EgwProductModel = &EgwProductModel{}
	egwProduct.Name = reqData.Name
	egwProduct.ShortDescription = reqData.ShortDescription
	egwProduct.Description = reqData.Description
	egwProduct.Price = reqData.Price

	err := e.insertProduct(req.Request.Context(), egwProduct)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, errors.New("error insert product"))
		return
	}

	// send product back
	respData := InsertResponseData{Product: *egwProduct}

	resp.WriteAsJson(respData)
}

func (e *EgwProductHttpHandler) insertProduct(ctx context.Context, egwProduct *EgwProductModel) error {
	err := e.productSvc.InsertProduct(ctx, egwProduct.ToDomain())
	if err != nil {
		return err
	}
	// retrieve their data from the DB to populate it (e.g. ID)
	productData, err := e.productSvc.FindByID(ctx, egwProduct.ID)
	if err != nil {
		return err
	}
	egwProduct.FromDomain(productData)
	return nil
}

func (e *EgwProductHttpHandler) UpdateProduct(req *restful.Request, resp *restful.Response) {
	/*
	var a UpdateRequestData
	req.ReadEntity(&a)

	// get product ID for update query
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string int %v", err)
	}
	reqId, err := params.StringFrom(req.Request, auth.USER_ID_CTX_KEY)
	if err != nil || len(reqId) == 0 {
		resp.WriteError(http.StatusBadRequest, errors.New("no id found for product"))
		return
	}

	ctx := req.Request.Context()
	dataProduct := &domain.EgwProduct{ID: id, Name: a.Name, ShortDescription: a.ShortDescription, Description: a.Description, Price: a.Price}

	err = e.productSvc.Update(ctx, dataProduct)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, errors.New("error updating product"))
		return
	}

	// return updated product as data
	var retProduct *EgwProductModel = &EgwProductModel{}
	retProduct.FromDomain(dataProduct)
	resp.WriteAsJson(retProduct)*/
}

// DeleteProduct performs the deletion of a product by ID.
/*func (e *EgwProductHttpHandler) DeleteProduct(req *restful.Request, resp *restful.Response) {
	// Extract the product ID from the request path or query parameters, assuming the ID is provided.
	productID := req.PathParameter("id")

	// Perform the deletion of the product with the given ID.
	err := e.deleteProduct(req.Request.Context(), productID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, errors.New("error deleting product"))
		return
	}

	// Return a success response.
	resp.Write([]byte("Product deleted successfully"))
}*/

/*func (e *EgwProductHttpHandler) deleteProduct(ctx context.Context, productID string) error {
	// Perform the deletion of the product with the given ID using the appropriate service method.
	err := e.productSvc.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}
	return nil
}*/
