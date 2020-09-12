FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 8085
RUN go build -o main .
CMD ["/app/main"]
