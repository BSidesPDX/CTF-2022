FROM nginx:latest
RUN apt-get update \
    && openssl req -x509 -nodes -newkey rsa:4096 -keyout /key.pem -out /cert.pem -sha256 -days 365 -subj "/C=US/ST=Oregon/L=Portland/O=BSides/OU=PDX/CN=definitely-not-clippy"
WORKDIR /app
COPY index.html /usr/share/nginx/html
COPY assets/ /usr/share/nginx/html/assets/
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8000
