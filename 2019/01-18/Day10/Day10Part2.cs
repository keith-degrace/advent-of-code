using System;
using System.Collections.Generic;
using System.Linq;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day10Part2
    {
        class Asteroid
        {
            public int x;
            public int y;
            public Dictionary<Tuple<int, int>, List<Asteroid>> directions = new Dictionary<Tuple<int, int>, List<Asteroid>>();
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
                var signX = dx < 0 ? -1 : 1;
                var signY = dy < 0 ? -1 : 1;

                var gcd = Math.Abs(getGCD(dx, dy));
                dx /= gcd;
                dy /= gcd;
            }


            return Tuple.Create(dx, dy);
        }

        static int getDistance(Asteroid asteroid1, Asteroid asteroid2)
        {
            var dx = Math.Abs(asteroid1.x - asteroid2.x);
            var dy = Math.Abs(asteroid1.y - asteroid2.y);

            return dx + dy;
        }

        static double getAngle(int x1, int y1, int x2, int y2)
        {
            return (Math.Atan2(y2 - y1, x2 - x1) * 180 / Math.PI + 270) % 360;
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

        static Asteroid getBestMonitoringStation(List<Asteroid> asteroids)
        {
            Asteroid bestAsteroid = null;
            var bestCount = 0;

            foreach (Asteroid asteroid in asteroids)
            {
                if (asteroid.directions.Count > bestCount)
                {
                    bestAsteroid = asteroid;
                    bestCount = asteroid.directions.Count;
                }
            }

            return bestAsteroid;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("10");
            List<Asteroid> asteroids = load(input);

            for (int i = 0; i < asteroids.Count; i++)
            {
                for (int j = 0; j < asteroids.Count; j++)
                {
                    if (i == j)
                        continue;

                    var slope = getSlope(asteroids[i], asteroids[j]);

                    if (!asteroids[i].directions.ContainsKey(slope))
                        asteroids[i].directions[slope] = new List<Asteroid>();

                    asteroids[i].directions[slope].Add(asteroids[j]);
                }

                foreach (var asteroidsInDirection in asteroids[i].directions.Values)
                {
                    asteroidsInDirection.Sort((Asteroid a, Asteroid b) =>
                    {
                        var distance1 = getDistance(asteroids[i], a);
                        var distance2 = getDistance(asteroids[i], b);

                        return distance1.CompareTo(distance2);
                    });
                }
            }

            Asteroid monitoringStation = getBestMonitoringStation(asteroids);

//            Console.WriteLine(monitoringStation.x + ","+ monitoringStation.y);

            var order = new List<Tuple<int, int>>(monitoringStation.directions.Keys);
            order.Sort((Tuple<int, int> a, Tuple<int, int> b) =>
            {
                var angle1 = getAngle(0, 0, a.Item1, a.Item2);
                var angle2 = getAngle(0, 0, b.Item1, b.Item2);

                return angle1.CompareTo(angle2);
            });

            var vaporiseCount = 0;

            while (vaporiseCount < 200)
            {
                foreach (var tuple in order)
                {
                    var asteroidsInDirection = monitoringStation.directions[tuple];
                    if (asteroidsInDirection.Count > 0)
                    {
                        vaporiseCount++;

                        if (vaporiseCount == 200)
                            Console.WriteLine(asteroidsInDirection[0].x * 100 + asteroidsInDirection[0].y);

                        asteroidsInDirection.RemoveAt(0);
                    }
                }
            }
        }
    }
}
/*
 
    .#....#####...#..
    ##...##.#####..##
    ##...#...#.#####.
    ..#.....X...###..
    ..#.#.....#....##

*/