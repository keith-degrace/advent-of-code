
def find1():
    houses = {}

    for elf in xrange(1, 1000000):
        for house in xrange(elf, 1000000, elf):
            if houses.has_key(house):
                houses[house] += elf * 10
            else:
                houses[house] = elf * 10
            
            if houses[house] >= 36000000:
                return house

def find2():
    houses = {}

    for elf in xrange(1, 1000000):
        
        house_count = 0
    
        for house in xrange(elf, 1000000, elf):
       
            if houses.has_key(house):
                houses[house] += elf * 11
            else:
                houses[house] = elf * 11
            
            if houses[house] >= 36000000:
                return house
               
            house_count += 1
            if house_count == 50:
                break
                
print find1()
print find2()