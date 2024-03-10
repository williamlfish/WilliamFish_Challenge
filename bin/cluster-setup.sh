set -e
export CLUSTER_NAME="comcastCCchallenge"
export CLUSTER_REGION="us-east-2"

eksctl create cluster --name=${CLUSTER_NAME} --auto-kubeconfig --region=${CLUSTER_REGION}

export KUBECONFIG="$HOME/.kube/eksctl/clusters/${CLUSTER_NAME}"

export CLUSTER_ACCOUNT=$(aws sts get-caller-identity --query Account --o text)

export CLUSTER_VPC=$(aws eks describe-cluster --region=${CLUSTER_REGION} --name ${CLUSTER_NAME}  --query "cluster.resourcesVpcConfig.vpcId" --output text)

export AWS_ROUTE53_DOMAIN="nboop.com"

aws route53 create-hosted-zone --name "${AWS_ROUTE53_DOMAIN}." --vpc VPCRegion=${CLUSTER_REGION},VPCId=${CLUSTER_VPC} --caller-reference "nboop" --output text

export HOSTED_ZONE_ID=$(aws route53 list-hosted-zones-by-name --dns-name "${AWS_ROUTE53_DOMAIN}." --query 'HostedZones[0].Id' --o text | awk -F "/" {'print $NF'})

aws iam create-policy --policy-name "ExternalDNSSA" --policy-document file://external-dns-policy.json --output text

eksctl utils associate-iam-oidc-provider --region=${CLUSTER_REGION} --cluster=${CLUSTER_NAME} --approve

eksctl create iamserviceaccount --name externalnds --namespace kube-system --cluster ${CLUSTER_NAME} --attach-policy-arn arn:aws:iam::637423595895:policy/ExternalDNSSA --approve --override-existing-serviceaccounts --region ${CLUSTER_REGION}

aws eks update-kubeconfig --name ${CLUSTER_NAME} --region ${CLUSTER_REGION} --output text

helm upgrade --wait --timeout 900s --install externaldns-release \
  --set provider=aws \
  --set aws.region=${CLUSTER_REGION} \
  --set txtOwnerId=${HOSTED_ZONE_ID} \
  --set domainFilters\[0\]="${AWS_ROUTE53_DOMAIN}" \
  --set serviceAccount.name=externalnds \
  --set serviceAccount.create=false \
  --set policy=sync \
  oci://registry-1.docker.io/bitnamicharts/external-dns --namespace externaldns


##instal ingress clr
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/cloud/deploy.yaml

##cert manager
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.4/cert-manager.yaml


