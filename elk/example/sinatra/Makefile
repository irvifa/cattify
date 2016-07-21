.PHONY:	build push

IMAGE = sinatra
TAG = 0.1
PROJECT =  kube-simulation

build:	
	docker build -t gcr.io/$(PROJECT)/$(IMAGE):$(TAG) .

push:	
	gcloud docker --server=gcr.io push gcr.io/$(PROJECT)/$(IMAGE):$(TAG)

