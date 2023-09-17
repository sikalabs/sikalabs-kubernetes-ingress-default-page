FROM golang:1.20 as build
WORKDIR /workdir
COPY go.* main.go index.html ./
COPY ./static ./static
RUN CGO_ENABLED=0 go build

FROM debian:12-slim
COPY --from=build \
  /workdir/sikalabs-kubernetes-ingress-default-page \
  /usr/local/bin/sikalabs-kubernetes-ingress-default-page
CMD [ "/usr/local/bin/sikalabs-kubernetes-ingress-default-page" ]
EXPOSE 8000
