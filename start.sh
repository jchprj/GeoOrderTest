curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
cd docker
# docker image remove geoordertest:latest
docker build -t geoordertest:latest  -f Dockerfile .
docker run -t -d -p 8080:8080 -p 8081:8081 --privileged --rm --name geoordertest_latest geoordertest:latest