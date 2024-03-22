# SPDX-License-Identifier: Apache-2.0

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
ARG HUGO_VERSION=0.115.4

##########################################################################
##    docker build --no-cache --target binary -t vela-hugo:binary .     ##
##########################################################################

FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b as binary

ARG HUGO_VERSION

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz  /tmp/hugo.tar.gz
 
RUN tar -xzf /tmp/hugo.tar.gz -C /bin

RUN chmod 0700 /bin/hugo

########################################################
##    docker build --no-cache -t vela-hugo:local .    ##
########################################################

FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

ARG HUGO_VERSION

ENV PLUGIN_HUGO_VERSION=${HUGO_VERSION}

RUN apk add --update --no-cache ca-certificates git libc6-compat libstdc++ nodejs npm

COPY --from=binary /bin/hugo /bin/hugo

COPY release/vela-hugo /bin/vela-hugo

ENTRYPOINT [ "/bin/vela-hugo" ]
