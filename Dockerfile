FROM alpine:latest
COPY gosling /usr/local/bin/
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["gosling", "-help"]
