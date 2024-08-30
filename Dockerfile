FROM golang:1.18.3-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17

CMD ["go", "run", "main.go"]
