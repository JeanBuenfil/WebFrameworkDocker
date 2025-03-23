package main

//Import para usar el framework gin
import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Endpoint de alumnos
	r.GET("/students", func(c *gin.Context) {
		students := []map[string]string{
			{"nombre": "Gloria Maria Ferrero", "matricula": "A01"},
			{"nombre": "Fuensanta Rodríguez", "matricula": "A02"},
			{"nombre": "Benita Baez", "matricula": "A03"},
			{"nombre": "Carlos Javier Vergara", "matricula": "A04"},
			{"nombre": "Jean Buenfil", "matricula": "A05"},
		}
		c.JSON(200, students)
	})

	// Endpoint de profesores
	r.GET("/teachers", func(c *gin.Context) {
		teachers := []map[string]string{
			{"nombre": "Michael Zurita", "numeroEmpleado": "E01"},
			{"nombre": "Maria Elena Gaspar", "numeroEmpleado": "E02"},
			{"nombre": "Rita Carbonell", "numeroEmpleado": "E03"},
			{"nombre": "Maria Luisa Bustos", "numeroEmpleado": "E04"},
			{"nombre": "Manuel Mesa", "numeroEmpleado": "E05"},
			{"nombre": "Leroy Jenkins", "numeroEmpleado": "E06"},
		}
		c.JSON(200, teachers)
	})

	// Iniciar servidor en el puerto 80
	r.Run(":80")
}
