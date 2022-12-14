FROM --platform=$BUILDPLATFORM golang:1.19-alpine AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Add git, curl and upx support
RUN apk add --no-cache git curl upx ca-certificates 

# Get dart-sass-embedded
ARG DART_SASS_EMBEDDED_UPSTREAM_VERSION=1.55.0
ARG DART_SASS_EMBEDDED_TARGETARCH=$TARGETARCH
RUN if [ "$TARGETARCH" = "amd64" ]; then DART_SASS_EMBEDDED_TARGETARCH=x64; fi; \
  wget https://github.com/sass/dart-sass-embedded/releases/download/$DART_SASS_EMBEDDED_UPSTREAM_VERSION/sass_embedded-$DART_SASS_EMBEDDED_UPSTREAM_VERSION-$TARGETOS-$DART_SASS_EMBEDDED_TARGETARCH.tar.gz -P /tmp/ && \
  tar -C /tmp/ -xzvf /tmp/sass_embedded-$DART_SASS_EMBEDDED_UPSTREAM_VERSION-$TARGETOS-$DART_SASS_EMBEDDED_TARGETARCH.tar.gz --strip-components=1  sass_embedded/dart-sass-embedded

WORKDIR /src

# Pull modules
COPY go.* ./
RUN go mod download

# Copy code into image
COPY . ./

# Build application for deployment
RUN --mount=type=cache,target=/root/.cache/go-build \
  --mount=type=cache,target=/go/pkg \
  CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -tags=jsoniter -trimpath -ldflags '-s -w' -o /tmp/myspace .

# Compress binary
RUN upx --best --lzma /tmp/myspace

# Create minimal image
FROM --platform=$TARGETPLATFORM gcr.io/distroless/base

# Add the dart-sass-embedded
COPY --from=builder /tmp/dart-sass-embedded /bin/dart-sass-embedded

# Add the binary
COPY --from=builder /tmp/myspace /myspace

EXPOSE 80/tcp

ENTRYPOINT ["/myspace"]
