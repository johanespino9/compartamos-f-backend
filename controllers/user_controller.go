package controllers

import (
	"compartamos-backend/models"
	"compartamos-backend/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserService services.UserService
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    var input models.User
    

    fmt.Printf("Input recibido: %+v\n", input)    
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Campos inválidos"})
        return
    }

    // Validar y parsear la fecha de nacimiento
    birthDate, err := time.Parse("2006-01-02", input.BirthDate)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Fecha de nacimiento inválida. Formato esperado: YYYY-MM-DD"})
        return
    }

    // Calcular la edad
    var age = calculateAge(birthDate)
    if age < 18 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario debe ser mayor de 18 años"})
        return
    }

    input.Age = age

    user, err := ctrl.UserService.CreateUser(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la creación de usuario"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
    users, err := ctrl.UserService.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
        return
    }

    c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := ctrl.UserService.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Campos inválidos"})
        return
    }

    updatedUser, err := ctrl.UserService.UpdateUser(id, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar datos del usuario"})
        return
    }

    c.JSON(http.StatusOK, updatedUser)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")

    err := ctrl.UserService.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}

func calculateAge(birthDate time.Time) int {
    today := time.Now()
    age := today.Year() - birthDate.Year()
    
    if today.YearDay() < birthDate.YearDay() {
        age--
    }
    return age
}
