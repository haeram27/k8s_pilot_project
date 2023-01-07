APPS = pilot_agent

TARGET = $(APPS)
FUNCS  = */*.go apps/*/*.go controllers/*/*.go

all : $(TARGET)

$(APPS) : main.go $(FUNCS)
	go build -buildvcs=false

.PHONY: tidy
tidy :
	go mod tidy
	
.PHONY: clean
clean :
	rm -f ${TARGET}
