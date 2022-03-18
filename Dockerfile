FROM golang:1.17

# ENV GOPROXY https://goproxy.cn
ENV GOPATH /opt/gopath
ENV DIR /playfair-server
WORKDIR $DIR
ENV PATH="${GOPATH}/bin:${PATH}"

COPY go.mod go.sum . $DIR/
RUN go get -d github.com/astaxie/beego && go get github.com/beego/bee
RUN go mod download && go mod verify
RUN chmod +x playfair/playfair

EXPOSE 8080

CMD bee run playfair-server

# docker cp . $(docker inspect -f '{{.Id}}' playfair-server):/playfair-server
