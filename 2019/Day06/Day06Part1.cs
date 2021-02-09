using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day06Part1
    {
        class Planet
        {
            public Planet orbitsAround;

            public int getTotalOrbitCount()
            {
                return getDirectOrbitCount() + getIndirectOrbitCount();
            }

            public int getDirectOrbitCount()
            {
                return orbitsAround != null ? 1 : 0;
            }

            public int getIndirectOrbitCount()
            {
                return orbitsAround != null ? orbitsAround.getTotalOrbitCount() : 0;
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("06");

            var planets = new Dictionary<string, Planet>();
            foreach (var line in input) {
                var planetIds = line.Split(")");

                Planet parent;
                if (planets.ContainsKey(planetIds[0]))
                    parent = planets[planetIds[0]];
                else
                    planets.Add(planetIds[0], parent = new Planet());

                Planet child;
                if (planets.ContainsKey(planetIds[1]))
                    child = planets[planetIds[1]];
                else
                    planets.Add(planetIds[1], child = new Planet());

                child.orbitsAround = parent;
            }

            var sum = 0;
            foreach (var planet in planets.Values)
            {
                sum += planet.getTotalOrbitCount();
            }

            Console.WriteLine(sum);

        }
    }
}
