package product

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
	ws.Route(ws.PUT("").To(httpHandler.UpdateProduct))

	wsCont.Add(ws)

	return httpHandler
}

// Performs login or register
func (e *EgwProductHttpHandler) InsertProduct(req *restful.Request, resp *restful.Response) {
	var reqData InsertRequestData
	req.ReadEntity(&reqData)

	var egwProduct *EgwProductModel = &EgwProductModel{}
	egwProduct.Name = reqData.Name
	egwProduct.ShortDescription = reqData.ShortDescription
	egwProduct.Description = reqData.Description
	egwProduct.Price = reqData.Price
	egwProduct.CreatedAt = time.Now()

	err = e.insertProduct(req.Request.Context(), egwProduct)
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