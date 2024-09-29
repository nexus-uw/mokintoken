# mokintoken

![](https://dockeri.co/image/nexusuw/mokintoken)

## what

a clientside encrypted note sharing webapp. built with ~~php7, Lumen~~ Go, sqlite, rollup, and docker.

[announcement blog post](https://ramsay.xyz/2020/03/27/mokintoken-released.html)

### 2024 update

[using golang as server](https://ramsay.xyz/2024/02/19/mokintoken-is-now-go.html)

## where

[https://mokintoken.ramsay.xyz](https://mokintoken.ramsay.xyz/?ref=readme)

[onion](http://mokinan4qvxi4ragyzgkewrmnnqslkcdglk6v5zruknwnnuvv2lu5uad.onion/)

## how to self host

```
touch mokintoken.sqlite
chown 65534:65534 mokintoken.sqlite
docker run -p 8080:8080 -v `pwd`/mokintoken.sqlite:/app/database/mokintoken.sqlite -e CLEANET=yoursite.af -e DARKNET=onion nexusuw/mokintoken
```

## local dev
```
npm i && npm run build
touch database/mokintoken.sqlite
go build && ./main

```

## todo

- csrf
- slim down docker container size (more)
- password protect notes (double encryption but stops someone from stumbling onto the contents if they just have the url)
- ratelimit
- allow uploading of ~~images~~/files
- inline CSP to html
- load test
- generate sha-BLAH during build and include in csp
- [web share target](https://developer.mozilla.org/en-US/docs/Web/API/Web_Share_API)
