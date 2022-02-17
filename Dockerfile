FROM golang:1.17

RUN wget https://github.com/google/protobuf/releases/download/v2.6.1/protobuf-2.6.1.tar.gz
RUN tar xzf protobuf-2.6.1.tar.gz
RUN cd protobuf-2.6.1
RUN apt-get update
RUN apt-get install -y sudo build-essential
RUN sudo ./configure
RUN sudo make
RUN sudo make install
RUN sudo ldconfig
RUN cd /app/go-grpc

CMD ["go", "run", "main.go"]