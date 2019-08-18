
# Build Stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build -o proton cli/main.go

# Runtime
FROM alpine
WORKDIR /app
COPY --from=build-env /src/proton /app/
ENTRYPOINT ./proton serve