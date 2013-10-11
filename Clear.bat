@echo off
title Clear
color 3
for /r . %%B in (*.log *.exe) do (
	echo "%%B"
	del "%%B"
)
pause