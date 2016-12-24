FROM nginx:1.11-alpine

COPY tools/nginx.conf /etc/nginx/nginx.conf
