FROM golang:alpine

WORKDIR /src
EXPOSE 5858
ADD . .
RUN go mod download && go build
ENTRYPOINT ["/"]