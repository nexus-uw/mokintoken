FROM --platform=$BUILDPLATFORM node:20-alpine as JSBUILD
COPY package.json package-lock.json rollup.config.js ./
RUN npm ci
COPY resources/js resources/js
RUN npm run build

FROM  --platform=$TARGETPLATFORM golang:1.21-alpine3.18 as GOBUILD

# Important:
#   Because this is a CGO enabled package, you are required to set it as 1.
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

# Force the go compiler to use modules
ENV GO111MODULE=on
ENV GOOS=linux
# Create the user and group files to run unprivileged 
RUN mkdir /user && \
    echo 'mokintoken:x:65534:65534:mokintoken:/:' > /user/passwd && \
    echo 'mokintoken:x:65534:' > /user/group

RUN apk update && apk add --no-cache git ca-certificates tzdata  gcc g++  openssh-client

RUN mkdir /build 
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download


RUN mkdir database && touch database/mokintoken.sqlite

COPY healthcheck ./healthcheck
RUN go build -o healthcheckCommand ./healthcheck/main.go

COPY *.go ./
RUN go get ./

RUN go build -o mokintoken ./mokintoken.go

# FROM scratch AS final would be nice, ned cgo for sqlite3 lib
FROM alpine:3.18 AS final
LABEL author="John Cena"
WORKDIR /app

# Import the user and group files
COPY --from=GOBUILD /user/group /user/passwd /etc/
COPY --from=GOBUILD  /build/mokintoken ./
COPY --from=GOBUILD  /build/healthcheckCommand ./
COPY --from=GOBUILD --chown=mokintoken:mokintoken /build/database ./database
COPY --from=JSBUILD  assets/* ./assets/
COPY ./views ./views

USER mokintoken:mokintoken


EXPOSE 8080
ENV CLEARNET "https://mokintoken.ramsay.xyz"
ENV DARKNET "http://mokinan4qvxi4ragyzgkewrmnnqslkcdglk6v5zruknwnnuvv2lu5uad.onion"
ENTRYPOINT ["/app/mokintoken"]
VOLUME /app/database/mokintoken.sqlite

HEALTHCHECK --interval=30s --timeout=1s --start-period=5s --retries=3 CMD [ "/app/healthcheck" ] 
