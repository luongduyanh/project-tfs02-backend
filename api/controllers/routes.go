package controllers

import "project-tfs02/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Search Route
	//tim kiem product theo product name
	s.Router.HandleFunc("/api/products/search/{name}", middlewares.SetMiddlewareJSON(s.SearchProductsByName)).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Orders routes ------------------ can middleware checkauth
	//tao 1 order moi
	s.Router.HandleFunc("/api/orders", middlewares.SetMiddlewareJSON(s.CreateOrder)).Methods("POST")
	//tat ca order - admin
	s.Router.HandleFunc("/api/orders", middlewares.SetMiddlewareJSON(s.GetOrders)).Methods("GET")
	//tat ca order theo userID
	s.Router.HandleFunc("/api/orders/", middlewares.SetMiddlewareJSON((s.GetOrdersByUserID))).Methods("GET")
	//tat ca orderlines trong 1 order
	s.Router.HandleFunc("/api/orders/{id}/order-lines", middlewares.SetMiddlewareJSON((s.GetOrderLinesByOrderID))).Methods("GET")
	
	//Posts routes
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//Products routes
	//tao 1 product moi
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateProduct))).Methods("POST")
	//lay tat ca product
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	//lay product theo ID
	s.Router.HandleFunc("/api/products/{id}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")
	//update product theo ID
	s.Router.HandleFunc("/api/products/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PUT")
	//delete product theo ID
	s.Router.HandleFunc("/api/products/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteProduct))).Methods("DELETE")

	//Brands routes
	//lay tat ca brands
	s.Router.HandleFunc("/api/brands", middlewares.SetMiddlewareJSON(s.GetBrands)).Methods("GET")
	//lay brand theo ID
	s.Router.HandleFunc("/api/brands/{id}", middlewares.SetMiddlewareJSON(s.GetBrand)).Methods("GET")
	//lay tat ca products cua brand
	s.Router.HandleFunc("/api/brands/{id}/products", middlewares.SetMiddlewareJSON(s.GetProductsByBrandID)).Methods("GET")

	//Categories routes
	//lay tat ca categories
	s.Router.HandleFunc("/api/categories", middlewares.SetMiddlewareJSON(s.GetCategories)).Methods("GET")
	//lay category theo ID
	s.Router.HandleFunc("/api/categories/{id}", middlewares.SetMiddlewareJSON(s.GetCategory)).Methods("GET")
	//lay tat ca sub-cates cua category
	s.Router.HandleFunc("/api/categories/{id}/sub-categories", middlewares.SetMiddlewareJSON(s.GetSubCatesByCateID)).Methods("GET")

	//SubCategories routes
	//lay tat ca subcates
	s.Router.HandleFunc("/api/sub-categories", middlewares.SetMiddlewareJSON(s.GetSubCategories)).Methods("GET")
	//lay subcate theo ID
	s.Router.HandleFunc("/api/sub-categories/{id}", middlewares.SetMiddlewareJSON(s.GetSubCategory)).Methods("GET")
	//lay tat ca products cua sub-cate
	s.Router.HandleFunc("/api/sub-categories/{id}/products", middlewares.SetMiddlewareJSON(s.GetProductsBySubCategoryID)).Methods("GET")
}
