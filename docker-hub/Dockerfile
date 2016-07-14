# Since we didn't build this image on our machine, Docker reaches out to Dockerhub and grabs it
# https://hub.docker.com/r/library/golang/
FROM golang:1.6

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

RUN go-wrapper download
RUN go-wrapper install

ENV PORT 8080
EXPOSE 8080

CMD go-wrapper run
