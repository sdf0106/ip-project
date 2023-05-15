FROM golang:1.20-alpine AS builder

RUN mkdir /app

COPY . /app
COPY . /app/configs/config.yml
WORKDIR /app

RUN CGO_ENABLE=0 go build -o ip_project ./cmd/app

RUN chmod +x /app/ip_project

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/ip_project /app
#COPY --from=0 /app/configs/config.yml ./config/

CMD ["./app/ip_project"]