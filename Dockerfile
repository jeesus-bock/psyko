FROM golang:1.23rc1-bookworm
EXPOSE 8080
WORKDIR /psyko
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -v -o /psyko
CMD ["/psyko/psyko"]
