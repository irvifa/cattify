.PHONY:	build push

IMAGE = filebeat
TAG = 0.5

build:	
	docker build -t gcr.io/kube-simulation/$(IMAGE):$(TAG) .

push:	
	gcloud docker --server=gcr.io push gcr.io/kube-simulation/$(IMAGE):$(TAG)
