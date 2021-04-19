package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	ErrorHandler "github.com/timopattikawa/mtp-restful-product/product/error"
	mod "github.com/timopattikawa/mtp-restful-product/product/model"
)

func (api *apiDeliveryProduct) Serve() {

	router := mux.NewRouter()

	router.HandleFunc("/products/", api.ProductGetAllHandle).Methods("GET")
	router.HandleFunc("/products/{id}/", api.ProductGetOneHandle).Methods("GET")
	router.HandleFunc("/products/", api.CreateProductHandle).Methods("POST")
	router.HandleFunc("/products/{id}/", api.DeleteProductHandle).Methods("DELETE")

	log.Println("Run on port 9000")
	http.ListenAndServe(":9000", router)
}

func (adp *apiDeliveryProduct) ProductGetAllHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	dataProduct, err := adp.GetAllDataProduct()

	if err != nil {
		tmpErr := ErrorHandler.NewHTTPError(
			http.StatusNotFound,
			"Sorry, Internal Server Error",
			"Internal Server Error",
			r.URL.Path,
		)
		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(body)
		return
	}

	response := mod.New(200, "OK", dataProduct)

	dataJson, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.Write(dataJson)
}

func (adp *apiDeliveryProduct) ProductGetOneHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	productIDString := params["id"]

	productIDInt, err := strconv.Atoi(productIDString)

	if err != nil {
		log.Println("error id param")
		http.Error(rw, "", http.StatusBadRequest)
		return
	}

	dataProduct, err := adp.GetOneProductByID(int64(productIDInt))

	if err != nil {
		log.Printf("[DELIVERY] Product Not Found with id: %d\n", productIDInt)

		tmpErr := ErrorHandler.NewHTTPError(
			http.StatusNotFound,
			fmt.Sprintf("Not found product %v", productIDInt),
			"Not Found",
			r.URL.Path,
		)

		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(http.StatusNotFound)
		rw.Write(body)
		return
	}

	response := mod.New(200, "OK", dataProduct)
	jsonDataProduct, err := json.Marshal(response)

	if err != nil {
		log.Println("[DELIVERY] Fail Marshal Data")
		http.Error(rw, "", http.StatusNotFound)
		return
	}

	rw.Write(jsonDataProduct)
}

func (adp *apiDeliveryProduct) CreateProductHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	log.Println(decoder)

	productRequest := mod.ProductRequest{}

	if err := decoder.Decode(&productRequest); err != nil {
		log.Printf("[DELIVERY] Fail to decode")
		tmpErr := ErrorHandler.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("Internal server error"),
			"Internal server error",
			r.URL.Path,
		)

		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(tmpErr.Status)
		rw.Write(body)
		return
	}

	log.Println(productRequest)

	status, err := adp.CreateProduct(productRequest)

	if err != nil {
		log.Printf("[DELIVERY] FAIL IN CREATE PRODUCT")

		tmpErr := ErrorHandler.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("Internal server error"),
			"Internal server error",
			r.URL.Path,
		)

		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(body)
		return
	}

	if status == http.StatusBadRequest {
		log.Printf("[DELIVERY] FAIL IN CREATE PRODUCT CAUSE PRODUCT HAS BEEN ADDED")

		tmpErr := ErrorHandler.NewHTTPError(
			status,
			fmt.Sprintf("Cannot create product because product has been added"),
			"Cannot create product",
			r.URL.Path,
		)

		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(status)
		rw.Write(body)
		return
	}

	response := mod.New(status, "OK", productRequest)

	jsonResp, err := json.Marshal(response)

	if err != nil {
		log.Printf("[DELIVERY] FAIL MARSHAL RESPONSE")
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(status)
	rw.Write(jsonResp)
}

func (adp *apiDeliveryProduct) DeleteProductHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)

	productIDString := params["id"]

	productIDInt, err := strconv.Atoi(productIDString)
	if err != nil {
		http.Error(rw, "ID NULL", http.StatusBadRequest)
	}

	status, err := adp.UseCaseProduct.DeleteProductByID(int64(productIDInt))
	if status == http.StatusNotFound && err != nil {
		log.Printf("[DELIVERY] NOT FOUND PRODUCT FOR DELETE")

		tmpErr := ErrorHandler.NewHTTPError(
			status,
			fmt.Sprintf("Not found product for delete"),
			"Not found product",
			r.URL.Path,
		)

		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(status)
		rw.Write(body)
		return
	}

	if status == http.StatusBadRequest && err != nil {
		log.Printf("[DELIVERY] INTERNAL SERVER ERROR FOR DELETE")

		tmpErr := ErrorHandler.NewHTTPError(
			status,
			fmt.Sprintf("internal server error for delete"),
			"INTERNAL SERVER ERROR",
			r.URL.Path,
		)

		body, err := tmpErr.ResponseBody()
		if err != nil {
			log.Println(err)
		}

		rw.WriteHeader(status)
		rw.Write(body)
		return
	}

	response := mod.New(status, "OK", nil)

	jsonResp, err := json.Marshal(response)

	if err != nil {
		log.Printf("[DELIVERY] FAIL MARSHAL RESPONSE")
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(status)
	rw.Write(jsonResp)
}
