import re

regex = re.compile('((.)\\2*)')

def apply(value):
    
    new_value = ''
    
    for match in regex.findall(value):
        new_value += str(len(match[0])) + match[1]
   
    return new_value

# Part 1
value = '3113322113'

for i in range(40):
    value = apply(value)
    
print len(value)

# Part 2
value = '3113322113'

for i in range(50):
    value = apply(value)
    
print len(value)