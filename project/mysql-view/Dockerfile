#FROM golang:alpine
FROM ubuntu
ADD  ./web /web
ADD ./layout.html /
RUN chmod a+x /web
RUN echo " ./web ">> /start.sh
RUN chmod a+x /start.sh


