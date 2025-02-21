FROM gcr.io/distroless/static-debian12

WORKDIR /workspace

COPY landns-agent ./app

ENTRYPOINT ["./app"]