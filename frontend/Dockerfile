# ベースステージ
FROM node:18.12.1-alpine AS base
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

# 環境変数の設定
ARG BACKEND_ENDPOINT
ENV REACT_APP_BACKEND_ENDPOINT=$BACKEND_ENDPOINT

# 開発環境ステージ
FROM base AS dev
COPY . .
CMD ["yarn", "start"]

# ビルドステージ
FROM base AS build
COPY . .
RUN yarn build

# 本番環境用
FROM nginx:alpine AS prod
COPY --from=build /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80