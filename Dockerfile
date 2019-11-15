FROM node:10.16.3-slim AS builder

ARG VERSION=3.2.3
RUN npm install gitbook-cli -g && gitbook fetch ${VERSION}

WORKDIR /work
COPY . /work

RUN node scripts/merge.js
RUN sed -i -e 's/solution/index/g' SUMMARY.md
RUN rm .gitignore
RUN gitbook build

FROM nginx:alpine

WORKDIR /usr/share/nginx/html
COPY --from=builder /work/_book /usr/share/nginx/html

EXPOSE 80