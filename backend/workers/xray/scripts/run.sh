docker build -t jellyfish-xray:v1 . && docker run -p 1111:1111 --log-opt max-size=10m --log-opt max-file=3 --restart=always -d --name jellyfish-xray jellyfish-xray:v1
