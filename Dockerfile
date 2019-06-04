FROM openshift/origin-release:golang-1.12 as builder
WORKDIR /go/src/github.com/aerogear/mobile-security-service
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
&& dep ensure \
&& go build -o mobile-security-service ./cmd/mobile-security-service/main.go

FROM node:10 as ui
WORKDIR /go/src/github.com/aerogear/mobile-security-service
COPY . .
RUN cd ui/ && npm install --production && npm run build

FROM centos:7
COPY --from=builder /go/src/github.com/aerogear/mobile-security-service /usr/bin/
COPY --from=ui /go/src/github.com/aerogear/mobile-security-service/ui/build /opt/mobile-security-service/ui/build/

USER 1001
ENV STATIC_FILES_DIR /opt/mobile-security-service/ui/build
EXPOSE 3000
ENTRYPOINT ["/usr/bin/mobile-security-service"]