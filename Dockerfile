FROM golang:1.22-bookworm

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY *.go ./

# Copy the localazy.json config file into the container
COPY localazy.json ./

# Copy the assets into the container
COPY ./assets ./assets

# Tell Railway that I need the LOCALAZY_READ_KEY secret
ARG LOCALAZY_READ_KEY

# Install Localazy CLI
RUN curl -sS https://dist.localazy.com/debian/pubkey.gpg | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/localazy.gpg && echo 'deb [arch=amd64 signed-by=/etc/apt/trusted.gpg.d/localazy.gpg] https://maven.localazy.com/repository/apt/ stable main' | sudo tee /etc/apt/sources.list.d/localazy.list && sudo apt-get update && sudo apt-get install -y localazy

# Download translations to 
RUN localazy download -r $LOCALAZY_READ_KEY

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bfstats-go

# Run
CMD ["/bfstats-go"]