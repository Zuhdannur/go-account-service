FROM node:20 AS prisma-builder

WORKDIR /app

RUN npm install -g prisma

COPY prisma ./prisma


# --- Build Go Application ---
FROM golang:1.24.1 AS go-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=prisma-builder /app/prisma ./prisma

RUN go run github.com/steebchen/prisma-client-go generate

RUN go build -o main ./cmd/main.go

RUN chmod +x main

EXPOSE 3000

CMD ["./main"]