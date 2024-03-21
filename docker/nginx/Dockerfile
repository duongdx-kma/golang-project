# Use the official Nginx image based on Alpine Linux
FROM nginx:alpine

# Remove the default Nginx configuration
RUN rm -rf /etc/nginx/conf.d/default.conf

# Expose ports 80 and 443
EXPOSE 80 443

# Start Nginx in daemon mode
CMD ["nginx", "-g", "daemon off;"]