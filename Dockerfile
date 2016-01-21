FROM scratch

MAINTAINER "Stève Sfartz" <steve.sfartz@gmail.com>

COPY . /deploy

EXPOSE 8080

ENTRYPOINT ["/deploy/retraite", "--port=8080"]



