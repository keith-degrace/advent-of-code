using System;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day16Part1
    {
        public static void solve()
        {
            var input = InputLoader.loadAsString("16");
            var signal = Array.ConvertAll(input.ToCharArray(), strDigit => int.Parse(strDigit + ""));

            var pattern = new int[] { 0, 1, 0, -1 };

            var nextSignal = new int[signal.Length];

            for (var phase = 0; phase < 100; phase++)
            {
                for (var i = 0; i < nextSignal.Length; i++)
                {
                    for (var j = 0; j < signal.Length; j++)
                    {
                        var patternIndex = ((((j + 1) / (i + 1)) % pattern.Length));
                        var patternDigit = pattern[patternIndex];

                        nextSignal[i] += signal[j] * patternDigit;
                    }

                    nextSignal[i] = Math.Abs(nextSignal[i]) % 10;
                }

                signal = nextSignal;
                nextSignal = new int[signal.Length];
            }

            for (var i=0; i<8; i++)
                Console.Write(signal[i]);
            Console.WriteLine();
        }
    }
}
