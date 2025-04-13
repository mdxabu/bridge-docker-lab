FROM alpine:latest

RUN apk add --no-cache iputils bash bind-tools

COPY top-domains.txt /top-domains.txt
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
