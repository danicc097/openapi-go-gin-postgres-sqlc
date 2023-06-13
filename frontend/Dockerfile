FROM node:18.5.0-alpine as nodebuild

WORKDIR /frontend

RUN npm i -g pnpm@8.3.1

ENV PATH /frontend/node_modules/.bin:$PATH
ENV JQ_VERSION=1.6

RUN wget --no-check-certificate https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-linux64 -O /usr/bin/jq
RUN chmod +x /usr/bin/jq
RUN apk add --no-cache bash git openssh gettext

COPY ./frontend/pnpm-lock.yaml /frontend/

# https://pnpm.io/cli/fetch (ensure buildkit used and cache target up to date: pnpm store path)
RUN --mount=type=cache,target=/root/.local/share/pnpm/store/v3 \
  pnpm fetch
COPY ./frontend/. /frontend/
# ensure config has all k-v as "<KEY>": "$<KEY>"
RUN jq \
  'to_entries | map_values({ (.key) : ("$" + .key) }) | reduce .[] as $item ({}; . + $item)' \
  ./config.template.json > ./config.tmp.json && mv ./config.tmp.json ./config.json
RUN cat config.json

# --offline since it's been fetched beforehand
RUN --mount=type=cache,target=/root/.local/share/pnpm/store/v3 \
  pnpm install -r --offline
RUN ls -h
#RUN pnpm run test:no-watch
# runtime env not available in vite.config

# copy symlink targets, not enough to just have build context containing them
# TODO: should this be necessary?
COPY ./*.json /

RUN pnpm run build

FROM nginx:1-alpine
ENV ENV_REPLACE_GLOB=/usr/share/nginx/html/assets/*.js
RUN apk add --no-cache bash moreutils
COPY --from=nodebuild /frontend/build /usr/share/nginx/html

COPY ./frontend/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./frontend/nginx/replace-envvars.sh /docker-entrypoint.d/
RUN chmod +x /docker-entrypoint.d/replace-envvars.sh

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
