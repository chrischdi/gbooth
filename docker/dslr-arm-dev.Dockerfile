FROM k8s.gcr.io/debian-base-arm:1.0.0

RUN apt-get update && \
  apt-get install -y \
    gphoto2 \
    libc6

RUN apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y \
    build-essential \
    libc6-dev \
    libgphoto2-dev \
    pkg-config 
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y \
    libc6-dev
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y \
    libgtk-3-dev 
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y \
    libvips-tools