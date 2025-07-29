package product

func NewProductController(httpHandler *product.HTTPHandler) *Controller {
	return &Controller{httpHandler: httpHandler}
}

func (c *Controller) HistoryProduct(w http.ResponseWriter, r *http.Request) {
	
}
