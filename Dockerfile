FROM ubuntu:18.04 as refal_host

RUN apt-get update
RUN apt-get install -y git dos2unix curl unzip sed g++
RUN apt install -y curl

# fetch refal-5-lambda
WORKDIR /usr/src
RUN git clone https://github.com/bmstu-iu9/simple-refal-distrib.git

# install refal-5-lambda
WORKDIR /usr/src/simple-refal-distrib
RUN ./bootstrap.sh

RUN apt-get update
RUN apt-get upgrade -y

ENV PATH="/usr/src/simple-refal-distrib/bin:${PATH}"
ENV RL_MODULE_PATH="/usr/src/simple-refal-distrib/lib:$RL_MODULE_PATH"

RUN curl https://storage.googleapis.com/golang/go1.16.2.linux-amd64.tar.gz -o go.tar.gz && \
    tar -zxf go.tar.gz && \
    rm -rf go.tar.gz && \
    mv go /go
ENV GOPATH /go
ENV PATH $PATH:/go/bin:$GOPATH/bin
# If you enable this, then gcc is needed to debug your app
ENV CGO_ENABLED 0
 
RUN mkdir -p /app
 
WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
RUN go build

RUN rlc dummy_terminator.ref
RUN rlc test_check.ref
RUN rlc test_gen.ref
RUN rlc total_results.ref

RUN dos2unix run_test.sh
RUN chmod +x run_test.sh
ENTRYPOINT ["./run_test.sh"]

