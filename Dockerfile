FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/fyriyadi/Go-Simple-Crud
RUN cd /build/Go-Simple-Crud && go build

EXPOSE 8080

ENTRYPOINT [ "/build/Go-Simple-Crud/main" ]