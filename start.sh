docker -v
if [ ! $? -eq 0 ]; then
  echo install docker
  curl -fsSL https://get.docker.com -o get-docker.sh
  sh get-docker.sh
fi
cd docker
# docker stop geoordertest_latest
# docker image remove geoordertest:latest
echo build image geoordertest:latest
docker build -t geoordertest:latest  -f Dockerfile .
echo run image geoordertest:latest
docker run -t -d -p 8080:8080 -p 8081:8081 --privileged --rm --name geoordertest_latest geoordertest:latest