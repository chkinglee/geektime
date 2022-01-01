FROM ubuntu
LABEL author="lilinzhen"

EXPOSE 8077
ENV VERSION="1.0.0"

COPY bin/amd64/httpserver /root/httpserver

ENTRYPOINT ["/root/httpserver"]

