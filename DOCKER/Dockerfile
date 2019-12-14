# docker build -t noah-uploader -f DOCKER/Dockerfile .
# docker run -d -p
FROM golang:1.12-buster as builder

ENV APP_PATH /home/noah-gate

COPY . ${APP_PATH}

WORKDIR ${APP_PATH}

RUN make create_vendor && \
    make build && \
    cp ./builds/gate /usr/local/bin/gate

EXPOSE 15000
CMD ["gate"]
STOPSIGNAL SIGTERM
