FROM ubuntu:24.04
RUN mkdir /app
WORKDIR /app
COPY  ./cmd .
ENTRYPOINT [ "/app/echo","defaultvalue" ]