FROM alpine:3.19.1 as base

ARG CLI_NAME="methodazure"
ARG TARGETARCH

RUN apk update && apk add bash jq

# Setup Method Directory Structure
RUN \
  mkdir -p /opt/method/${CLI_NAME}/ && \
  mkdir -p /opt/method/${CLI_NAME}/var/data && \
  mkdir -p /opt/method/${CLI_NAME}/var/data/tmp && \
  mkdir -p /opt/method/${CLI_NAME}/var/conf && \
  mkdir -p /opt/method/${CLI_NAME}/var/log && \
  mkdir -p /opt/method/${CLI_NAME}/service/bin && \
  mkdir -p /mnt/output

COPY scripts/* /opt/method/${CLI_NAME}/service/bin/

FROM base as amd64
COPY build/linux-amd64/${CLI_NAME} /opt/method/${CLI_NAME}/service/bin/${CLI_NAME}

FROM base as arm64
COPY build/linux-arm64/${CLI_NAME} /opt/method/${CLI_NAME}/service/bin/${CLI_NAME}

FROM ${TARGETARCH} as final
WORKDIR /opt/method/${CLI_NAME}/
ENTRYPOINT ["./service/bin/init.sh"]
