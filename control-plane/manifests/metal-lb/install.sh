#!/bin/sh

helm repo add metallb https://metallb.github.io/metallb
helm repo update metallb

helm install metallb metallb/metallb --namespace metallb-system --create-namespace