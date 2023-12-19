#!/bin/bash

OUT_KUBE_GET_SECRET=$(kubectl get secret hf-deploy-secret -n dev)
OUT_KUBE_GET_REG_SECRET=$(kubectl get secret hfregcred -n dev)
OUT_KUBE_GET_NS=$(kubectl get ns | grep dev)

KUBE_REG_SERVER=$1
KUBE_REG_USERNAME=$2
KUBE_REG_PASSWORD=$3
KUBE_REG_EMAIL=$4

if [ ! -z "$OUT_KUBE_GET_NS" -a "$OUT_KUBE_GET_NS" != " " ]; then
    echo "namespace already exists, skipping..."
else
    kubectl create namespace dev
fi

if [ ! -z "$OUT_KUBE_GET_SECRET" -a "$OUT_KUBE_GET_SECRET" != " " ]; then
    echo "secret already exists, skipping..."
else 
    kubectl create secret generic hf-deploy-secret --from-env-file=.env --namespace=dev
fi

if [ ! -z "$OUT_KUBE_GET_REG_SECRET" -a "$OUT_KUBE_GET_REG_SECRET" != " " ]; then
    echo "registry secret already exists, skipping..."
else
    kubectl create secret docker-registry hfregcred --docker-server="$KUBE_REG_SERVER" --docker-username="$KUBE_REG_USERNAME" --docker-password="$KUBE_REG_PASSWORD" --docker-email="$KUBE_REG_EMAIL" --namespace=dev
fi

sed -i s:{{REPOSITORY_API_URL}}:$REPOSITORY_API_URL:g ./infrastructure/kubernetes/deployment/app.yaml;
sed -i s:{{IMAGE_TAG}}:$IMAGE_TAG:g ./infrastructure/kubernetes/deployment/app.yaml;

sed -i s:{{REPOSITORY_POSTGRES_URL}}:$REPOSITORY_POSTGRES_URL:g ./infrastructure/kubernetes/deployment/postgres.yaml;
sed -i s:{{IMAGE_TAG}}:$IMAGE_TAG:g ./infrastructure/kubernetes/deployment/postgres.yaml;

sed -i s:{{REPOSITORY_SWAGGER_URL}}:$REPOSITORY_SWAGGER_URL:g ./infrastructure/kubernetes/deployment/swagger.yaml;
sed -i s:{{IMAGE_TAG}}:$IMAGE_TAG:g ./infrastructure/kubernetes/deployment/swagger.yaml;

sed -i s:{{PWD}}:$PWD:g ./infrastructure/kubernetes/volume/postgres.yaml;

echo "teste $PWD"
