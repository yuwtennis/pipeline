# pipeline

A pipeline powered by apache beam for general use.  
I would like to introduce this as practical example.

## Available pipelines

### realestate_download

This pipeline targets mansion price in Japan which is downloaded from [MLIT](https://www.land.mlit.go.jp/webland/servlet/MainServlet).  
Data will be indexed into elasticsearch.

#### Sample dashboard using kibana

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
