#!/bin/bash

OUT_KUBE_GET_SECRET=$(kubectl get secret hf-deploy-secret)
OUT_KUBE_GET_NS=$(kubectl get ns | grep dev)

if [ ! -z "$OUT_KUBE_GET_SECRET" -a "$OUT_KUBE_GET_SECRET" != " " ]; then
    echo "secret already exists, initializing..."
else 
    kubectl create secret generic hf-deploy-secret --from-env-file=.env
fi

if [ ! -z "$OUT_KUBE_GET_NS" -a "$OUT_KUBE_GET_NS" != " " ]; then
    echo "namespace already exists, initializing..."
else 
    kubectl create namespace dev
fi
