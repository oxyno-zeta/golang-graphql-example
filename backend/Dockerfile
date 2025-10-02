FROM alpine:3.22

ENV USER=appuser
ENV APP=golang-graphql-example
ENV UID=1000
ENV GID=1000

# renovate: datasource=repology depName=alpine_3_22/ca-certificates versioning=loose
ENV CA_CERTIFICATES_VERSION="20250911-r0"

RUN apk add --update --no-cache ca-certificates=${CA_CERTIFICATES_VERSION} && \
    rm -Rf /var/cache/apk/* && \
    addgroup -g $GID $USER && \
    adduser -D -g "" -h "/$USER" -G "$USER" -H -u "$UID" "$USER"

WORKDIR /$USER

COPY $APP /$USER/$APP

RUN chown -R $UID:$GID /$USER

USER $USER

ENTRYPOINT [ "/appuser/golang-graphql-example" ]
