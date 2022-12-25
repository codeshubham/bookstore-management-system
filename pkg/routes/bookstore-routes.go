package routes

import (
	"github.com/codeshubham/bookstore-management-system/pkg/controllers"
	"github.com/gorrila/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Method("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")

}
