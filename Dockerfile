FROM golang:alpine
ENV CGO_ENABLED 0
RUN apk update --no-cache
RUN apk add git \bash \make
ADD . /go/src/devops-test
WORKDIR /go/src/devops-test
RUN make

FROM alpine:latest  
EXPOSE 8080
WORKDIR /devops-test
COPY --from=0 /go/src/devops-test .
CMD ["./endpoint"] 
