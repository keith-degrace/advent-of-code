cities = []
distances = {}

for line in open('advent09.txt'):
    tokens = line.split()

    city1 = tokens[0]
    city2 = tokens[2]
    distance = int(tokens[4])

    if city1 not in cities:
        cities.append(city1)

    if city2 not in cities:
        cities.append(city2)

    distances[(city1,city2)] = distance
    distances[(city2,city1)] = distance

def get_permutations(cities):
    
    if len(cities) == 1:
        return [cities]

    permutations = []
    
    for city in cities:
        remaining_cities = list(cities)
        remaining_cities.remove(city)

        for sub_permutation in get_permutations(remaining_cities):
            permutations.append([city] + sub_permutation)

    return permutations

def get_distance(path):
    return sum(distances[(path[i-1], path[i])] for i in range(1, len(path)))

print min(get_distance(path) for path in get_permutations(cities))
print max(get_distance(path) for path in get_permutations(cities))
