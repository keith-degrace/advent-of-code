using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day14Part2
    {
        class Ingredient
        {
            public string name;
            public int amount;

            public Ingredient(string name, int amount)
            {
                this.name = name;
                this.amount = amount;
            }
        }

        class Recipe
        {
            public Ingredient produces;
            public List<Ingredient> requires = new List<Ingredient>();
        }

        static Dictionary<string, Recipe> parse(string[] input)
        {
            var recipies = new Dictionary<string, Recipe>();

            foreach (var line in input)
            {
                var recipeParts = line.Split(" => ");

                var recipe = new Recipe();

                foreach (var ingredientString in recipeParts[0].Split(", "))
                {
                    var ingredientParts = ingredientString.Split(" ");

                    var ingredient = new Ingredient(ingredientParts[1], int.Parse(ingredientParts[0]));

                    recipe.requires.Add(ingredient);
                }

                var producesParts = recipeParts[1].Split(" ");

                recipe.produces = new Ingredient(producesParts[1], int.Parse(producesParts[0]));

                recipies.Add(recipe.produces.name, recipe);
            }

            return recipies;
        }

        static int getOreCost(Dictionary<string, Recipe> recipies, Dictionary<string, int> availableIngredients, int ingredientAmount, string ingredientName)
        {
            if (ingredientName == "ORE")
                return ingredientAmount;

            var recipe = recipies[ingredientName];

            var recipeAmount = ((ingredientAmount - 1) / recipe.produces.amount) + 1;

            var cost = 0;

            foreach (var requiredIngredient in recipe.requires)
            {
                var amountToCost = requiredIngredient.amount * recipeAmount;
                if (availableIngredients.ContainsKey(requiredIngredient.name))
                {
                    var amountAvailableToUse = Math.Min(amountToCost, availableIngredients[requiredIngredient.name]);

                    amountToCost -= amountAvailableToUse;
                    availableIngredients[requiredIngredient.name] -= amountAvailableToUse;
                }

                if (amountToCost > 0)
                    cost += getOreCost(recipies, availableIngredients, amountToCost, requiredIngredient.name);
            }

            var remaining = (recipeAmount * recipe.produces.amount) - ingredientAmount;
            if (remaining > 0)
                availableIngredients[ingredientName] = availableIngredients.GetValueOrDefault(ingredientName) + remaining;

            return cost;
        }


        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("14");

            var recipies = parse(input);

            Dictionary<string, int> availableIngredients = new Dictionary<string, int>();

            var fuelCount = 0;
            var fuelAmountPerIteration = 1000;
            var ore = 1000000000000;
            while (true)
            {
                var cost = getOreCost(recipies, availableIngredients, fuelAmountPerIteration, "FUEL");
                if (cost > ore) {
                    if (fuelAmountPerIteration > 1)
                    {
                        fuelAmountPerIteration /= 2;
                        continue;
                    }
                    break;
                }

                ore -= cost;

                fuelCount += fuelAmountPerIteration;
            }

            Console.WriteLine(fuelCount);
        }
    }
}
