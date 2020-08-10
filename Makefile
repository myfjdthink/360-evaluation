NAME=deoooo.tencentcloudcr.com/deo/360-evaluation
TAG=master

build:
	echo building ${NAME}:${TAG}
	docker build -t ${NAME}:${TAG} .
	docker push "${NAME}:${TAG}"
