import re

ingredients = []

for line in open('advent15.txt').readlines():
    r = re.match('(.*): capacity (.*), durability (.*), flavor (.*), texture (.*), calories (.*)', line)

    ingredient = {}
    ingredient['name'] = r.group(1)
    ingredient['capacity'] = int(r.group(2))
    ingredient['durability'] = int(r.group(3))
    ingredient['flavor'] = int(r.group(4))
    ingredient['texture'] = int(r.group(5))
    ingredient['calories'] = int(r.group(6))
    
    ingredients.append(ingredient)

def get_all_combinations(amount, ingredient_count):

    combinations = []
    
    for i in range(1, amount):
        amount_left = amount - i
        ingredients_left = ingredient_count - 1

        if amount_left == 0:
            combinations.append([i])
        elif ingredients_left == 1:
            combinations.append([i] + [amount_left])
        else:
            for sub_combination in get_all_combinations(amount_left, ingredients_left):
                combinations.append([i] + sub_combination)
    
    return combinations

max_score = 0
    
for combination in get_all_combinations(100, 4):

    capacity = 0
    durability = 0
    flavor = 0
    texture = 0
    calories = 0

    for i in range(len(ingredients)):
        capacity   += combination[i] * ingredients[i]['capacity']
        durability += combination[i] * ingredients[i]['durability']
        flavor     += combination[i] * ingredients[i]['flavor']
        texture    += combination[i] * ingredients[i]['texture']
        calories   += combination[i] * ingredients[i]['calories']

    capacity = max(0, capacity)
    durability = max(0, durability)
    flavor = max(0, flavor)
    texture = max(0, texture)
    calories = max(0, calories)
        
    score = capacity * durability * flavor * texture
    
    if calories == 500:
        max_score = max(max_score, score)
    
print max_score