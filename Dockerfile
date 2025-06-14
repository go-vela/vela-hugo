# SPDX-License-Identifier: Apache-2.0

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
# renovate: datasource=github-tags depName=gohugoio/hugo extractVersion=^v(?<version>.*)$
ARG HUGO_VERSION=0.147.7

##########################################################################
##    docker build --no-cache --target binary -t vela-hugo:binary .     ##
##########################################################################

FROM alpine:3.22.0@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715 AS binary

ARG HUGO_VERSION

ENV HUGO_RELEASE_URL="https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}"
ENV HUGO_FILENAME="hugo_${HUGO_VERSION}_linux-amd64.tar.gz"
ENV HUGO_CHECKSUM_FILENAME="hugo_${HUGO_VERSION}_checksums.txt"

RUN wget -q "${HUGO_RELEASE_URL}/${HUGO_FILENAME}" -O "${HUGO_FILENAME}" && \
  wget -q "${HUGO_RELEASE_URL}/${HUGO_CHECKSUM_FILENAME}" -O "${HUGO_CHECKSUM_FILENAME}" && \
  grep "${HUGO_FILENAME}" "${HUGO_CHECKSUM_FILENAME}" | sha256sum -c && \
  tar -xf "${HUGO_FILENAME}" && \
  mv hugo /bin/hugo && \
  chmod 0700 /bin/hugo

########################################################
##    docker build --no-cache -t vela-hugo:local .    ##
########################################################

FROM alpine:3.22.0@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715

ARG HUGO_VERSION

ENV PLUGIN_HUGO_VERSION=${HUGO_VERSION}

RUN apk add --update --no-cache ca-certificates git libc6-compat libstdc++ nodejs npm

COPY --from=binary /bin/hugo /bin/hugo

COPY release/vela-hugo /bin/vela-hugo

ENTRYPOINT [ "/bin/vela-hugo" ]
