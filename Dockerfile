# Usar una imagen base de Go
FROM golang:1.24.1

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos del proyecto
COPY . .

# Descargar dependencias
RUN go mod tidy

# Compilar la aplicación
RUN go build -o sicei main.go

# Exponer el puerto 80
EXPOSE 80

# Ejecutar la aplicación
CMD ["./sicei"]
