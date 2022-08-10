FROM golang:1.19-alpine

RUN mkdir /api
WORKDIR /api
COPY . .

RUN go build -o hlcloud-api main.go
EXPOSE 1337

CMD ["./hlcloud-api"]