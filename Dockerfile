# Fase de compilación
FROM golang:1.23.4-alpine as builder
 
WORKDIR /app
 
# Copia los archivos de dependencias y descarga módulos
COPY go.mod go.sum ./
RUN go mod tidy
 
# Copia todo el proyecto (incluyendo .env localmente, si existe)
COPY . .
 
# Compila el binario
RUN go build -o main .
 
# Si existe .env, se copia a env.file; si no, se crea un archivo vacío.
RUN if [ -f .env ]; then cp .env env.file; else touch env.file; fi
 
# Fase final
FROM alpine:latest
 
# Instala dependencias necesarias
RUN apk --no-cache add postgresql-client
 
WORKDIR /root/
 
# Copia el binario compilado
COPY --from=builder /app/main .
 
# Copia el archivo env.file (lo renombramos a .env en la imagen final)
COPY --from=builder /app/env.file .env
 
EXPOSE 8088
 
CMD ["./main"]