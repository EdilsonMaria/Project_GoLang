# Usar a imagem oficial do Golang como base
FROM golang:1.21

# Definindo o diretório de trabalho dentro do container
WORKDIR /app

# Copiando o código-fonte para o container
COPY . .

# Construindo o binário
RUN go build -o main .

# Expondo a porta 8080
EXPOSE 8080

# Comando para iniciar o aplicativo
CMD ["./main"]
