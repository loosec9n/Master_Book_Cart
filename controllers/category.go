package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/utils"
	"encoding/json"
	"log"
	"net/http"
)

func (c Controller) AddCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category models.ProductCategory

		//fetching data ->
		json.NewDecoder(r.Body).Decode(&category)

		//adding category to database ->
		category, err := c.ProductRepo.AddCategory(category)
		if err != nil {
			log.Println("Failed to add category")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Failed to add category", err))
			return
		}

		log.Println("category added by admin")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode((utils.PrepareResponse(true, "Succesfully added category", &category)))

	}
}

func (c Controller) ViewCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//data form -> database
		category, err := c.ProductRepo.ViewCategory()

		if err != nil {
			log.Println("Failed to fetch category from the databse")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "No category found", nil))
		}
		log.Println("Found Category")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "category found", &category))
	}
}

func (c Controller) AddAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.ProductAuthor

		//fetching data
		json.NewDecoder(r.Body).Decode(&author)

		//adding author to database author table
		author, err := c.ProductRepo.AddAuthor(author)

		if err != nil {
			log.Println("error adding author")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error adding author", err))
			return
		}

		log.Println("Successfuly added author")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "sucess adding author", &author))

	}
}

func (c Controller) ViewAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//fetcing data from Author Table
		author, err := c.ProductRepo.ViewAuthor()

		if err != nil {
			log.Println("author not found from the author table")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "author not found", nil))
			return
		}

		log.Println("Sucessfully found author")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "author found", &author))
	}
}
