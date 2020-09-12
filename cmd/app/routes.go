package app

func (receiver *server) InitRoutes() {
	mux := receiver.router.(*exactMux)
	mux.GET("/", receiver.handlerProductsList())
	mux.POST("/", receiver.handlerProductsList())

	mux.POST("/market/addproduct", receiver.handlerAddProduct())
	mux.POST("/market/removeproduct", receiver.handleProductsRemove())

	mux.POST("/market/addsaleslist", receiver.handlerAddListSales())
	mux.GET("/favicon.ico", receiver.handleFavicon())
}
