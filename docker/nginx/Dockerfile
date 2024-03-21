FROM nginx:alpine

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY nginx.conf.template /etc/nginx
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE 80 443

ENTRYPOINT ["/bin/bash", "-c", "/usr/local/bin/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]