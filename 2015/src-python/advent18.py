def is_light_on(grid, x, y):
    
    if (y < 0) or (y >= len(grid) or (x < 0) or (x >= len(grid[y]))):
        return False

    return grid[y][x] == '#'

def get_light_count(grid):
    
    light_count = 0

    for line in grid:
        light_count += line.count('#')
        
    return light_count
    
def generate_new_grid(grid):

    new_grid = [list(line) for line in grid]

    for y in range(len(grid)):
        for x in range(len(grid[y])):
        
            neighbooring_light_count = 0
            if is_light_on(grid, x-1, y-1):
                neighbooring_light_count += 1
            if is_light_on(grid, x-1,   y):
                neighbooring_light_count += 1
            if is_light_on(grid, x-1, y+1):
                neighbooring_light_count += 1
            if is_light_on(grid,   x, y+1):
                neighbooring_light_count += 1
            if is_light_on(grid, x+1, y+1):
                neighbooring_light_count += 1
            if is_light_on(grid, x+1,   y):
                neighbooring_light_count += 1
            if is_light_on(grid, x+1, y-1):
                neighbooring_light_count += 1
            if is_light_on(grid,   x, y-1):
                neighbooring_light_count += 1
            
            if grid[y][x] == '#':
                if neighbooring_light_count in [2,3]:
                    new_grid[y][x] = '#'
                else:
                    new_grid[y][x] = '.'
            else:
                if neighbooring_light_count == 3:
                    new_grid[y][x] = '#'
                else:
                    new_grid[y][x] = '.'
            
    return new_grid
   
# Part 1
grid = [list(line.strip()) for line in open('advent18.txt').readlines()]

for i in range(100):
    grid = generate_new_grid(grid)
    
print get_light_count(grid)
                
# Part 1
grid = [list(line.strip()) for line in open('advent18.txt').readlines()]
grid[0][0] = '#'
grid[0][99] = '#'
grid[99][0] = '#'
grid[99][99] = '#'

for i in range(100):
    grid = generate_new_grid(grid)
    grid[0][0] = '#'
    grid[0][99] = '#'
    grid[99][0] = '#'
    grid[99][99] = '#'
    
print get_light_count(grid)
            
   
   