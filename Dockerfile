# SPDX-License-Identifier: Apache-2.0

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
ARG HUGO_VERSION=0.115.4

##########################################################################
##    docker build --no-cache --target binary -t vela-hugo:binary .     ##
##########################################################################

FROM alpine:3.18.2@sha256:82d1e9d7ed48a7523bdebc18cf6290bdb97b82302a8a9c27d4fe885949ea94d1 as binary

ARG HUGO_VERSION

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz  /tmp/hugo.tar.gz
 
RUN tar -xzf /tmp/hugo.tar.gz -C /bin

RUN chmod 0700 /bin/hugo

########################################################
##    docker build --no-cache -t vela-hugo:local .    ##
########################################################

FROM alpine:3.18.2@sha256:82d1e9d7ed48a7523bdebc18cf6290bdb97b82302a8a9c27d4fe885949ea94d1

ARG HUGO_VERSION

ENV PLUGIN_HUGO_VERSION=${HUGO_VERSION}

RUN apk add --update --no-cache ca-certificates git libc6-compat libstdc++ nodejs npm

COPY --from=binary /bin/hugo /bin/hugo

COPY release/vela-hugo /bin/vela-hugo

ENTRYPOINT [ "/bin/vela-hugo" ]
