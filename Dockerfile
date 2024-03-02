FROM golang:1.22.0-bookworm as builder
WORKDIR /bytesize
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /bytesize/bytesize

FROM scratch
WORKDIR /app
COPY --from=builder /bytesize/bytesize /app/bytesize
EXPOSE 8080
ENV APP_ENV=production
ENTRYPOINT [ "/app/bytesize" ]



