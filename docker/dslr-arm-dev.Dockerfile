FROM k8s.gcr.io/debian-base-arm:1.0.0

RUN apt-get update && \
  apt-get install -y \
    gphoto2 \
    libc6

RUN apt-get update && \
  apt-get install -y \
    build-essential \
    libc6-dev \
    libgphoto2-dev \
    pkg-config