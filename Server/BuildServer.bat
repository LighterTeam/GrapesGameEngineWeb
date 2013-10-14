@echo off
title BuildGame

del bin\GameServer.exe

set oldgopath=%GOPATH%
set GOPATH="%CD%\Common";"%CD%\GameServer";

cd GameServer
go build
copy GameServer.exe ..\bin

set GOPATH=%oldgopath%

pause

