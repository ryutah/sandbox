# appengine-over-iap

## Prepare

### Setup Cloud IAP to App Engine

### Set secrets for Secret Manager API

1. Enable Secret Manager API

    ```console
    gcloud services enable secretmanager.googleapis.com
    ```

1. Create Secret

    ```console
    echo ${IAP_CLIENT_ID} |
      gcloud beta secrets create appeingine-iap-client \
      --replication-policy automatic \
      --data-file -
    ```

1. Add IAM Policy to access Secret

```console
export PROJECT_ID=$(gcloud config get-value project)
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member "${PROJECT_ID}@appspot.gserviceaccount.com " \
  --role 'roles/secretmanager.secretAccessor'
```

### Create Cloud Tasks Queue and check location

1. Create Cloud Tasks Queue

    ```console
    gcloud tasks queues create iap-example
    ```

1. Get Location

    ```console
    gcloud tasks queues describe iap-example --format 'value("name")' | \
      sed -e 's/^.*locations\/\(.*\)\/queues.*$/\1/g'
    ```

### Set app.yaml Environment Variables

```yaml
runtime: go112
service: service-a
env_variables:
  QUEUE_LOCATION: [CLOUD_TASKS_QUEUE_LOCATION]
```

## Deploy and settings

### Deploy

```console
gcloud --quiet app deploy ./service-a ./service-b
```

### Add Cloud IAP Secured-Uers to service-b

## Access

```console
curl https://service-a-dot-[PROJECT_ID].appspot.com
```

### Access Over Cloud Tasks

```console
curl https://service-a-dot-[PROJECT_ID].appspot.com/tasks
```
