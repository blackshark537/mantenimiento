FROM --platform=amd64 golang:alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o ./mantenimiento

FROM --platform=amd64 golang:alpine
WORKDIR /root/
COPY --from=builder /bin /

CMD [ "/mantenimiento", "serve" ]
