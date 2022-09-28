ARG image=registry.access.redhat.com/ubi8/ubi-minimal:8.6
FROM $image

ENV HOME=/tmp
WORKDIR /tmp

COPY test-connector /

EXPOSE 8080
USER 10001

ENTRYPOINT ["/test-connector"]
CMD [ "run" ]
