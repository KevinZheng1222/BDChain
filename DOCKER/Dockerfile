FROM alpine:3.7

# This is the release of bdc to pull in.
ENV TM_VERSION 0.17.1
ENV TM_SHA256SUM d57008c63d2d9176861137e38ed203da486febf20ae7d388fb810a75afff8f24

# bdc will be looking for genesis file in /bdc (unless you change
# `genesis_file` in config.toml). You can put your config.toml and private
# validator file into /bdc.
#
# The /bdc/data dir is used by bdc to store state.
ENV DATA_ROOT /bdc
ENV TMHOME $DATA_ROOT

# Set user right away for determinism
RUN addgroup tmuser && \
    adduser -S -G tmuser tmuser

# Create directory for persistence and give our user ownership
RUN mkdir -p $DATA_ROOT && \
    chown -R tmuser:tmuser $DATA_ROOT

# jq and curl used for extracting `pub_key` from private validator while
# deploying bdc with Kubernetes. It is nice to have bash so the users
# could execute bash commands.
RUN apk add --no-cache bash curl jq

RUN apk add --no-cache openssl && \
    wget https://github.com/bdc/bdc/releases/download/v${TM_VERSION}/bdc_${TM_VERSION}_linux_amd64.zip && \
    echo "${TM_SHA256SUM}  bdc_${TM_VERSION}_linux_amd64.zip" | sha256sum -c && \
    unzip -d /bin bdc_${TM_VERSION}_linux_amd64.zip && \
    apk del openssl && \
    rm -f bdc_${TM_VERSION}_linux_amd64.zip

# Expose the data directory as a volume since there's mutable state in there
VOLUME $DATA_ROOT

# p2p port
EXPOSE 46656
# rpc port
EXPOSE 46657

ENTRYPOINT ["bdc"]

CMD ["node", "--moniker=`hostname`"]
