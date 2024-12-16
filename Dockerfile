FROM golang:1.23

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-multiplication

EXPOSE 8080

CMD [ "/go-multiplication" ]
