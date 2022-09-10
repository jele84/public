FROM centos:7

RUN yum install -y epel-release && yum install -y nginx && yum clean all
RUN echo "daemon off;" >> /etc/nginx/nginx.conf
COPY index.html /usr/share/nginx/html

CMD [ "nginx" ]
