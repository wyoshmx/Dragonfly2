ARG BASE_IMAGE=alpine:3.12
ARG GOPROXY
ARG GOTAGS
ARG GOGCFLAGS

FROM golang:1.16.6-alpine as server-builder
WORKDIR /go/src/d7y.io/dragonfly/v2
RUN apk --no-cache add bash make gcc libc-dev git
COPY . /go/src/d7y.io/dragonfly/v2
RUN make build-manager && make install-manager

FROM node:12-alpine as console-builder
WORKDIR /build
COPY ./manager/console/package.json /build
RUN npm install --loglevel warn --progress false
COPY ./manager/console /build
RUN npm run build

FROM ${BASE_IMAGE}
WORKDIR /opt/dragonfly/bin
ENV PATH=/opt/dragonfly/bin:$PATH
RUN mkdir -p /opt/dragonfly/bin/manager/console \
    && echo "hosts: files dns" > /etc/nsswitch.conf
COPY --from=server-builder /opt/dragonfly/bin/manager /opt/dragonfly/bin/server
COPY --from=console-builder /build/dist /opt/dragonfly/bin/manager/console/dist
EXPOSE 8080 65003
ENTRYPOINT ["/opt/dragonfly/bin/server"]
