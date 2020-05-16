FROM golang:1.14-alpine AS build

WORKDIR /go/src/github.com/HackDalton/talking-loudly
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine

COPY --from=build /go/bin/talking-loudly ./

CMD ["./talking-loudly"]