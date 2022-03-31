FROM rust:latest AS builder

RUN rustup target add x86_64-unknown-linux-musl
RUN apt update && apt install -y musl-tools musl-dev
RUN update-ca-certificates

ENV USER=website
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /website
COPY ./website/src /website/src
COPY ./website/templates /website/templates
COPY ./website/Cargo.toml /website/Cargo.toml
COPY ./website/Cargo.lock /website/Cargo.lock

RUN cargo build --target x86_64-unknown-linux-musl --release

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /website
COPY --from=builder /website/target/x86_64-unknown-linux-musl/release/website ./
COPY ./documents /documents
COPY ./go /go
COPY ./rust/src /rust/src

USER website:website

EXPOSE 15050
ENTRYPOINT [ "/website/website" ]