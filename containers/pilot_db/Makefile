BUILD_TAG = agent_pilot:db 

all : 
	docker build -t $(BUILD_TAG) .

.PHONY: push
push :
	docker tag $(BUILD_TAG) sanmaul/$(BUILD_TAG)
	docker push sanmaul/$(BUILD_TAG)
