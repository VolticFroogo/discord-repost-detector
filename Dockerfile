FROM golang:1.15.6

WORKDIR /go/src/github.com/VolticFroogo/discord-repost-detector
COPY . .
RUN go build -o main .

CMD ["./main"]
