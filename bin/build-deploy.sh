set -e
export CLUSTER_NAME="comcastCCchallenge"

aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 637423595895.dkr.ecr.us-east-2.amazonaws.com

export TAG=$(git rev-parse --short HEAD)
export IMAGE="637423595895.dkr.ecr.us-east-2.amazonaws.com/ccv"
docker build -t "$IMAGE:$TAG" .
docker push "$IMAGE:$TAG"



export KUBECONFIG="$HOME/.kube/eksctl/clusters/${CLUSTER_NAME}"



export NS='comcast-challenge'

if ! kubectl get ns $NS
then
  echo 'Creating the namespace'
  kubectl create ns $NS
fi



envsubst < infra/deployment.yaml | kubectl apply -n $NS -f -

#IF WE HAD SECRETS TO CREATE...
#kubectl create secret generic  comcast-challenge \
#  --save-config \
#  --from-env-file=.prod.env \
#  --dry-run=client -o yaml | kubectl apply -n $NS -f -