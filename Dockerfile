FROM golang:alpine AS go-builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0

WORKDIR /work
COPY ./website /work

RUN go mod download
RUN go build -o main *.go

FROM node:alpine AS node-builder

ENV NODE_ENV production

WORKDIR /work

COPY ./website/templates /work/templates
COPY ./website/tailwind.config.js /work/

RUN npx tailwindcss -o ./built.css --minify

FROM scratch

ENV GIN_MODE release

WORKDIR /output
COPY ./website/assets/favicon.png /output/assets/
COPY ./website/templates /output/templates
COPY ./website/data.json /output/data.json
COPY ./documents /documents
COPY ./go /go
COPY ./rust/src /rust/src
COPY ./java/src /java/src
COPY --from=go-builder /work/main /output/
COPY --from=node-builder /work/built.css /output/assets/

EXPOSE 15050
ENTRYPOINT [ "./main" ]