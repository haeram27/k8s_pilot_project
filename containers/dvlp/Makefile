TARGET = kubectl

BUILD_TAG = agent_pilot:dvlp

all :
ifeq ("$(wildcard $(TARGET))","")
	curl -LO "https://dl.k8s.io/release/v1.26.0/bin/linux/amd64/kubectl"
endif
	docker build -t $(BUILD_TAG) .

.PHONY: push
push :
	docker tag $(BUILD_TAG) sanmaul/$(BUILD_TAG)
	docker push sanmaul/$(BUILD_TAG)

clean :
	rm -rf $(TARGET)
