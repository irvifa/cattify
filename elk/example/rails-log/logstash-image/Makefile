.PHONY:	build push

IMAGE = lsth-rb
TAG = 0.7

build:	
	docker build -t gcr.io/kube-simulation/$(IMAGE):$(TAG) .

push:	
	gcloud docker --server=gcr.io push gcr.io/kube-simulation/$(IMAGE):$(TAG)
