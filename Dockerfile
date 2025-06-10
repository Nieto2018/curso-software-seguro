# docker pull golang:alpine
FROM golang:alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de tu aplicación
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila tu aplicación Go
RUN go build /app/github-tracker/main.go

CMD ["/app/main"]