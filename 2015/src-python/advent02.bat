@echo off

setlocal ENABLEDELAYEDEXPANSION

REM Part 1
set totalRequiredPaper=0

for /f "delims=x tokens=1,2,3" %%i in ('type advent02.txt') do (

    set /a surface1 = %%i * %%j
    set /a surface2 = %%j * %%k
    set /a surface3 = %%k * %%i
    
    set /a area = 2 * !surface1! + 2 * !surface2! + 2 * !surface3!

    set smallestSurface=!surface1!
    if !surface2! LSS !smallestSurface! set smallestSurface=!surface2!
    if !surface3! LSS !smallestSurface! set smallestSurface=!surface3!
    
    set /a requiredPaper = !area! + !smallestSurface!

    set /a totalRequiredPaper += !requiredPaper!
)

echo !totalRequiredPaper!

REM Part 2
set totalRequiredRibbon=0

for /f "delims=x tokens=1,2,3" %%i in ('type advent02.txt') do (

    set /a perimeter1 = %%i + %%i + %%j + %%j
    set /a perimeter2 = %%j + %%j + %%k + %%k
    set /a perimeter3 = %%k + %%k + %%i + %%i

    set /a volume = %%i * %%j * %%k
    
    set requiredRibbonForBox=!perimeter1!
    if !perimeter2! LSS !requiredRibbonForBox! set requiredRibbonForBox=!perimeter2!
    if !perimeter3! LSS !requiredRibbonForBox! set requiredRibbonForBox=!perimeter3!
    
    set /a requiredRibbonForBow = !volume!
    set /a requiredRibbon = !requiredRibbonForBox! + !requiredRibbonForBow!

    set /a totalRequiredRibbon += !requiredRibbon!
)
    
echo !totalRequiredRibbon!