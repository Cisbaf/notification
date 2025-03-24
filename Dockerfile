# Etapa 1: Compilação
FROM golang:1.23.5 AS builder 

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY . .

# Baixa as dependências do Go
RUN go mod download

# Compila a aplicação Go
RUN go build -o main .

# Comando para rodar a aplicação
CMD ["./main"]