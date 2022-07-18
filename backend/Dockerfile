FROM golang:1.18.4-alpine3.16

RUN apk add --no-cache git
#Install GCC
RUN apk add build-base

WORKDIR /app/ebiznes_go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./out/ebiznes_go .

EXPOSE 8080

CMD ["./out/ebiznes_go"]