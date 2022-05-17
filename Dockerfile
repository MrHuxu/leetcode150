FROM rust:latest AS builder

RUN rustup target add x86_64-unknown-linux-musl
RUN apt update && apt install -y musl-tools musl-dev
RUN update-ca-certificates

WORKDIR /website
COPY ./website/src ./src
COPY ./website/templates ./templates
COPY ./website/Cargo.toml ./Cargo.toml
COPY ./website/Cargo.lock ./Cargo.lock

RUN cargo build --target x86_64-unknown-linux-musl --release

FROM scratch

WORKDIR /website
COPY --from=builder /website/target/x86_64-unknown-linux-musl/release/website ./
COPY ./website/src/data.json ./src/data.json
COPY ./documents /documents
COPY ./go /go
COPY ./rust/src /rust/src
COPY ./java/src /java/src

EXPOSE 15050
ENTRYPOINT [ "/website/website" ]
