@echo off
poetry run python convert_icon.py icon.png
rsrc -manifest app.manifest -ico icon.ico -o rsrc.syso
go build
