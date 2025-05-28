FROM debian:stable-slim

ARG GO_VERSION=1.22.0

RUN apt-get update && apt-get install -y \
    curl \
    git

RUN curl -fsSL https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get install -y nodejs
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /workspaces

CMD ["bash"]