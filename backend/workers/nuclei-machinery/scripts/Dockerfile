FROM alpine:3.13.2
# gobuster_3.4.0_Linux_i386 release
COPY nuclei /usr/local/bin/nuclei
COPY config.json config.json
COPY nuclei_service nuclei_service
RUN ["chmod", "u+x", "nuclei_service"]
CMD ["/nuclei_service","worker"]