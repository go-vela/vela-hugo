# Copyright (c) 2021 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
ARG HUGO_VERSION=0.76.5

##########################################################################
##    docker build --no-cache --target binary -t vela-hugo:binary .     ##
##########################################################################

FROM alpine as binary

ARG HUGO_VERSION

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz  /tmp/hugo.tar.gz
 
RUN tar -xzf /tmp/hugo.tar.gz -C /bin

RUN chmod 0700 /bin/hugo

########################################################
##    docker build --no-cache -t vela-hugo:local .    ##
########################################################

FROM alpine

ARG HUGO_VERSION

ENV PLUGIN_HUGO_VERSION=${HUGO_VERSION}

RUN apk add --update --no-cache ca-certificates git libc6-compat libstdc++ nodejs

COPY --from=binary /bin/hugo /bin/hugo

COPY release/vela-hugo /bin/vela-hugo

ENTRYPOINT [ "/bin/vela-hugo" ]
