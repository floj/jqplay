FROM golang:1-alpine as builder
ENV CGO_ENABLED=0

RUN apk add --no-cache bash nodejs npm git curl \
	&& npm install --global grunt-cli bower
RUN  curl -Lfo /bin/jq "https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64" \
	&& chmod +x /bin/jq
RUN addgroup -g 10001 scratch \
	&& adduser -h /app -D -G scratch -u 10001 scratch

COPY . /app
RUN chown -R scratch:scratch /app
USER scratch
WORKDIR /app

RUN bin/build
RUN touch /app/jqplay.boltdb

FROM scratch
ENV PORT=8080
ENV DATABASE_FILE=/jqplay.boltdb
ENV PATH=/bin
COPY --from=builder /etc/passwd /etc/group /etc/
COPY --from=builder /bin/jq /bin/jq
COPY --from=builder /app/jqplay /bin/jqplay
COPY --from=builder /app/public /srv/public
COPY --from=builder /app/jqplay.boltdb /jqplay.boltdb
WORKDIR /srv
USER scratch
EXPOSE 8080
CMD ["/bin/jqplay"]