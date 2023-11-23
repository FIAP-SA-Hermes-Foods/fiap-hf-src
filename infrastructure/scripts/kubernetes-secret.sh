#!/bin/bash

OUT_KUBE_GET_SECRET=$(kubectl get secret hf-deploy-secret)

if [ ! -z "$OUT_KUBE_GET_SECRET" -a "$OUT_KUBE_GET_SECRET" != " " ]; then
    echo "secret already exists, initializing..."
else 
    kubectl create secret generic hf-deploy-secret --from-env-file=.env
fi

