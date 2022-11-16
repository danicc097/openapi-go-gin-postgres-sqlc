#
# See https://playwright.dev/python/docs/docker
# Use a Docker image version that matches Playwright version.
# use supported regular ubuntu
#
FROM ubuntu:focal-20221019

WORKDIR /src
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

COPY requirements.txt /src/requirements.txt
COPY requirements-dev.txt /src/requirements-dev.txt

RUN --mount=type=cache,target=/root/.cache/pip \
  apt-get update \
  && apt-get install -y software-properties-common \
  && add-apt-repository ppa:deadsnakes/ppa \
  && apt-get update \
  && apt-get install -y python3.9 wget \
  && BUILD_DEPS='libpq-dev libffi-dev musl-dev gcc' \
  && DEBIAN_FRONTEND=noninteractive \
  && apt-get install -y netcat python3.9-distutils ${BUILD_DEPS} \
  && wget https://bootstrap.pypa.io/get-pip.py \
  && python3.9 get-pip.py \
  && pip3 install pip-tools \
  && pip-sync requirements.txt requirements-dev.txt \
  && playwright install \
  && playwright install-deps \
  && apt-get autoremove -y ${BUILD_DEPS} \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists /tmp/* /var/tmp/*
