APP_NAME=$(basename $PWD)
export APP_NAME=$(echo $APP_NAME|sed -e 's/_/-/g'|cut -c1-7)
echo $APP_NAME
function buildDocker() {
  echo "Building Dockerfile-based application..."
  # build image
  docker build \
    -t "letanthang/$APP_NAME" .
}

function pushDocker() {
  # docker login -u letanthang --password $DOCKER_REGISTRY_KEY
  docker push letanthang/$APP_NAME
}

function genK8sResource() {
  cd provision/
  sed 's/APP_NAME/'"$APP_NAME"'/g' k8s/deployment.yaml > zdeployment.yaml
  sed 's/APP_NAME/'"$APP_NAME"'/g' k8s/service.yaml > zservice.yaml
  sed 's/APP_NAME/'"$APP_NAME"'/g' k8s/ingress.yaml > zingress.yaml
  sed 's/APP_NAME/'"$APP_NAME"'/g' k8s/* > zall.yaml
}

function deployLocal() {
  echo "Run docker with name app_$APP_NAME"
  docker rm app_$APP_NAME --force
  docker run -d --name app_$APP_NAME -p 9090:9090 --network gitnet letanthang/$APP_NAME
}

function deployK8s() {
  cd provision/
  cat zall.yaml
  kubectl apply -f zall.yaml
}

buildDocker
pushDocker
genK8sResource
# deployLocal
deployK8s