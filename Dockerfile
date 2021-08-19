FROM golang:1.16-alpine3.14

WORKDIR /app
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./

EXPOSE 8080

CMD [ "./spotter-vector" ]