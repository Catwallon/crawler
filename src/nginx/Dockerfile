FROM nginx:latest

RUN apt-get update && apt-get install -y \
	wget \
    php-fpm \
    php-mysqli \
    php-gd \
    php-xml \
    php-mbstring \
    && apt-get clean

COPY ./default.conf /etc/nginx/conf.d/
RUN wget https://www.phpmyadmin.net/downloads/phpMyAdmin-latest-all-languages.tar.gz
RUN tar -xvf phpMyAdmin-*-all-languages.tar.gz
RUN mkdir -p /usr/share/nginx/html/phpmyadmin
RUN cp -r phpMyAdmin-*-all-languages/* /usr/share/nginx/html/phpmyadmin
COPY ./config.inc.php /usr/share/nginx/html/phpmyadmin

EXPOSE 80

CMD service php8.2-fpm start && chmod 666 /run/php/php8.2-fpm.sock && nginx -g "daemon off;"