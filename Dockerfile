
FROM golang:alpine

ENV GO111MODULE=on
ENV API_PORT=8888

WORKDIR /app
COPY . ./
RUN go build -o /go-api

CMD [ "/go-api" ]