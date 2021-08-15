FROM golang:1.16 AS builder

# Set GOPATH to build Go app
ENV GOPATH=/go

# Set apps source directory
ENV SRC_DIR=${GOPATH}/src/github.com/Joker666/spotangels-demo/

# Define current working directory
WORKDIR ${SRC_DIR}

# Copy apps scource code to the image
COPY . ${SRC_DIR}

# Build App
RUN ./build.sh

ENTRYPOINT [ "spotangels-demo" ]
