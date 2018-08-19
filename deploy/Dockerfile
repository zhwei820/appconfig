FROM ubuntu:16.04
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update && apt-get -y install supervisor

RUN mkdir -p /data
WORKDIR /data

ADD ./beego_api /data
ADD ./supervisord.conf /etc/supervisord.conf
ONBUILD ADD ./beego_api /data
ONBUILD ADD ./supervisord.conf /etc/supervisord.conf


EXPOSE 8080