APPS = cppm

TARGET = $(APPS)
FUNCS  = models/*.go controllers/*.go routers/*.go

$(APPS) : main.go $(FUNCS)
	go build -buildvcs=false

.PHONY: tidy
tidy :
	go mod tidy
	
.PHONY: clean
clean :
	rm -f ${TARGET} data.db lastupdate.tmp
