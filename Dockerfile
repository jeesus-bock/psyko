FROM golang:1.23rc1-alpine3.20
WORKDIR /psyko
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /psyko
CMD ["/psyko/psyko"]
