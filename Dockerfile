#
# For production:
#   docker build --target production -t CIPHERC2 .
#   docker run -it --rm -v $HOME/.CIPHERC2:/home/CIPHERC2/.CIPHERC2 CIPHERC2 
#
# For unit testing:
#   docker build --target test .
#

# STAGE: base
## Compiles CIPHERC2 for use
FROM golang:1.23.5 AS base

### Base packages
RUN apt-get update --fix-missing && apt-get -y install \
    git build-essential zlib1g zlib1g-dev wget zip unzip

### Add CIPHERC2 user
RUN groupadd -g 999 CIPHERC2 && useradd -r -u 999 -g CIPHERC2 CIPHERC2
RUN mkdir -p /home/CIPHERC2/ && chown -R CIPHERC2:CIPHERC2 /home/CIPHERC2

### Build CIPHERC2:
WORKDIR /go/src/github.com/bishopfox/CIPHERC2
ADD . /go/src/github.com/bishopfox/CIPHERC2/
RUN make clean-all 
RUN make 
RUN cp -vv CIPHERC2-server /opt/CIPHERC2-server 

# STAGE: test
## Run unit tests against the compiled instance
## Use `--target test` in the docker build command to run this stage
FROM base AS test

RUN apt-get update --fix-missing \
    && apt-get -y upgrade \
    && apt-get -y install \
    curl

RUN /opt/CIPHERC2-server unpack --force 

### Run unit tests
RUN /go/src/github.com/bishopfox/CIPHERC2/go-tests.sh

# STAGE: production
## Final dockerized form of CIPHERC2
FROM debian:bookworm-slim AS production

### Install production packages
RUN apt-get update --fix-missing \
    && apt-get -y upgrade \
    && apt-get -y install \
    libxml2 libxml2-dev libxslt-dev locate gnupg \
    libreadline6-dev libcurl4-openssl-dev git-core \
    libssl-dev libyaml-dev openssl autoconf libtool \
    ncurses-dev bison curl xsel postgresql \
    postgresql-contrib postgresql-client libpq-dev \
    curl libapr1 libaprutil1 libsvn1 \
    libpcap-dev libsqlite3-dev libgmp3-dev \
    nasm

### Install MSF for stager generation
RUN curl https://raw.githubusercontent.com/rapid7/metasploit-omnibus/master/config/templates/metasploit-framework-wrappers/msfupdate.erb > msfinstall \
    && chmod 755 msfinstall \
    && ./msfinstall \
    && mkdir -p ~/.msf4/ \
    && touch ~/.msf4/initial_setup_complete 

### Cleanup unneeded packages
RUN apt-get remove -y curl gnupg \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

### Add CIPHERC2 user
RUN groupadd -g 999 CIPHERC2 \
    && useradd -r -u 999 -g CIPHERC2 CIPHERC2 \
    && mkdir -p /home/CIPHERC2/ \
    && chown -R CIPHERC2:CIPHERC2 /home/CIPHERC2 \
    && su -l CIPHERC2 -c 'mkdir -p ~/.msf4/ && touch ~/.msf4/initial_setup_complete'

### Copy compiled binary
COPY --from=base /opt/CIPHERC2-server  /opt/CIPHERC2-server

### Unpack CIPHERC2:
USER CIPHERC2
RUN /opt/CIPHERC2-server unpack --force 

WORKDIR /home/CIPHERC2/
VOLUME [ "/home/CIPHERC2/.CIPHERC2" ]
ENTRYPOINT [ "/opt/CIPHERC2-server" ]


# STAGE: production-slim (about 1Gb smaller)
FROM debian:bookworm-slim as production-slim

### Install production packages
RUN apt-get update --fix-missing \
    && apt-get -y upgrade

### Cleanup unneeded packages
RUN apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

### Add CIPHERC2 user
RUN groupadd -g 999 CIPHERC2 \
    && useradd -r -u 999 -g CIPHERC2 CIPHERC2 \
    && mkdir -p /home/CIPHERC2/ \
    && chown -R CIPHERC2:CIPHERC2 /home/CIPHERC2

### Copy compiled binary
COPY --from=base /opt/CIPHERC2-server  /opt/CIPHERC2-server

### Unpack CIPHERC2:
USER CIPHERC2
RUN /opt/CIPHERC2-server unpack --force 

WORKDIR /home/CIPHERC2/
VOLUME [ "/home/CIPHERC2/.CIPHERC2" ]
ENTRYPOINT [ "/opt/CIPHERC2-server" ]
