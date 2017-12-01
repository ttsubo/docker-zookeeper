FROM ubuntu:16.04

MAINTAINER Toshiki Tsuboi <t.tsubo2000@gmail.com>

ENV GOPATH  /golang
ADD golang golang
ADD log log
ADD root root

RUN apt-get update \
 && apt-get install -y golang \
 && apt-get install -y git \
 && apt-get install -y vim \
 && apt-get install -y iputils-ping \
 && apt-get install -y net-tools \
 && apt-get install -y screen

# Setup environment varialbles for go
ENV PATH  /usr/local/go/bin:/golang/bin:$PATH

RUN go get github.com/samuel/go-zookeeper/zk

