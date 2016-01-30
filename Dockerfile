# Ce Dockerfile expose la commande interactive de bilan retraite.
FROM scratch

MAINTAINER "Stève Sfartz" <steve.sfartz@gmail.com>

# Copies the directory in which the docker build command is launched
COPY . /deploy

ENTRYPOINT ["/deploy/maretraite"]



