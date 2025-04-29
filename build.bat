@echo off
rsrc -manifest app.manifest -ico icon.ico -o rsrc.syso
go build -ldflags "-H windowsgui"
move tuesday.exe Tuesday.exe
