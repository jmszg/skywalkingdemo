ARG BASEIMAGE=registry.cn-hangzhou.aliyuncs.com/zgjhub/apline:3.17.2
FROM $BASEIMAGE

COPY build/amd64/skywalkingdemo-server1 /
COPY conf/app.yaml /conf/app.yaml

CMD ["/skywalkingdemo-server1"]
