FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get -u github.com/gofiber/fiber/v2

COPY *.go ./

RUN go build -o /fast-http-todo


EXPOSE 3000

CMD [ "/fast-http-todo" ]