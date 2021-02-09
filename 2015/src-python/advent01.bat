@echo off

setlocal ENABLEDELAYEDEXPANSION

REM Part 1
set data=
for /f "delims=" %%i in ('type advent01.txt') do set data=!data! %%i

set floor=0

:loop1

    set char=!data:~0,1!
    set data=!data:~1!
    
    if "!char!"=="(" set /a floor += 1
    if "!char!"==")" set /a floor -= 1
    
    if "!data!" neq "" goto loop1

echo !floor!


REM Part 2
set data=
for /f "delims=" %%i in ('type advent01.txt') do set data=!data! %%i

set floor=0
set position=0

:loop2

    set char=!data:~0,1!
    set data=!data:~1!
    
    if "!char!"=="(" set /a floor += 1
    if "!char!"==")" set /a floor -= 1
    
    if "!floor!" neq "-1" (
        set /a position += 1
        goto loop2
    )

echo !position!


