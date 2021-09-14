FROM alpine:latest

RUN apk add curl

RUN mkdir /root/.kube

COPY .kube/config /root/.kube/config

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

RUN chmod +x kubectl

RUN mv kubectl /usr/local/bin/kubectl
