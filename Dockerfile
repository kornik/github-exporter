FROM golang:1.20-alpine

WORKDIR /usr/src/github-exporter

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/src/github-exporter/github-exporter
EXPOSE 2112

CMD ["/usr/src/github-exporter/github-exporter"]