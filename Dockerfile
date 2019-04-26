FROM debian:stretch-slim

WORKDIR /app
COPY build/front /app/
COPY templates /app/templates/

EXPOSE 3000
ENTRYPOINT [ "/app/front" ]