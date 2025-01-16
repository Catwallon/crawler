FROM nginx:latest

RUN apt-get update && apt-get install -y \
	wget \
	php-fpm \
	php-mysqli \
	php-gd \
	php-xml \
	php-mbstring \
	npm \
	&& apt-get clean

ARG API_PORT
ARG API_HOST

ENV VITE_API_HOST=$API_HOST
ENV VITE_API_PORT=$API_PORT

WORKDIR /root

# Copy nginx config
COPY default.conf /etc/nginx/conf.d/

# Build and install frontend
COPY src/ .
RUN npm install && npm run build
RUN cp -R dist/* /usr/share/nginx/html

# Install phpMyAdmin
RUN wget https://www.phpmyadmin.net/downloads/phpMyAdmin-latest-all-languages.tar.gz
RUN tar -xvf phpMyAdmin-*-all-languages.tar.gz
RUN mkdir -p /usr/share/phpmyadmin
RUN cp -r phpMyAdmin-*-all-languages/* /usr/share/phpmyadmin
COPY ./config.inc.php /usr/share/phpmyadmin

EXPOSE 80

CMD env && service php8.2-fpm start && chmod 666 /run/php/php8.2-fpm.sock && nginx -g "daemon off;"