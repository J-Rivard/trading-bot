FROM golang

# Set the Current Working Directory inside the container
WORKDIR /app/trading-bot

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build ./cmd/trading-bot/main.go

EXPOSE 8080

CMD [ "./main" ]