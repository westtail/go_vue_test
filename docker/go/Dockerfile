# goのイメージをベースにする
FROM golang:1.13.1-alpine

# goのライブラリの環境変数
ENV GO111MODULE=on

# aptとfreshのインストール
RUN apk update && \
    apk add emacs curl git && \
    go get github.com/pilu/fresh

# mySQLのインストール
RUN  go get github.com/go-sql-driver/mysql