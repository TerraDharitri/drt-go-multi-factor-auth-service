FROM golang:1.23.6 AS builder

MAINTAINER DharitrI

WORKDIR /terradharitri
COPY . .

WORKDIR /terradharitri/cmd/multi-factor-auth

RUN go build -o tcs

# ===== SECOND STAGE ======
FROM ubuntu:20.04
COPY --from=builder /terradharitri/cmd/multi-factor-auth /terradharitri

EXPOSE 8080

WORKDIR /terradharitri

ENTRYPOINT ["./tcs"]
CMD ["--log-level", "*:DEBUG", "--start-swagger-ui"]
