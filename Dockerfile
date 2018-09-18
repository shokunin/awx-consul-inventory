# Note you cannot run golang binaries on Alpine directly
FROM            debian:buster-slim

MAINTAINER      chris.mague@shokunin.co

COPY            awx-consul-inventory /awx-consul-inventory

WORKDIR		/
ENV		GIN_MODE=release

EXPOSE          8080

ENTRYPOINT      [ "/awx-consul-inventory" ]
