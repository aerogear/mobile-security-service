FROM centos:7
ARG BINARY=./mobile-security-service-server
EXPOSE 3000

COPY ${BINARY} /opt/mobile-security-service-server
ENTRYPOINT ["/opt/mobile-security-service-server"]