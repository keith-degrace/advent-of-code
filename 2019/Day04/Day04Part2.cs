using System;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day04Part2
    {
        private static bool hasTwoAdjacentChars(string password)
        {
            if (password.Length < 2)
                return false;

            var groupSize = 1;
            for (var i = 0; i<password.Length-1; i++)
            {
                if (password[i] == password[i + 1])
                {
                    groupSize++;
                }
                else
                {
                    if (groupSize == 2)
                        return true;

                    groupSize = 1;
                }
            }

            return groupSize == 2;
        }

        private static bool isNeverDecreasing(string password)
        {
            if (password.Length > 1)
            {
                for (var i = 1; i < password.Length; i++)
                {
                    if (password[i] < password[i - 1])
                        return false;
                }
            }

            return true;
        }

        private static bool isValid(string password)
        {
            return hasTwoAdjacentChars(password) && isNeverDecreasing(password);
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("04").Split("-");

            var min = Int32.Parse(input[0]);
            var max = Int32.Parse(input[1]);

            var validCount = 0;

            for (var value = min; value <= max; value++)
            {
                if (isValid(String.Format("{0}", value)))
                    validCount++;
            }

            Console.WriteLine(validCount);
        }
    }
}
