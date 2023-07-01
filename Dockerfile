FROM golang:1.20.5-bookworm

WORKDIR /app

COPY . /app

RUN go get .

ENV PORT=":80"

#ENV ENV=prod
CMD ["go", "run", "main.go"]
