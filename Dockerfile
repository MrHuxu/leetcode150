FROM node:11.15.0 AS node-builder

ENV NODE_ENV production

WORKDIR /work
COPY ./website/client /work/client
COPY ./website/data.json /work/
COPY ./website/package.json /work/
COPY ./website/package-lock.json /work/
COPY ./website/config/webpack.config.js /work/config/

RUN npm install
RUN ./node_modules/webpack/bin/webpack.js --config config/webpack.config.js

FROM golang:latest AS go-builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
ENV CGO_ENABLED 0

WORKDIR /work
COPY ./website/main.go /work/
COPY ./website/server /work/server
COPY ./website/go.mod /work/
COPY ./website/go.sum /work/

RUN go mod download
RUN go build main.go

FROM scratch

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./website/config/server.json /output/config/
COPY ./website/server/templates /output/server/templates
COPY ./problems /problems
COPY ./website/data.json /output/
COPY --from=node-builder /work/client/public/bundle.js /output/client/public/
COPY --from=go-builder /work/main /output/

EXPOSE 15050
ENTRYPOINT [ "./main" ]