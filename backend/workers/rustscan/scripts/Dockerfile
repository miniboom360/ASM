FROM rustscan/rustscan:1.10.0
USER root

COPY rustscan_service rustscan_service
COPY config.json config.json
RUN ["chmod", "u+x", "rustscan_service"]
CMD ["/rustscan_service","worker"]