echo "Building hellok8s version ${1}"
docker build --build-arg VER_TAG=$1 -t vallerion/hellok8s:$1 .
docker build --platform=linux/amd64 --build-arg VER_TAG=$1 -t vallerion/hellok8s:$1-amd64 .
docker push vallerion/hellok8s:$1
docker push vallerion/hellok8s:$1-amd64