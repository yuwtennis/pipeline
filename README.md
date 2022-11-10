# pipeline

![Action Status](https://github.com/yuwtennis/pipelines/actions/workflows/main.yml/badge.svg)

* [Motivation](#motivation)
* [Available pipelines](#available-pipelines)
  * [Downloading real-estate data](#downloading-real-estate-data)
* [Sample dashboard using kibana](#sample-dashboard-using-kibana)
  * [Real Estate](#real-estate)
* [Getting started](#getting-started)
* [Conventions](#conventions)
  * [Commit message](#commit-message)
* [Sample infrastructure]

## Motivation

A pipeline powered by apache beam for general use.  
I would like to introduce this as practical example.

## Available pipelines

### Downloading real estate data

This pipeline targets mansion price in Japan which is downloaded from [MLIT](https://www.land.mlit.go.jp/webland/servlet/MainServlet).  
Data will be indexed into elasticsearch.

```markdown
cd cmd/real_estates
go run main.go
```

### Indexing to elasticsearch using cross language platform

```markdown
# Spin up expansion service
export JOB_SERVER_PORT=20000
python -m apache_beam.runners.portability.local_job_service_main -p $JOB_SERVER_PORT

# In different shell, run pipeline
cd cmd/xlang
go run main.go \
  --staging_location xxx \
  --zone xxx \
  --subnetwork xxx
```


## Sample dashboard using kibana

### Real Estate

![](images/RealEstate1.png)
![](images/RealEstate2.png)

## Getting Started

1. Build image
```markdown
make build
```

2. Run app
```markdown
docker run --rm pipelines:latest
```

## Conventions

### Commit message

I follow [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/#summary).

## Sample infrastructure

There is a sample terraform script for you to build infrastructure on Google cloud.  
You could use this when running on dataflow.

```markdown
cd scripts/terraform
export TF_VAR_project=MY_PROJECT
export TF_VAR_region=MY_REGION

terraform plan
terraform apply
```