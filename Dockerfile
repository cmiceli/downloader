FROM ghcr.io/mikenye/docker-youtube-dl:latest
USER root

# Set the working directory
WORKDIR /etc/downloader

COPY ./downloader /usr/bin/downloader
RUN mkdir -p /etc/downloader

# Your additional configuration
ENTRYPOINT ["/usr/bin/downloader"]

