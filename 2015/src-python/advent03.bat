@echo off

goto main

REM Part 1
:part1
    setlocal ENABLEDELAYEDEXPANSION
    
    for /f "delims=" %%i in ('type advent03.txt') do set data="%%i"

    set delivered=0

    set position_x=0
    set position_y=0

    :loop1

        set char=!data:~0,1!
        set data=!data:~1!
        
        if "!char!"=="<" set /a position_x += 1
        if "!char!"==">" set /a position_x -= 1
        if "!char!"=="^" set /a position_y += 1
        if "!char!"=="v" set /a position_y -= 1

        if !%position_x%x%position_y%! NEQ 1 (set /a delivered += 1)
        
        set %position_x%x%position_y%=1
            
        if "!data!" neq "" goto loop1
        
    echo !delivered!
    exit /b

REM Part 2
:part2
    setlocal ENABLEDELAYEDEXPANSION
    
    for /f "delims=" %%i in ('type advent03.txt') do set data="%%i"

    set delivered=0

    set santa_position_x=0
    set santa_position_y=0

    set robot_position_x=0
    set robot_position_y=0

    set turn=santa
    
    :loop2

        set char=!data:~0,1!
        set data=!data:~1!
        
        if "!turn!"=="santa" (
        
            if "!char!"=="<" set /a santa_position_x += 1
            if "!char!"==">" set /a santa_position_x -= 1
            if "!char!"=="^" set /a santa_position_y += 1
            if "!char!"=="v" set /a santa_position_y -= 1

            if !%santa_position_x%x%santa_position_y%! NEQ 1 (set /a delivered += 1)
            
            set %santa_position_x%x%santa_position_y%=1
            
            set turn=robot
            
        ) else (
            
            if "!char!"=="<" set /a robot_position_x += 1
            if "!char!"==">" set /a robot_position_x -= 1
            if "!char!"=="^" set /a robot_position_y += 1
            if "!char!"=="v" set /a robot_position_y -= 1

            if !%robot_position_x%x%robot_position_y%! NEQ 1 (set /a delivered += 1)
            
            set %robot_position_x%x%robot_position_y%=1
            
            set turn=santa
        )
            
        if "!data!" neq "" goto loop2
        
    echo !delivered!
    exit /b

:main
    call :part1
    call :part2
