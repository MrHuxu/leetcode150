FROM golang:latest AS go-builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0

WORKDIR /work
COPY ./website /work

RUN go mod download
RUN go build -o main *.go

FROM scratch

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./website/data /output/data
COPY ./website/templates /output/templates
COPY ./documents /documents
COPY ./go /go
COPY ./rust /rust
COPY --from=go-builder /work/main /output/

EXPOSE 15050
ENTRYPOINT [ "./main" ]