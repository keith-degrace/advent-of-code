
using System;
using System.Collections.Generic;
using System.Linq;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day01Part1
    {
        class Asteroid
        {
            public int x;
            public int y;
            public HashSet<Tuple<int, int>> asteroidSight = new HashSet<Tuple<int, int>>();
        }

        static int getGCD(int[] numbers)
        {
            return numbers.Aggregate(getGCD);
        }

        static int getGCD(int a, int b)
        {
            return b == 0 ? a : getGCD(b, a % b);
        }

        static Tuple<int, int> getSlope(Asteroid asteroid1, Asteroid asteroid2)
        {
            var dx = asteroid1.x - asteroid2.x;
            var dy = asteroid1.y - asteroid2.y;

            if (dx == 0)
            {
                dy /= Math.Abs(dy);
            }
            else if (dy == 0)
            {
                dx /= Math.Abs(dx);
            }
            else
            {
                // Console.WriteLine("Before : " + dx + "/" + dy);

                var signX = dx < 0 ? -1 : 1;
                var signY = dy < 0 ? -1 : 1;

                var gcd = Math.Abs(getGCD(dx, dy));
                dx /= gcd;
                dy /= gcd;

                // Console.WriteLine("After  : " + dx + "/" + dy + "   (" + gcd + ")");
            }


            return Tuple.Create(dx, dy);
        }

        static List<Asteroid> load(string[] input)
        {
            List<Asteroid> asteroids = new List<Asteroid>();

            for (var y = 0; y < input.Length; y++)
            {
                for (var x = 0; x < input[y].Length; x++)
                {
                    if (input[y][x] == '#')
                    {
                        var asteroid = new Asteroid();
                        asteroid.x = x;
                        asteroid.y = y;
                        asteroids.Add(asteroid);
                    }
                }
            }

            return asteroids;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("10");
            List<Asteroid> asteroids = load(input);

            for (int i = 0; i < asteroids.Count; i++)
            {
                for (int j = 0; j < asteroids.Count; j++)
                {
                    if (i != j)
                        asteroids[i].asteroidSight.Add(getSlope(asteroids[i], asteroids[j]));
                }
            }

            var best = 0;
            foreach (var asteroid in asteroids)
                best = Math.Max(best, asteroid.asteroidSight.Count);

            Console.WriteLine(best);
        }
    }
}
