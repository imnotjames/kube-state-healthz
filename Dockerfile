ARG GOVERSION=1.22
ARG GOARCH=amd64
ARG GOOS=linux

FROM golang:${GOVERSION} as builder
ENV GOARCH=${GOARCH}
ENV GOOS=${GOOS}
WORKDIR /go/src/github.com/imnotjames/kube-state-healthz
COPY . /go/src/github.com/imnotjames/kube-state-healthz
RUN CGO_ENABLED=0 go build -o /kube-state-healthz

FROM gcr.io/distroless/static:latest-${GOARCH} as runtime
COPY --from=builder /kube-state-healthz /
USER nobody
ENTRYPOINT ["/kube-state-healthz"]
CMD ["serve"]
EXPOSE 8000
