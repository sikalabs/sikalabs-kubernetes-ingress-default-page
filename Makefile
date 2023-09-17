IMAGE = sikalabs/sikalabs-kubernetes-ingress-default-page

prettier-check:
	yarn run prettier-check

prettier-write:
	yarn run prettier-write

dev:
	slu serve-files

build-and-push:
	docker build --platform linux/amd64 -t $(IMAGE) .
	docker push $(IMAGE)
