FROM mcr.microsoft.com/playwright:v1.35.0-jammy

WORKDIR /e2e

RUN npm i -g pnpm@8.3.1

ENV PATH /e2e/node_modules/.bin:$PATH

COPY package.json /e2e/
COPY pnpm-lock.yaml /e2e/
RUN --mount=type=cache,target=/root/.local/share/pnpm/store/v3 \
  pnpm install -r

RUN ls -lhA /e2e/node_modules/.bin
RUN ls -lhA /e2e
