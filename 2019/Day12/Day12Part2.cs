using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day12Part2
    {
        class Moon
        {
            public int[] initial = new int[3];
            public int[] current = new int[3];
            public int[] velocity = new int[3];

            public bool isInitial(int dimension)
            {
                return (current[dimension] == initial[dimension]) && (velocity[dimension] == 0);
            }

            override public string ToString()
            {
                return "[" + current[0] + "," + current[1] + "," + current[2] + "][" + velocity[0] + "," + velocity[1] + "," + velocity[2] + "]";
            }
        }

        static List<Moon> load(string[] input)
        {
            var regex = new Regex("<x=(.+), y=(.+), z=(.+)>");

            var moons = new List<Moon>();
            foreach (var line in input)
            {
                MatchCollection matches = regex.Matches(line);

                var moon = new Moon();

                moon.initial[0] = int.Parse(matches[0].Groups[1].Value);
                moon.initial[1] = int.Parse(matches[0].Groups[2].Value);
                moon.initial[2] = int.Parse(matches[0].Groups[3].Value);

                moon.current = (int[]) moon.initial.Clone();

                moons.Add(moon);
            }

            return moons;
        }
        public static Int64 lcm(Int64 a, Int64 b)
        {
            Int64 num1, num2;
            if (a > b)
            {
                num1 = a; num2 = b;
            }
            else
            {
                num1 = b; num2 = a;
            }

            for (int i = 1; i < num2; i++)
            {
                Int64 mult = num1 * i;
                if (mult % num2 == 0)
                {
                    return mult;
                }
            }
            return num1 * num2;
        }

        static Int64 getStepsToRepeat(List<Moon> moons, int dimension)
        {
            for (Int64 step = 0; ;step++)
            {
                // Apply gravity.
                for (var i = 0; i < moons.Count - 1; i++)
                {
                    for (var j = i + 1; j < moons.Count; j++)
                    {
                        var gravity = moons[i].current[dimension] < moons[j].current[dimension] ? 1 : moons[i].current[dimension] > moons[j].current[dimension] ? -1 : 0;
                        moons[i].velocity[dimension] += gravity;
                        moons[j].velocity[dimension] -= gravity;
                    }
                }

                // Apply velocity
                foreach (var moon in moons)
                    moon.current[dimension] += moon.velocity[dimension];

                // Check if we're back to initial.
                var backToInitial = true;
                foreach (var moon in moons)
                {
                    if (!moon.isInitial(dimension))
                    {
                        backToInitial = false;
                        break;
                    }
                }

                if (backToInitial)
                    return step + 1;
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("12");

            var moons = load(input);

            var stepX = getStepsToRepeat(moons, 0);
            var stepY = getStepsToRepeat(moons, 1);
            var stepZ = getStepsToRepeat(moons, 2);

            Console.WriteLine(lcm(lcm(stepX, stepY), stepZ));
        }
    }
}

