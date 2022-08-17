package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	Routers()
}

func Routers() {
	InitDB()
	defer db.Close()
	log.Println("Starting the HTTP server on port 9080")
	router := mux.NewRouter()
	router.HandleFunc("/products",
		GetProducts).Methods("GET")
	router.HandleFunc("/products",
		CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}",
		GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}",
		UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}",
		DeleteProduct).Methods("DELETE")
	router.HandleFunc("/categories",
		GetCategories).Methods("GET")
	router.HandleFunc("/products",
		GetProductByCategory).Methods("GET")
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})
}

/***************************************************/

//Get all produits
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	result, err := db.Query("SELECT id, name," +
		"price,description,quantity,id_category from product")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.Name,
			&product.Price, &product.Description,&product.Quantity,&product.IDCategory)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}
//Get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categories []Category
	result, err := db.Query("SELECT id, libelle from category")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var category Category
		err := result.Scan(&category.ID, &category.Libelle)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}
	json.NewEncoder(w).Encode(categories)
}


//Create product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO product(name," +
		"price,description,quantity,id_category) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	price := keyVal["price"]
	description := keyVal["description"]
	quantity := keyVal["quantity"]
	id_category := keyVal["id_category"]
	_, err = stmt.Exec(name, price, description,quantity,id_category)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New produit was created")
}

//Get product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, name,"+
		"price,description,quantity,id_category from product WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var product Product
	for result.Next() {
		err := result.Scan(&product.ID, &product.Name,
			&product.Price, &product.Description,&product.Quantity,&product.IDCategory)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(product)
}
//Get Product by category
func GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT id, name,"+
		"price,description,quantity,id_category from product WHERE id_category = ?", params["id_category"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var product Product
	for result.Next() {
		err := result.Scan(&product.ID, &product.Name,
			&product.Price, &product.Description,&product.Quantity,&product.IDCategory)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(product)
}
//Update product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE product SET name = ?," +
		"price= ?, description=?,quantity=?,id_category=? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	price := keyVal["price"]
	description := keyVal["description"]
	quantity := keyVal["quantity"]
	id_category := keyVal["id_category"]
	_, err = stmt.Exec(name, price, description,quantity,id_category,
		params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "product with ID = %s was updated",
		params["id"])
}

//Delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Product with ID = %s was deleted",
		params["id"])
}

/***************************************************/

type Product struct {
	ID        string `json:"id"`
	Name string `json:"name"`
	Price  string `json:"price"`
	Description     string `json:"description"`
	Quantity  string `json:"quantity"`
	IDCategory     string `json:"id_category"`

}
type Category struct {
    ID    int `json:"id"`
    Libelle  string `json:"libelle"`
    
}


//Db configuration
var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/gocrud")
	if err != nil {
		panic(err.Error())
	}
}
/*func InitDB() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "gocrudusers"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}*/


/***************************************************/

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
