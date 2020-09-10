package app

import (
	"html/template"
	"io/ioutil"
	"log"
	"nellymarket/pkg/models"
	"net/http"
	"path/filepath"
	"strconv"
)

func (receiver server) handlerProductsList() func( http.ResponseWriter,  *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		list, err := receiver.marketSvc.ProductsList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		list2, err := receiver.marketSvc.SalesList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := struct {
			Title   string
			Products []models.Prices
			H1 string
			List []models.Sales
		}{
			Title:   "Nelly Market",
			Products: list,
			H1: "Sales list",
			List: list2,
		}

		err = tpl.Execute(writer, data)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func (receiver server) handlerAddProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := request.FormValue("name")
		category := request.FormValue("category")
		price, err := strconv.Atoi(request.FormValue("price"))
		if price == 0 {
			http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		}
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		product := models.Prices{
			Name:     name,
			Category: category,
			Price:    price,
			Removed:  false,
		}
		err = receiver.marketSvc.AddProducts(product)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}
}

func (receiver *server) handleProductsRemove() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := request.FormValue("id")
		newID, err := strconv.Atoi(id)

		err = receiver.marketSvc.RemoveById(newID)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}
}

func (receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {
	file, err := ioutil.ReadFile(filepath.Join(receiver.assetsPath, "favicon.ico"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}

func (receiver server) handlerAddListSales() func(w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		product := request.FormValue("product")
		client := request.FormValue("client")
		count, err := strconv.Atoi(request.FormValue("count"))
		if count == 0 {
			http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		}
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}


		products := models.Sales{
			Product: product,
			Count:   count,
			Client:  client,
		}

		err = receiver.marketSvc.AddSalesInDB(products)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}
}

//func (receiver server) handlerSalesList() func( http.ResponseWriter,  *http.Request) {
//	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
//	if err != nil {
//		panic(err)
//	}
//	return func(writer http.ResponseWriter, request *http.Request) {
//		list, err := receiver.marketSvc.SalesList()
//		if err != nil {
//			log.Print(err)
//			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//			return
//		}
//
//		data := struct {
//			H1   string
//			Sales []models.Sales
//		}{
//			H1:   "Sales list",
//			Sales: list,
//		}
//
//		err = tpl.Execute(writer, data)
//		if err != nil {
//			log.Print(err)
//			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//			return
//		}
//	}
//}