FROM golang:1.10beta1

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apt-get update \
  && apt-get install -y build-essential git curl wget \
                        zlib1g-dev libxml2-dev libxslt1-dev \
                        openssl less libssl-dev libreadline-dev vim

COPY ./ /go/src/
COPY ./ /go/src/app/

WORKDIR /go/src/app

RUN go get github.com/pilu/fresh

RUN curl https://glide.sh/get | sh
# RUN glide install
