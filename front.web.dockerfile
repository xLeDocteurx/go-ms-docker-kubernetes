FROM mhart/alpine-node:12 AS build
# FROM node:10 AS build
# FROM node:13.8 AS build
# FROM node:12.0 AS build
LABEL stage=builder

RUN mkdir -p /app

COPY ./front/web /app/front/web

COPY ./utils /app/utils

COPY ./models /app/models

WORKDIR /app/utils

RUN npm install

WORKDIR /app/models

RUN npm install

WORKDIR /app/front/web

RUN npm install

# RUN ./localeClean.sh

# RUN npm run test

RUN npm run build --loglevel verbose

FROM alpine:latest AS app
LABEL stage=builder

RUN mkdir -p /app

COPY --from=build /app/front/web/dist /app

FROM nginx:latest

COPY --from=app /app /app

COPY ./nginx/container/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

# RUN certbot --nginx --agree-tos -m admin@timebom.be -d timebom.be -d www.timebom.be --rsa-key-size 4096

# LABEL maintainer="Jonas Alfredsson <jonas.alfredsson@protonmail.com>"

# ARG BUILDX_QEMU_ENV

# # Do a single run command to make the intermediary containers smaller.
# RUN set -ex && \
# # Install packages necessary during the build phase (for all architectures).
#     apt-get update && \
#     apt-get install -y --no-install-recommends \
#             build-essential \
#             cargo \
#             curl \
#             libffi6 \
#             libffi-dev \
#             libssl-dev \
#             openssl \
#             procps \
#             python3 \
#             python3-dev \
#     && \
# # Install the latest version of PIP, Setuptools and Wheel.
#     curl -L 'https://bootstrap.pypa.io/get-pip.py' | python3 && \
# # Handle an extremely specific issue when building the cryptography package for
# # 32-bit architectures within QEMU running on a 64-bit host (issue #30).
#     if [ "${BUILDX_QEMU_ENV}" = "true" ] && [ "$(getconf LONG_BIT)" = "32" ]; then \
#         pip3 install -U cryptography==3.3.2; \
#     fi && \
# # Install certbot.
#     pip3 install -U cffi certbot \
#     && \
# # Remove everything that is no longer necessary.
#     apt-get remove --purge -y \
#             build-essential \
#             cargo \
#             curl \
#             libffi-dev \
#             libssl-dev \
#             python3-dev \
#     && \
#     apt-get autoremove -y && \
#     apt-get clean && \
#     rm -rf /var/lib/apt/lists/* && \
#     rm -rf /root/.cache && \
#     rm -rf /root/.cargo && \
# # Create new directories and set correct permissions.
#     mkdir -p /var/www/letsencrypt && \
#     mkdir -p /etc/nginx/user_conf.d && \
#     chown www-data:www-data -R /var/www \
#     && \
# # Make sure there are no surprise config files inside the config folder.
#     rm -f /etc/nginx/conf.d/*

# # Copy in our "default" Nginx server configurations, which make sure that the
# # ACME challenge requests are correctly forwarded to certbot and then redirects
# # everything else to HTTPS.
# COPY ./letsencrypt/nginx/ /etc/nginx/conf.d/

# # Copy in all our scripts and make them executable.
# COPY ./letsencrypt/scripts/ /scripts
# RUN chmod +x -R /scripts && \
# # Make so that the parent's entrypoint script is properly triggered (issue #21).
#     sed -ri '/^if \[ "\$1" = "nginx" -o "\$1" = "nginx-debug" \]; then$/,${s//if echo "$1" | grep -q "nginx"; then/;b};$q1' /docker-entrypoint.sh

# # Create a volume to have persistent storage for the obtained certificates.
# VOLUME /etc/letsencrypt

# # The Nginx parent Docker image already expose port 80, so we only need to add
# # port 443 here.
# EXPOSE 443
# EXPOSE 80

# # Change the container's start command to launch our Nginx and certbot
# # management script.
# CMD [ "/scripts/start_nginx_certbot.sh" ]
