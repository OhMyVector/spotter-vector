# To debug container set "debug" tag
ARG DEPLOY_TAG=latest

# Build
FROM golang:1.16-alpine3.14 AS build

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /spotter-vector ./cmd/api/main.go

# Deploy
FROM gcr.io/distroless/base-debian10:$DEPLOY_TAG

WORKDIR /

COPY ./public ./public
COPY conf.local.yaml ./

COPY --from=build /spotter-vector /spotter-vector

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "./spotter-vector" ]