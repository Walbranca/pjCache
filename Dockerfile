FROM golang

WORKDIR /app

COPY go-get.sh /app/go-get.sh
RUN sh go-get.sh

COPY . /app
RUN go build -o app .
CMD [ "./app" ]

EXPOSE 80
