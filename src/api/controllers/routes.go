package controllers

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.GET("/", s.Home)

	// // Login Route
	s.Router.POST("/login", s.Login)

	// //Users routes
	s.Router.POST("/users", s.CreateUser)
	s.Router.GET("/users", s.GetUsers)
	s.Router.GET("/users/:id", s.GetUser)
	s.Router.PUT("/users/:id", s.UpdateUser)
	s.Router.POST("/users/:id", s.DeleteUser)

	// //Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
