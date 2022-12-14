package handlers

import (
	"context"
	"net/http"

	protos "github.com/0fatih/currency/protos/currency"

	"github.com/0fatih/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll returns the products from the data stored
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	rw.Header().Add("Content-Type", "application/json")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Println("[ERROR] serializing produtc", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching produc", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// get exchange ret
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["TRY"]),
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[ERROR] error getting new rate", err)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
	}

	prod.Price = prod.Price * resp.Rate

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just in case
		p.l.Println("[ERROR] serializing product", err)
	}
}
