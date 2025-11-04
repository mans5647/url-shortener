FROM nginx:alpine

RUN mkdir -p /usr/share/nginx/mysite

COPY . /usr/share/nginx/mysite/
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80