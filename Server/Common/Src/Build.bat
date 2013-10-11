set oldgopath=%GOPATH%
set GOPATH="%CD%\..\..\Common";

cd TSCommon
go build 
go install
cd ..
cd TSTCP
go build
go install
cd ..
cd TSUtil
go build
go install
cd ..
cd TSHTTP
go build
go install
cd ..
cd TSLog
go build
go install
cd ..
cd TSJson
go build
go install

set GOPATH=%oldgopath%
pause
