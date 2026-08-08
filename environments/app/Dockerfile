FROM golang:1.27rc2-alpine3.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY *.go ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /agent_app

FROM gcr.io/distroless/static-debian13 AS runner
COPY --from=builder /agent_app /agent_app
EXPOSE 8080
USER nonroot:nonroot
CMD ["/agent_app"]
