FROM golang:1.23rc2-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=80
ENV DOCKER=true

EXPOSE $PORT

CMD [ "go", "run", "cmd/app/main.go" ]