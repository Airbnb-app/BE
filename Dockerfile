FROM golang:1.19-alpine3.16

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o airbnb-app

CMD ["./airbnb-app"]