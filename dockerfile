FROM --platform=amd64 golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./src ./src
COPY *.go ./

RUN go build -o /mantenimiento
#RUN CGO_ENABLED=0 go build -ldflags=”-s -w” -o /mantenimiento

FROM --platform=amd64 alpine:latest
WORKDIR /root/
COPY --from=builder /mantenimiento /

#EXPOSE 3000

#USER nonroot:nonroot

CMD [ "/mantenimiento", "serve" ]
