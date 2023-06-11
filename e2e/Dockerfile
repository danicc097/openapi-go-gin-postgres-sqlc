#
# See https://playwright.dev/python/docs/docker
# Use a Docker image version that matches Playwright version.
# use supported regular ubuntu
#
FROM mcr.microsoft.com/playwright:v1.35.0-jammy

WORKDIR /e2e

RUN npm i -g pnpm@8.3.1

ENV PATH /e2e/node_modules/.bin:$PATH

COPY pnpm-lock.yaml /e2e

RUN --mount=type=cache,target=/root/.local/share/pnpm/store/v3 \
  pnpm fetch

COPY . /e2e

RUN --mount=type=cache,target=/root/.local/share/pnpm/store/v3 \
  pnpm install -r --offline
