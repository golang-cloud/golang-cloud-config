FROM hub.skyinno.com/common/alpine:latest
MAINTAINER FAE Config Server "fae@fiberhome.com"
EXPOSE 8888
ADD golang-cloud-config /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/golang-cloud-config"]