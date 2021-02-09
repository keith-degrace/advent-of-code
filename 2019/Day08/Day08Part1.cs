using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day08Part1
    {

        public static void solve()
        {
            var input = InputLoader.loadAsString("08");
            var width = 25;
            var height = 6;
            var depth = input.Length / (width * height);

            var smallest = new int[] { int.MaxValue, 0, 0 };

            for (var z = 0; z < depth; z++)
            {
                var counts = new int[] { 0, 0, 0 };

                for (var y = 0; y < height; y++)
                {
                    for (var x = 0; x < width; x++)
                    {
                        var value = int.Parse(input[z * width * height + y * width + x] + "");

                        if (value < counts.Length)
                        {
                            counts[value]++;
                        }
                    }
                }

                if (counts[0] < smallest[0])
                    smallest = counts;
            }

            Console.WriteLine(smallest[1] * smallest[2]);
        }
    }
}
