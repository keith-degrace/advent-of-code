data = open('advent01.txt').read()

# Part 1
print data.count('(') - data.count(')')
   
# Part 2
floor = 0;

for i in range(0, len(data)):

    if data[i] == '(':
        floor = floor + 1;
    else:
        floor = floor - 1;

    if floor == -1:
        print i + 1
        break
