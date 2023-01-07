APPS = apps.exe srvlookup.exe

TARGET = $(APPS)
FUNCS  = 

all : $(TARGET)

apps.exe : apps.go
	go build -o apps.exe apps.go
	
srvlookup.exe : srvlookup.go
	go build -o srvlookup.exe srvlookup.go

clean :
	rm -f ${TARGET}
