echo "Building version ${1}"
docker build --build-arg VER_TAG=$1 -t vallerion/k8s-cource-1:$1 .
docker push vallerion/k8s-cource-1:$1