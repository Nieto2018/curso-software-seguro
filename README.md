# Curso de Ciberseguridad para Desarrollo Web

## Introducción

Implementa buenas prácticas de ciberseguridad en aplicaciones web: protege datos, controla accesos, automatiza pruebas, gestiona secretos y roles con AWS, Auth0, Terraform y GitHub Actions.

URL curso: <https://platzi.com/cursos/software-seguro/>

## Docker

### Crear red

```bash
docker network create --driver=bridge curso-software-seguro-network
```

### Build

```bash
docker build -t curso-software-seguro:latest .
```

### Run

#### Desplegar contenedor aplicación GO

```bash
docker run -it --rm \
  --network curso-software-seguro-network \
  --name curso-software-seguro-app \
  curso-software-seguro:latest
```

#### Desplegar contenedor Ngrok

```bash
docker run -it --rm \
  --network curso-software-seguro-network \
  -e NGROK_AUTHTOKEN=<token> \
  ngrok/ngrok http curso-software-seguro-app:8080/hello
```
