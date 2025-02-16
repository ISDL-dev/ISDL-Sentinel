#!/bin/sh

# 環境変数を埋め込んで nginx.conf を生成
envsubst '$ENV_TYPE' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Nginx を起動
exec "$@"
