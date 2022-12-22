using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day18Part1
    {
        public static void solve()
        {
            var input = "  this   is a test ";

            var stack = new LinkedList<string>();

            var currentWord = "";

            foreach (var c in input)
            {
                if (c != ' ')
                {
                    currentWord += c;
                }
                else if (currentWord.Length > 0)
                {
                    stack.AddFirst(currentWord);
                    currentWord = "";
                }
            }

            foreach (var word in stack)
            {
                Console.Write(word + " ");
            }
        }
    }
}
