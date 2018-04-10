build:
	@sh -c "'$(CURDIR)/build.sh'"

push:
	@sh -c "'$(CURDIR)/push.sh'"

build_develop:
	docker build -t "bdc/bdc:develop" -f Dockerfile.develop .

push_develop:
	docker push "bdc/bdc:develop"

.PHONY: build build_develop push push_develop
