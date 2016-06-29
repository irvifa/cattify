.PHONY:	build push

IMAGE = filebeat-elasticsearch
TAG = 0.1

build:	
	docker build -t gcr.io/kube-simulation/$(IMAGE):$(TAG) .

push:	
	gcloud docker --server=gcr.io push gcr.io/kube-simulation/$(IMAGE):$(TAG)
