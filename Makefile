all: controller webhook  activator

controller:
	@echo "build controller image"
	./build/build-image.sh controller

webhook:
	@echo "build webhook image"
	./build/build-image.sh webhook

activator:
	@echo "build activator image"
	./build/build-image.sh activator
kpa:
	@echo "build autoscaler image"
	./build/build-image.sh kpa

ingressslb:
	@echo "build ingressslb image"
	./build/build-image.sh ingressslb

queue:
	@echo "build queue-proxy image"
	./build/build-image.sh queue

serving_controller:
	@echo "build serving-controller image"
	./build/build-image.sh serving-controller
