curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
cd docker
# docker image remove geoordertest:latest
docker build -t geoordertest:latest  -f Dockerfile .
docker run -it --privileged --rm --name geoordertest_latest geoordertest:latest