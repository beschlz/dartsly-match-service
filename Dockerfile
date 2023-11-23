FROM golang:1.21 as builder

COPY . /src
WORKDIR /src

RUN make install
RUN make compile_arm

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=builder /src/bin/dartsly-match-service /dartsly-match-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/dartsly-match-service"]