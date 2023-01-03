docker build -t machinery-services-gobuster:v1 . && docker run --log-opt max-size=10m --log-opt max-file=3 --restart=always -d --name machinery-services-gobuster machinery-services-gobuster:v1
