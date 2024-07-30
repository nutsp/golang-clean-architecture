package app

func (app *App) InitRoute() {
	api := app.echo.Group("/api")

	v1 := api.Group("/v1")
	v1.POST("/users", app.userHandler.CreateUserHandler)
	v1.PUT("/users", app.userHandler.UpdateUserHandler)
	v1.GET("/users/:id", app.userHandler.GetUserInfoHandler)
}
