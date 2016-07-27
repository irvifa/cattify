.PHONY:	build push

IMAGE = lsth-flbt
TAG = 0.7
USERNAME = rvu95

build:	
	docker build -t $(USERNAME)/$(IMAGE):$(TAG) .

push:	
	docker push $(USERNAME)/$(IMAGE):$(TAG)
