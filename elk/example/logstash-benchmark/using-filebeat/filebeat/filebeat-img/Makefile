.PHONY:	build push

IMAGE = filebeat
TAG = 0.1
USERNAME = rvu95

build:	
	docker build -t $(USERNAME)/$(IMAGE):$(TAG) .

push:	
	docker push $(USERNAME)/$(IMAGE):$(TAG)
