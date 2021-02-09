using System;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day01Part2
    {
        private static int getRequiredFuel(int mass)
        {
            if (mass == 0)
                return 0;

            var requiredFuel = Math.Max((int) Math.Floor(mass / 3.0) - 2, 0);

            return requiredFuel + getRequiredFuel(requiredFuel);
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("01");

            var sum = 0;
            foreach (var line in input)
            {
                var mass = Int32.Parse(line);
                sum += getRequiredFuel(mass);
            }

            Console.WriteLine(sum);
        }
    }
}
