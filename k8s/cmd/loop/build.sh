echo "Building loop version ${1}"
docker build --build-arg VER_TAG=$1 -t vallerion/loop:$1 .
docker build --platform=linux/amd64 --build-arg VER_TAG=$1 -t vallerion/loop:$1-amd64 .
docker push vallerion/loop:$1
docker push vallerion/loop:$1-amd64