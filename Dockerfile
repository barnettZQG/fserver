FROM alpine:3.4
COPY fserver /fserver
EXPOSE 8080
LABEL traefik.tags=function-fserver
LABEL traefik.frontend.rule=Host:fs.faas.pro
LABEL traefik.port=8080
ENTRYPOINT ["/fserver"]

