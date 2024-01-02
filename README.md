# mokintoken

![](https://dockeri.co/image/nexusuw/mokintoken)

## what

a clientside encrypted note sharing webapp. built with php7, Lumen, sqlite, rollup, and docker.

[announcement blog post](https://ramsay.xyz/2020/03/27/mokintoken-released.html)

### 2024 update

using golang as server -> TODO blog

## where

[https://mokintoken.ramsay.xyz](https://mokintoken.ramsay.xyz/?ref=readme)

[onion](http://mokinan4qvxi4ragyzgkewrmnnqslkcdglk6v5zruknwnnuvv2lu5uad.onion/)

## how to self host

```
docker run -p 8080:8080 -v `pwd`/mokintoken.sqilte:/database/mokintoken.sqilte nexusuw/mokintoken
```

## todo

- csrf
- slim down docker container size
- password protect notes (double encryption but stops someone from stumbling onto the contents if they just have the url)
- ratelimit
- allow uploading of images/files
- inline CSP to html
- load test
- generate sha-BLAH during build and include in csp
