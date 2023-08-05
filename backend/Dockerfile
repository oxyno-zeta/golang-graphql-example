FROM alpine:3.18

ENV USER=appuser
ENV APP=golang-graphql-example
ENV UID=1000
ENV GID=1000

RUN apk add --update --no-cache ca-certificates=20230506-r0 && \
    rm -Rf /var/cache/apk/* && \
    addgroup -g $GID $USER && \
    adduser -D -g "" -h "/$USER" -G "$USER" -H -u "$UID" "$USER"

WORKDIR /$USER

COPY $APP /$USER/$APP

RUN chown -R $UID:$GID /$USER

USER $USER

ENTRYPOINT [ "/appuser/golang-graphql-example" ]
