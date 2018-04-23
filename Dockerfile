FROM golang:alpine AS build-env

ARG pkg=discord-resource

COPY . $GOPATH/src/$pkg

RUN set -ex \
      && apk add --no-cache --virtual .build-deps \
              git \
      && go get -v $pkg/... \
      && apk del .build-deps

RUN go install $pkg/...

FROM alpine
RUN set -ex \
      && apk add --no-cache ca-certificates

RUN mkdir -p /opt/resource
COPY --from=build-env /go/bin/check /opt/resource/
COPY --from=build-env /go/bin/in /opt/resource/
COPY --from=build-env /go/bin/out /opt/resource/

CMD echo "This container is not used directly."; exit 1