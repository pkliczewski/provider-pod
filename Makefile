QUAY_HOST = quay.io
QUAY_ORG = pkliczewski
IMAGENAME = provider-pod
TAG = latest

all: build push

build:
	docker build -t $(QUAY_HOST)/$(QUAY_ORG)/$(IMAGENAME):$(TAG) .

push:
	docker push $(QUAY_HOST)/$(QUAY_ORG)/$(IMAGENAME):$(TAG)

.PHONY: all build push
