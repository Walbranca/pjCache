FROM golang


WORKDIR /app/scripts
COPY scripts .
RUN sh go-get.sh

WORKDIR /go/src
COPY .temp/private-repositories .

WORKDIR /app/src
ENTRYPOINT [ "go", "run", "." ]

EXPOSE 80
