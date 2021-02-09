using System;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day02Part2
    {
        private static int executeProgram(int[] codes, int noun, int verb)
        {
            var codesCopy = (int[])codes.Clone();

            codesCopy[1] = noun;
            codesCopy[2] = verb;

            for (int i = 0; codesCopy[i] != 99; i += 4)
            {
                if (codesCopy[i] == 1)
                {
                    codesCopy[codesCopy[i + 3]] = codesCopy[codesCopy[i + 1]] + codesCopy[codesCopy[i + 2]];
                }
                else if (codesCopy[i] == 2)
                {
                    codesCopy[codesCopy[i + 3]] = codesCopy[codesCopy[i + 1]] * codesCopy[codesCopy[i + 2]];
                }
            }

            return codesCopy[0];
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("02").Split(",");

            var codes = new int[input.Length];
            for (int i = 0; i < input.Length; i++)
                codes[i] = Int32.Parse(input[i]);

            for (int noun = 0; noun <= 99; noun++)
            {
                for (int verb = 0; verb <= 99; verb++)
                {
                    var result = executeProgram(codes, noun, verb);
                    if (result == 19690720)
                    {
                        Console.WriteLine(100 * noun + verb);
                        return;
                    }
                }
            }
        }
    }
}
