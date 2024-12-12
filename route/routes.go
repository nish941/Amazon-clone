package routes

import (
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
    // User-related routes
    userRoutes := router.Group("/user")
    {
        userRoutes.POST("/signup", nil) // Placeholder
        userRoutes.POST("/login", nil) // Placeholder
    }

    // Product-related routes
    productRoutes := router.Group("/products")
    {
        productRoutes.GET("/", nil) // Placeholder
    }
}
