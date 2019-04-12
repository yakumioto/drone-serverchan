FROM alpine:3.9

LABEL maintainer="Mioto <yaku.mioto@gmail.com>"

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /drone/src

COPY serverchan /usr/local/bin/serverchan

ENTRYPOINT ["serverchan"]