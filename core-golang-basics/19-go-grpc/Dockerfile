FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o mygrpc .

FROM alpine:latest
COPY --from=build /app/mygrpc /usr/local/bin/
RUN apk add --no-cache ca-certificates
COPY grpc-health-probe /bin/
EXPOSE 50051
HEALTHCHECK --interval=10s --timeout=3s \
  CMD grpc-health-probe -addr=localhost:50051 || exit 1
CMD ["mygrpc"]
