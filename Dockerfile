FROM --platform=$BUILDPLATFORM node:14-alpine as JSBUILD
COPY package.json package-lock.json rollup.config.js ./
RUN npm ci
COPY resources/js resources/js
RUN npm run build

FROM  --platform=$TARGETPLATFORM php:7-apache as MAIN

# Use the default production configuration
RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini"

# Copy composer.lock and composer.json
COPY composer.lock composer.json /var/www/

# Set working directory
WORKDIR /var/www

# Install dependencies
RUN apt-get update && apt-get install -y \
  curl git zip unzip

# Install composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

#change to run on 8080 so that non-root user can start the server.
COPY mokintoken.conf /etc/apache2/sites-available/000-default.conf
RUN sed -i 's/80/8080/g' /etc/apache2/ports.conf
# https://www.digitalocean.com/community/questions/why-do-my-laravel-routes-not-work
RUN  a2enmod rewrite

RUN groupadd -g 1000 mokintoken
RUN useradd -u 1000 -ms /bin/bash -g mokintoken mokintoken

COPY --chown=mokintoken:mokintoken . .
RUN composer install
COPY --from=JSBUILD --chown=mokintoken:mokintoken public/* public/
RUN apt-get remove -y curl git zip unzip && apt-get clean && rm -rf /var/lib/apt/lists/*



ENV APACHE_RUN_USER=mokintoken
USER mokintoken

EXPOSE 8080
VOLUME /var/www/database/database.sqilte
