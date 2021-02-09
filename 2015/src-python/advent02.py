gifts = []
for line in open('advent02.txt').readlines():
    gifts.append([int(n) for n in line.split('x')])

# Part 1
totalRequiredPaper = 0

for gift in gifts:

    surface1 = gift[0] * gift[1]
    surface2 = gift[1] * gift[2]
    surface3 = gift[2] * gift[0]

    area = 2 * surface1 + 2 * surface2 + 2 * surface3

    smallestSurface = min(surface1, surface2, surface3)
    
    requiredPaper = area + smallestSurface

    totalRequiredPaper += requiredPaper
    
print totalRequiredPaper

# Part 2
totalRequiredRibbon = 0

for gift in gifts:

    perimeter1 = gift[0] + gift[0] + gift[1] + gift[1]
    perimeter2 = gift[1] + gift[1] + gift[2] + gift[2]
    perimeter3 = gift[2] + gift[2] + gift[0] + gift[0]

    volume = gift[0] * gift[1] * gift[2]
    
    requiredRibbonForBox = min(perimeter1, perimeter2, perimeter3)
    requiredRibbonForBow = volume

    requiredRibbon = requiredRibbonForBox + requiredRibbonForBow
    
    totalRequiredRibbon += requiredRibbon
    
print totalRequiredRibbon