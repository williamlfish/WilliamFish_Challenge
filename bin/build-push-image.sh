#637423595895.dkr.ecr.us-east-2.amazonaws.com/ccv
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 637423595895.dkr.ecr.us-east-2.amazonaws.com


TAG=$(git rev-parse --short HEAD)