FROM nginx:latest

COPY ./ssl/www.isdl-sentinel.com.pem /etc/nginx/ssl/www.isdl-sentinel.com.pem
COPY ./ssl/www.isdl-sentinel.com.nopass.key /etc/nginx/ssl/www.isdl-sentinel.com.nopass.key

RUN apt-get update && apt-get install -y gettext-base

COPY nginx.conf.template /etc/nginx/nginx.conf.template

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
