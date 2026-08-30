FROM alpine
WORKDIR /app

ADD shared shared
ADD build build

ENV GOTOOLCHAIN=local

ENTRYPOINT build/trip-service