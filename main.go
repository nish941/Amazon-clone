package main

import (
    "github.com/gin-gonic/gin"
    "amazon-clone/routes"
)

func main() {
    router := gin.Default()

    // Initialize routes
    routes.InitializeRoutes(router)

    // Start server
    router.Run(":8080")
}
