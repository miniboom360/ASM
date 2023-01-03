FROM alpine:3.13.2
# gobuster_3.4.0_Linux_i386 release
COPY gobuster gobuster
COPY gobuster_service gobuster_service
COPY wordlist.txt wordlist.txt
COPY config.json config.json
RUN ["chmod", "u+x", "gobuster_service"]
RUN ["chmod", "u+x", "gobuster"]
CMD ["/gobuster_service","worker"]