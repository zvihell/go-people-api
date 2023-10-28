FROM golang:1.20-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o go-people-api ./cmd/main.go

CMD [ "./go-people-api" ]