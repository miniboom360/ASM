#docker run -it robd003/python3.10:latest /bin/bash启动直接进入docker内部
FROM robd003/python3.10:latest

COPY oneforall oneforall
COPY python /python
COPY config.json config.json
COPY oneforall_run.sh oneforall_run.sh
RUN ["chmod", "u+x", "oneforall_run.sh"]
CMD ["/oneforall_run.sh"]