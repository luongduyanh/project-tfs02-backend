package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	es "project-tfs02/api/es_client"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"strings"

	"github.com/gorilla/mux"
	elastic "github.com/olivere/elastic/v7"
)

func (server *Server) SearchProductsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ReplaceAll(vars["name"], "-", " ")
	product := models.Product{}
	productGotten, err := product.FindProductsByName(server.DB, name)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, productGotten)

}

type ESClient struct {
	*elastic.Client
}

type BookManager struct {
	esClient *ESClient
}

func NewBookManager(es *ESClient) *BookManager {
	return &BookManager{esClient: es}
}

func (bm *BookManager) SearchBooks(name string) []*models.Product {
	ctx := context.Background()
	if bm.esClient == nil {
		fmt.Println("Nil es client")
		return nil
	}
	// build query to search for title
	query := elastic.NewSearchSource()
	query.Query(elastic.NewMatchQuery("name", name))

	// get search's service
	searchService := bm.esClient.
		Search().
		Index("products").
		SearchSource(query)

	// perform search query
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Cannot perform search with ES", err)
		return nil
	}
	// get result
	var products []*models.Product

	for _, hit := range searchResult.Hits.Hits {
		var product models.Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			fmt.Println("Get data error: ", err)
			continue
		}
		products = append(products, &product)
	}
	return products
}

func (server *Server) EsSearchByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := string(vars["name"])
	url := "http://localhost:9200"
	esclient, _ := es.NewESClient(url)

	// search
	bm := NewBookManager((*ESClient)(esclient))
	productGotten := bm.SearchBooks(name)
	utils.JSON(w, http.StatusOK, productGotten)
}
