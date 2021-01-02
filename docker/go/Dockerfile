FROM golang:1.13.1-alpine
# goのイメージをベースにする
ENV GO111MODULE=on
# goのライブラリの環境変数
RUN apk update && \
    apk add emacs curl git && \
    go get github.com/pilu/fresh
# aptとfreshのインストール