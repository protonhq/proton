
# Build Stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
RUN mkdir /src
WORKDIR /src
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .
RUN go build -o proton cli/main.go

# Runtime
FROM alpine
WORKDIR /app
COPY --from=build-env /src/proton /app/
ENTRYPOINT ./proton serve