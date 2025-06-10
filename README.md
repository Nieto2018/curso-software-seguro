# Curso de Ciberseguridad para Desarrollo Web

<https://platzi.com/cursos/software-seguro/>

## Docker

### Crear red

```bash
# docker network create --driver=bridge --subnet=192.170.1.0/24 --ip-range=192.170.1.0/24 --gateway=192.170.1.254  gitlab-ci
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

docker run -it --rm \
  --network curso-software-seguro-network \
  --name curso-software-seguro-app \
  -p 8080:8080 \
  curso-software-seguro:latest
```

#### Desplegar contenedor Ngrok

```bash
docker run -it --rm \
  --network curso-software-seguro-network \
  -e NGROK_AUTHTOKEN=1wo3JQxO9j3aiOeJep5dm69chSP_737gTXsZ5GuHiJRwtRayu \
  ngrok/ngrok http curso-software-seguro-app:8080/hello
```
# curso-software-seguro
