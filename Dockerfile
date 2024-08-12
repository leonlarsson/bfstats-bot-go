FROM golang:1.22-bookworm AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy source files, including assets and localazy config
COPY . .

# Tell Railway that I need the LOCALAZY_READ_KEY secret
ARG LOCALAZY_READ_KEY

# Install Localazy CLI
RUN curl -sS https://dist.localazy.com/debian/pubkey.gpg | gpg --dearmor -o /etc/apt/trusted.gpg.d/localazy.gpg && echo 'deb [arch=amd64 signed-by=/etc/apt/trusted.gpg.d/localazy.gpg] https://maven.localazy.com/repository/apt/ stable main' | tee /etc/apt/sources.list.d/localazy.list && apt-get update && apt-get install -y localazy

# Download translations to assets/locales
RUN localazy download -r $LOCALAZY_READ_KEY

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bfstats-go

# Use a minimal image to actually run the binary
FROM scratch

# Copy the binary from the build stage
COPY --from=build /bfstats-go /bfstats-go

# Copy the assets and config file to the final image (this has the downloaded translations)
COPY --from=build /app/assets /assets

# Run
CMD ["/bfstats-go"]