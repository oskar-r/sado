FROM golang:1.11.1-alpine3.8 AS build
WORKDIR /go/src/my-archive/backend
RUN apk update && apk add curl git ca-certificates && rm -rf /var/cache/apk/*
COPY Gopkg.lock Gopkg.lock 
COPY Gopkg.toml Gopkg.toml
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure -vendor-only
COPY . .
WORKDIR /go/src/my-archive/backend/api
RUN go build -o /go/bin/backend

FROM alpine
WORKDIR /usr/bin
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build /go/bin .
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=build /go/src/my-archive/backend/authz .
COPY --from=build /go/src/my-archive/backend/test-certs .
COPY --from=build /go/src/my-archive/backend/test-certs/mino-public.crt /usr/local/share/ca-certificates/mycert.crt
RUN update-ca-certificates