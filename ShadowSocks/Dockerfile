FROM ubuntu:12.04
MAINTAINER kanchuanqi@gmail.com
RUN apt-get update
RUN apt-get install -y python-pip
RUN pip install shadowsocks
ENTRYPOINT ["ssserver", "-p", "8000", "-m", "rc4-md5"]
CMD ["-k", "LoveWYN1008"]
EXPOSE 8000
