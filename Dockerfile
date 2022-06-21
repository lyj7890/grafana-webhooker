FROM alpine

WORKDIR /run
COPY ./bin/webhooker  .

ENTRYPOINT ["/run/webhooker"]
