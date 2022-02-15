FROM registry.access.redhat.com/ubi8/ubi-minimal

ENV HOME=/tmp
WORKDIR /tmp

COPY write-asset-client /
USER 1001

ENTRYPOINT ["/write-asset-client"]
