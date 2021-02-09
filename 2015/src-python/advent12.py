import json

j = json.load(open('advent12.txt'))

def get_sum(node):
    
    if type(node) is int:
        return node
        
    if type(node) is dict:
        sum = 0
        
        for key in node:
            sum += get_sum(node[key])
            
        return sum
            
    if type(node) is list:
        sum = 0
        
        for child in node:
            sum += get_sum(child)
            
        return sum
    
    return 0
        
print get_sum(j)


def get_sum_2(node):
    
    if type(node) is int:
        return node
        
    if type(node) is dict:
        sum = 0
        
        if "red" in node.itervalues():
            return 0
        
        for key in node:                
            sum += get_sum_2(node[key])
            
        return sum
            
    if type(node) is list:
        sum = 0
        
        for child in node:
            sum += get_sum_2(child)
            
        return sum
    
    return 0
        
print get_sum_2(j)