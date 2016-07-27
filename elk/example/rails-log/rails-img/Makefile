.PHONY: build push

IMAGE = rails-log
TAG = 0.1
PROJECT =  rvu95

build:	
	docker build -t $(PROJECT)/$(IMAGE):$(TAG) .

push:	
	docker push $(PROJECT)/$(IMAGE):$(TAG)



