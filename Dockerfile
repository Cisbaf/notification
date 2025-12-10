# Etapa 1: Compilação
FROM golang:1.24 AS builder 

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiando requirements
COPY go.mod go.sum ./

# Baixa as dependências do Go
RUN go mod download

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila a aplicação Go
RUN go build -o main .

# Comando para rodar a aplicação
CMD ["./main"]
