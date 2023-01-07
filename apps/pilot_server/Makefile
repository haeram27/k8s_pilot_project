GO = go

APPS = pilot_server

TARGET = $(APPS)

all : $(TARGET)

$(APPS) : 
	${GO} build -buildvcs=false

.PHONY: tidy
tidy :
	go mod tidy
	
.PHONY: clean
clean :
	rm -f ${TARGET}
