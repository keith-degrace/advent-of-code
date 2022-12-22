using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day12Part1
    {
        class Moon
        {
            public int x;
            public int y;
            public int z;
        
            public int vx = 0;
            public int vy = 0;
            public int vz = 0;

            override public string ToString()
            {
                return "[" + x + "," + y + "," + z +"][" + vx + "," + vy + "," + vz + "]";
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("12");

            var regex = new Regex("<x=(.+), y=(.+), z=(.+)>");

            var moons = new List<Moon>();
            foreach (var line in input)
            {
                MatchCollection matches = regex.Matches(line);

                var moon = new Moon();

                moon.x = int.Parse(matches[0].Groups[1].Value);
                moon.y = int.Parse(matches[0].Groups[2].Value);
                moon.z = int.Parse(matches[0].Groups[3].Value);

                moons.Add(moon);
            }

            for (var step = 0; step < 1000; step++)
            {
                // Apply gravity.
                for (var i = 0; i < moons.Count - 1; i++)
                {
                    for (var j = i + 1; j < moons.Count; j++)
                    {
                        var gravityX = moons[i].x < moons[j].x ? 1 : moons[i].x > moons[j].x ? -1 : 0;
                        moons[i].vx += gravityX;
                        moons[j].vx -= gravityX;

                        var gravityY = moons[i].y < moons[j].y ? 1 : moons[i].y > moons[j].y ? -1 : 0;
                        moons[i].vy += gravityY;
                        moons[j].vy -= gravityY;

                        var gravityZ = moons[i].z < moons[j].z ? 1 : moons[i].z > moons[j].z ? -1 : 0;
                        moons[i].vz += gravityZ;
                        moons[j].vz -= gravityZ;
                    }
                }

                // Apply velocity
                foreach (var moon in moons)
                {
                    moon.x += moon.vx;
                    moon.y += moon.vy;
                    moon.z += moon.vz;
                }
            }

            var sum = 0;
            foreach (var moon in moons)
                sum += (Math.Abs(moon.x) + Math.Abs(moon.y) + Math.Abs(moon.z)) * (Math.Abs(moon.vx) + Math.Abs(moon.vy) + Math.Abs(moon.vz));

            Console.WriteLine(sum);
        }
    }
}
