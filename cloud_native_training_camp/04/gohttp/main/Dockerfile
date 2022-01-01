FROM ubuntu
LABEL author="lilinzhen"
LABEL homework="module_3"

EXPOSE 80
ENV VERSION="CENTOS_lilinzhen"
# go build -o gohttp main.go
ADD gohttp /root/gohttp

ENTRYPOINT ["/root/gohttp"]
#ENTRYPOINT ["top"]
#CMD ["-b"]
