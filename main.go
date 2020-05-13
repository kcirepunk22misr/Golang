package main

import (
	"database/sql"
	"encoding/json"

	"net/http"
	"github.com/go-chi/chi"
	"GolangRestApi/conexion"
	_ "github.com/go-sql-driver/mysql"
)
var databaseConnection *sql.DB

type Product struct {
	ID int `json:"id"`
	Product_Code string `json:"product_code"`
	Description string `json:description`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	databaseConnection = conexion.InitDB();
	defer databaseConnection.Close()
	r := chi.NewRouter()
	r.Get("/products", AllProduct)
	r.Post("/products", CreateProduct)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProduct)
	http.ListenAndServe(":3000", r)


}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("INSERT products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully created"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("DELETE FROM products WHERE id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var product Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("UPDATE products SET product_code=?, description=? WHERE id=?")
	catch(err)
	_, er := query.Exec(product.Product_Code, product.Description, id)
	catch(er)
	defer query.Close()

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "update successfulli"})
}

func AllProduct(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'') FROM products`;
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	respondWithJSON(w, http.StatusOK, products)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}