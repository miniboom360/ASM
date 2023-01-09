FROM ubuntu:14.04
COPY xray xray
COPY xray_run.sh xray_run.sh
RUN ["chmod", "u+x", "xray"]
RUN ["sudo", "apt-get", "install", "-y", "qemu", "gcc-aarch64-linux-gnu"]
RUN ["chmod", "u+x", "xray_run.sh"]
CMD ["/bin/bash", "-c", "/xray_run.sh"]