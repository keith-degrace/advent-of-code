using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day06Part2
    {
        class Planet
        {
            public string id;
            public Planet orbitsAround;

            public Planet(string id)
            {
                this.id = id;
            }
        }

        static Dictionary<string, Planet> load(string[] input)
        {
            var planets = new Dictionary<string, Planet>();

            foreach (var line in input)
            {
                var planetIds = line.Split(")");

                Planet parent;
                if (planets.ContainsKey(planetIds[0]))
                    parent = planets[planetIds[0]];
                else
                    planets.Add(planetIds[0], parent = new Planet(planetIds[0]));

                Planet child;
                if (planets.ContainsKey(planetIds[1]))
                    child = planets[planetIds[1]];
                else
                    planets.Add(planetIds[1], child = new Planet(planetIds[1]));

                child.orbitsAround = parent;
            }

            return planets;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("06");

            var planets = load(input);

            var you = planets["YOU"];
            var san = planets["SAN"];

            LinkedList<string> youAncestors = new LinkedList<string>();
            for (var current = you.orbitsAround; current != null; current = current.orbitsAround)
                youAncestors.AddFirst(current.id);

            LinkedList<string> sanAncestors = new LinkedList<string>();
            for (var current = san.orbitsAround; current != null; current = current.orbitsAround)
                sanAncestors.AddFirst(current.id);


            var currentYou = youAncestors.First;
            var currentSan = sanAncestors.First;

            while (currentYou != null && currentSan != null)
            {
                if (currentYou.Value != currentSan.Value)
                    break;

                currentYou = currentYou.Next;
                currentSan = currentSan.Next;
            }

            var distanceToYou = 0;
            for (;  currentYou != null; currentYou = currentYou.Next)
                distanceToYou++;

            var distanceToSan = 0;
            for (; currentSan != null; currentSan = currentSan.Next)
                distanceToSan++;

            Console.WriteLine(distanceToSan + distanceToYou);
        }
    }
}
