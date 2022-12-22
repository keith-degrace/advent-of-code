using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day03Part2
    {
        class Wire
        {
            int x;
            int y;
            int step;

            private Tuple<int, int> getSlope(char direction)
            {
                switch (direction)
                {
                    case 'U':
                        return Tuple.Create(0, 1);

                    case 'D':
                        return Tuple.Create(0, -1);

                    case 'L':
                        return Tuple.Create(-1, 0);

                    case 'R':
                        return Tuple.Create(1, 0);
                }

                throw new Exception();
            }

            public void move(string input, Action<int, int, int> track)
            {
                var direction = input[0];
                var amount = Int32.Parse(input.Substring(1));

                var slope = getSlope(direction);

                for (var d = 0; d < amount; d++)
                {
                    x += slope.Item1;
                    y += slope.Item2;
                    step++;

                    track(x, y, step);
                }
            }

            public void move(string[] inputs, Action<int, int, int> track)
            {
                foreach (var input in inputs)
                    move(input, track);
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsStringArray("03");

            var wire1Inputs = input[0].Split(",");
            var wire2Inputs = input[1].Split(",");

            var wire1 = new Wire();
            var wire2 = new Wire();

            var wire1Path = new Dictionary<Tuple<int, int>, int>();
            wire1.move(wire1Inputs, (int x, int y, int step) =>
            {
                if (!wire1Path.ContainsKey(Tuple.Create(x, y)))
                    wire1Path.Add(Tuple.Create(x, y), step);
            });

            var closest = Int32.MaxValue;
            wire2.move(wire2Inputs, (int x, int y, int step) =>
            {
                if (wire1Path.ContainsKey(Tuple.Create(x, y)))
                    closest = Math.Min(closest, step + wire1Path[Tuple.Create(x, y)]);
            });

            Console.WriteLine(closest);
        }
    }
}
