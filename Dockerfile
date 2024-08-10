# syntax = docker/dockerfile:1.2

FROM golang:1.21.5-bullseye AS dep

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

FROM dep as builder

COPY . .

WORKDIR /app

RUN --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p bin \
    && go build -ldflags="-X main.Version=${VERSION}" -o ./bin/exchange-backend ./cmd/

FROM golang:1.21.5-bullseye

ENV GIN_MODE release
ENV DEPLOYMENT_ENVIRONMENT production

RUN mkdir -p /usr/local/bin
COPY --from=builder /app/bin/exchange-backend /usr/local/bin/

CMD ["exchange-backend"]
