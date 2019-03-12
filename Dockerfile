FROM openshift/origin-release:golang-1.10 as builder

WORKDIR /go/src/github.com/aerogear/mobile-security-service
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
&& dep ensure \
&& go build -o mobile-security-service ./cmd/mobile-security-service/main.go

FROM centos:7

COPY --from=builder /go/src/github.com/aerogear/mobile-security-service /usr/bin/

EXPOSE 3000

USER 1001

ENTRYPOINT ["/usr/bin/mobile-security-service"]
