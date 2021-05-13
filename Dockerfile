FROM golang:1.15-alpine AS build
ADD . /app
ENV CGO_ENABLED 0
WORKDIR /app
RUN go build -o replica ./cmd/replica

FROM alpine:3.12
COPY --from=build /app/replica /app/replica
ENTRYPOINT ["/app/replica"]


