ARG GOVERSION=1.22
ARG GOARCH

FROM golang:${GOVERSION} as builder
ENV GOARCH=${GOARCH}
WORKDIR /go/src/github.com/imnotjames/kube-state-healthz
COPY . /go/src/github.com/imnotjames/kube-state-healthz
RUN go build -o kube-state-healthz

FROM gcr.io/distroless/static:latest-${GOARCH} as runtime
COPY --from=builder /go/src/github.com/imnotjames/kube-state-healthz/kube-state-healthz /
USER nobody
ENTRYPOINT ["/kube-state-healthz"]
EXPOSE 8000
