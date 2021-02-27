# docker/ui/Dockerfile
FROM node:14.9.0-alpine3.10

CMD ["/bin/sh"]

WORKDIR /ui

RUN apk update && \
    npm install -g http-server
