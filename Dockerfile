FROM golang:alpine AS build-env

RUN apk add --no-cache git

COPY . /src
WORKDIR /src
RUN go get -d ./...
RUN go build -o /assets/out ./out
RUN cp ./check/check /assets/check
RUN cp ./in/in /assets/in

FROM alpine
RUN set -ex \
      && apk add --no-cache ca-certificates

RUN mkdir -p /opt/resource
COPY --from=build-env /assets/check /opt/resource/
COPY --from=build-env /assets/in /opt/resource/
COPY --from=build-env /assets/out /opt/resource/

CMD echo "This container is not used directly."; exit 1
