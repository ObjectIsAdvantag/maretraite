FROM scratch

MAINTAINER "Stève Sfartz" <steve.sfartz@gmail.com>

COPY . /retraite

EXPOSE 8080

ENTRYPOINT ["/retraite/bilan", "--port=8080", "-logtostderr=true", "-v=5"]



