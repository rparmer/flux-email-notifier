FROM --platform=${BUILDPLATFORM} golang:alpine AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /build/
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o dist/flux-email-notifier .

FROM --platform=${TARGETPLATFORM} alpine
WORKDIR /app/
COPY --from=builder /build/dist/flux-email-notifier /app/flux-email-notifier
RUN adduser --disabled-password --uid 1000 -D fluxuser
RUN chown fluxuser:fluxuser /app/flux-email-notifier
USER fluxuser
ENTRYPOINT ["/app/flux-email-notifier"]
