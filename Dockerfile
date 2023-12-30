FROM --platform=$BUILDPLATFORM node:14-alpine as JSBUILD
COPY package.json package-lock.json rollup.config.js ./
RUN npm ci
COPY resources/js resources/js
RUN npm run build

FROM  --platform=$TARGETPLATFORM golang:1.21 as GOBUILD
# Force the go compiler to use modules
ENV GO111MODULE=on
# Create the user and group files to run unprivileged 
RUN mkdir /user && \
    echo 'mokintoken:x:65534:65534:mokintoken:/:' > /user/passwd && \
    echo 'mokintoken:x:65534:' > /user/group
#RUN apk update && apk add --no-cache git ca-certificates tzdata  gcc g++  openssh-client

RUN mkdir /build 
COPY . /build/
WORKDIR /build 
# Import the code TODO: isolate to only requried go files
COPY ./ ./ 
RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux go build -o mokintoken .


FROM alpine AS final
LABEL author="John Cena"
# RUN apk add sqlite3 ?
# Import the user and group files
COPY --from=GOBUILD /user/group /user/passwd /etc/
COPY --from=GOBUILD /build/mokintoken /
COPY --from=JSBUILD --chown=mokintoken:mokintoken assets/* /assets/
COPY ./views /views
COPY ./database /database

WORKDIR /

USER mokintoken

EXPOSE 8080
ENTRYPOINT /mokintoken
VOLUME /database/mokintoken.sqilte
