# gke-google-managed-ssl

Example to use Google Managed SSL to Ingress

## Prepared

### Install k8s tools

```console
brew install skaffoled # For local server
brew install kustomize
```

### Set up ingress-nginx to Docker for Mac

```console
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.29.0/deploy/static/mandatory.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.29.0/deploy/static/provider/cloud-generic.yaml
```

## Start Local Server

```console
./scripts/serve.sh
```

## Deploy to GKE

### Create Resources

```console
./scripts/create_resouces.sh
```

### Set DNS

#### Check Static IP Address

```console
gcloud compute addresses describe --global example-cluster-ig --format 'value("address")'
```

### Build Image

```console
./scripts/submit_image.sh
```

### Deploy

```console
./scripts/deploy.sh
```
