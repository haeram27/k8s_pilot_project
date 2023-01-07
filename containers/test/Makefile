TARGET = apps.exe

SRC_PATH = /src/agent_pilot/apps/test
SRC_APPS = $(SRC_PATH)/apps.exe

BUILD_POD = $(shell kubectl get pods -n agent-pilot | grep pilot-dvlp | head -n 1 | awk '{print $$1}')
BUILD_TAG = agent_pilot:test
	
all : $(TARGET)

apps.exe :
	/bin/sh -c 'cd $(SRC_PATH); make clean; git pull; go mod vendor;'
	kubectl exec $(BUILD_POD) -n agent-pilot -- /bin/sh -c 'cd $(SRC_PATH); make;'
	cp -f $(SRC_APPS) .

	docker build -t $(BUILD_TAG) .

.PHONY: push
push :
	docker tag $(BUILD_TAG) sanmaul/$(BUILD_TAG)
	docker push sanmaul/$(BUILD_TAG)

clean :
	rm -rf $(TARGET)
