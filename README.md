# mokintoken

![](https://dockeri.co/image/nexusuw/mokintoken)

## what

a clientside encrypted note sharing webapp. built with php7, Lumen, sqlite, rollup, and docker.

[announcement blog post](https://ramsay.xyz/2020/03/27/mokintoken-released.html)

### updated 

using golang as server -> TODO blog

## where

[https://mokintoken.ramsay.xyz](https://mokintoken.ramsay.xyz/?ref=readme)

[onion](http://mokinan4qvxi4ragyzgkewrmnnqslkcdglk6v5zruknwnnuvv2lu5uad.onion/)

## how to self host

```
touch database.sqlite
docker run -v `pwd`/database.sqlite:/var/www/database/database.sqlite nexusuw/mokintoken php artisan migrate
docker run -p 8080:8080 -v `pwd`/database.sqlite:/var/www/database/database.sqlite nexusuw/mokintoken
```

## todo

- csrf
- slim down docker container size
- password protect notes (double encryption but stops someone from stumbling onto the contents if they just have the url)
- ratelimit
- allow uploading of images/files
- turn off server logging?
- inline CSP to html
- load test
