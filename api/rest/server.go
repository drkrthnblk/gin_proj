package rest

func BuidServer() *gin.Engine {
	server := gin.Default()
	server.Use(gin.CustomRecovery(customisedHandleRecovvery))
	v1 := server.Group("/orders")
}