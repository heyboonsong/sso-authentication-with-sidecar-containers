FROM rust:1.71.1 as builder

WORKDIR /builder

COPY . .
ENV TODO_URL=http://todo-service
ENV USER_URL=http://user-service
RUN cargo build -p web-bff --release
# RUN cargo zigbuild --release --target x86_64-unknown-linux-musl

FROM gcr.io/distroless/cc

COPY --from=builder /builder/target/release/web-bff /
CMD ["./web-bff"]