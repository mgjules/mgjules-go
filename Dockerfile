FROM golang:1.19-alpine AS builder

# Add git, curl and upx support
RUN apk add --no-cache git curl upx gcc g++

WORKDIR /src

# Pull modules
COPY go.* ./
RUN go mod download

# Copy code into image
COPY . ./

# Build application for deployment
RUN go build -tags=jsoniter -trimpath -ldflags '-s -w -linkmode external -extldflags "-static"' -o /tmp/myspace .

# Compress binary
RUN upx --best --lzma /tmp/myspace

# Create minimal image
FROM scratch
COPY --from=builder /tmp/myspace /myspace
CMD ["/myspace"]
