FROM golang:1.25-alpine

WORKDIR /electro98/noteapp/

COPY . .
RUN go build -v -o /electro98/noteapp/noteapp

CMD ["/electro98/noteapp/noteapp"]