OCIR_URL := us-phoenix-1.ocir.io
TENANCY := ax0uvqijwxae
REGISTRY := bsidespdx2022
DOCKER_IMAGE := 50-bof

CONTAINER_TAG := ${OCIR_URL}/${TENANCY}/${REGISTRY}/${DOCKER_IMAGE}

target: container push deploy serve

clean: deleteDocker deleteKube

# Docker Configurations

container:
	cp ../distFiles/50-bof ../src/
	docker build --no-cache --tag ${CONTAINER_TAG} ../src/

push:
	docker push ${CONTAINER_TAG}

deleteDocker:
	docker image rm ${CONTAINER_TAG}

# Kubernetes

deploy:
	cat deployment/deploy.yaml | kubectl create --namespace bsidespdxctf  -f -

serve:
	kubectl create --namespace bsidespdxctf -f deployment/service.yaml

deleteKube:
	cat deployment/deploy.yaml | kubectl delete --namespace bsidespdxctf -f -
	kubectl delete --namespace bsidespdxctf -f deployment/service.yaml
	oci dns record domain delete --zone-name-or-id bsidespdxctf.party --domain basic-bof-1.bsidespdxctf.party --force