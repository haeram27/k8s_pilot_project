FROM alpine:3.16

RUN mkdir -p /src/bin/conf
COPY pilot_server /src/bin/.
COPY conf/* /src/bin/conf/.

WORKDIR /src/bin
ENTRYPOINT ["./pilot_server"]
