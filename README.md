# mokintoken

![](https://dockeri.co/image/nexusuw/mokintoken)

## what

a clientside encrypted note sharing webapp. built with php7, Lumen, sqlite, rollup, and docker.

[(todo) announcement blog post]()

## where

[https://mokintoken.ramsay.xyz](https://mokintoken.ramsay.xyz/?ref=readme)

[(todo) onion]()

## how to self host

```
touch database.sqlite
docker run -v `pwd`/database.sqlite:/var/www/database/database.sqlite nexusuw/mokintoken php artisan migrate
docker run -p 8080:8080 -v `pwd`/database.sqlite:/var/www/database/database.sqlite nexusuw/mokintoken
```

## local setup

1. `sudo apt-get install php-sqlite3 php-mbstring php-7`
2. [install nodejs](https://nodejs.org/en/download/package-manager/)
3. [install composer](https://getcomposer.org/download/)
4. `composer install`
5. `npm install`
6. `php artisan migrate`
7. `npm run build`
8. `php -S 0.0.0.0:8000 -t public`
9. `npm run watch`

## todo

- slim down docker container size
- publish multiple archs for containers
- expire old notes
- password protect notes
- max view count
