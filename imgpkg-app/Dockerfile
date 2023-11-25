FROM ubuntu:latest

RUN apt-get -y update

RUN apt-get -y install wget

RUN wget https://github.com/carvel-dev/imgpkg/releases/download/v0.39.0/imgpkg-linux-amd64

RUN mv imgpkg-linux-amd64 /usr/local/bin/imgpkg
RUN chmod +x /usr/local/bin/imgpkg

CMD ["/bin/bash", "-c", "sleep infinity"]