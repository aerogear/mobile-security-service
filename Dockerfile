FROM centos:7
ARG BINARY=./mobile-security-service
EXPOSE 3000

COPY ${BINARY} /opt/mobile-security-service
ENTRYPOINT ["/opt/mobile-security-service"]