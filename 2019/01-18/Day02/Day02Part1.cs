using System;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day20Part1
    {
        public static void solve()
        {
            var input = InputLoader.loadAsString("02").Split(",");

            var codes = new int[input.Length];
            for (int i = 0; i < input.Length; i++)
                codes[i] = Int32.Parse(input[i]);

            codes[1] = 12;
            codes[2] = 2;

            for (int i=0; codes[i] != 99; i+= 4)
            {
                if (codes[i] == 1) {
                    codes[codes[i + 3]] = codes[codes[i + 1]] + codes[codes[i + 2]];
                }
                else if (codes[i] == 2) {
                    codes[codes[i + 3]] = codes[codes[i + 1]] * codes[codes[i + 2]];
                }
            }

            Console.WriteLine(codes[0]);
        }
    }
}
