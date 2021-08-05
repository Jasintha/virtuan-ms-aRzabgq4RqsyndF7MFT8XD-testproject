FROM golang:1.16 as builder

WORKDIR /

COPY . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go mod tidy
RUN swag init
RUN go mod tidy
RUN go build -o app .

CMD ["./app"]


FROM virtuan/alpine-3


WORKDIR /root


COPY --from=builder /app .
COPY --from=builder /start.sh .

EXPOSE $PORT

CMD ["sh", "start.sh"]
