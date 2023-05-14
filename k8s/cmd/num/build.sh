echo "Building num version ${1}"
docker build --build-arg VER_TAG=$1 -t vallerion/num:$1 .
docker build --platform=linux/amd64 --build-arg VER_TAG=$1 -t vallerion/num:$1-amd64 .
docker push vallerion/num:$1
docker push vallerion/num:$1-amd64