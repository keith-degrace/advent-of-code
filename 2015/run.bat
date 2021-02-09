@echo off

set GOPATH=%~dp0

go install au "%1"
if errorlevel 1 exit /b 1

bin\%~n1
