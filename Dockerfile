# To debug container set "debug" tag
ARG DEPLOY_TAG=latest

# Build
FROM golang:1.16-alpine3.14 AS build

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /spotter-vector

# Deploy
FROM gcr.io/distroless/base-debian10:$DEPLOY_TAG

WORKDIR /

COPY ./static ./static
COPY --from=build /spotter-vector /spotter-vector

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "./spotter-vector" ]