FROM nginx:alpine

COPY nginx.conf.template /etc/nginx

COPY entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]