containers = [int(line) for line in open('advent17.txt').readlines()]

combinations_of_150 = []

for i in range(0, 2**len(containers)):
    
    liters = 0
    container_count = 0
    
    for container_index in range(0, len(containers)):
        
        if ((i >> container_index) & 1) == 1:
            liters += containers[container_index]
            container_count += 1

    if liters == 150:
        combinations_of_150.append(container_count)

print len(combinations_of_150)

print combinations_of_150.count(min(combinations_of_150))