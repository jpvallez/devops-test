FROM golang:alpine
ADD . /go/src/devops-test
WORKDIR /go/src/devops-test
RUN go build -o endpoint  .

FROM alpine:latest  
EXPOSE 8080
WORKDIR /devops-test
COPY --from=0 /go/src/devops-test .
CMD ["./endpoint"] 