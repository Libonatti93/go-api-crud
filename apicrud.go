package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// Estrutura do modelo de dados (ex: Produto)
type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

var products []Product
var nextID = 1

// Função para listar todos os produtos
func getProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

// Função para obter um produto por ID
func getProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for _, product := range products {
        if product.ID == id {
            c.JSON(http.StatusOK, product)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Função para criar um novo produto
func createProduct(c *gin.Context) {
    var newProduct Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newProduct.ID = nextID
    nextID++
    products = append(products, newProduct)
    c.JSON(http.StatusCreated, newProduct)
}

// Função para atualizar um produto
func updateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for i, product := range products {
        if product.ID == id {
            if err := c.ShouldBindJSON(&products[i]); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            products[i].ID = id // Garantir que o ID não seja alterado
            c.JSON(http.StatusOK, products[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Função para deletar um produto
func deleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for i, product := range products {
        if product.ID == id {
            products = append(products[:i], products[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func main() {
    r := gin.Default()

    // Endpoints da API
    r.GET("/products", getProducts)
    r.GET("/products/:id", getProductByID)
    r.POST("/products", createProduct)
    r.PUT("/products/:id", updateProduct)
    r.DELETE("/products/:id", deleteProduct)

    // Inicia o servidor na porta 8080
    r.Run(":8080")
}
