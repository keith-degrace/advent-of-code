using System;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day05Part2
    {
        enum OpCode
        {
            Add = 1,
            Mul = 2,
            Input = 3,
            Output = 4,
            JumpIfTrue = 5,
            JumpIfFalse = 6,
            LessThan = 7,
            Equals = 8,
            Terminate = 99,
        };

        private static OpCode getOpCode(int input)
        {
            return (OpCode)(input % 100);
        }

        enum ParameterMode
        {
            Position = 0,
            Immediate = 1,
        }

        private static ParameterMode getParameterMode(int input, int which)
        {
            return (ParameterMode)((input / (100 * Math.Pow(10, which-1))) % 10);
        }

        private static int getParameterValue(int[] values, int current, int parameter)
        {
            var paramMode = getParameterMode(values[current], parameter);
            var paramValue = values[current + parameter];

            return (paramMode == ParameterMode.Position) ? values[paramValue] : paramValue;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("05").Split(",");

            var values = new int[input.Length];
            for (int i = 0; i < input.Length; i++)
                values[i] = Int32.Parse(input[i]);

            var register = 5;

            var current = 0;
            while (true)
            {
                var opCode = getOpCode(values[current]);

                if (opCode == OpCode.Add)
                {
                    var value1 = getParameterValue(values, current, 1);
                    var value2 = getParameterValue(values, current, 2);
                    var value3 = values[current + 3];

                    values[value3] = value1 + value2;

                    current += 4;
                }
                else if (opCode == OpCode.Mul)
                {
                    var value1 = getParameterValue(values, current, 1);
                    var value2 = getParameterValue(values, current, 2);
                    var value3 = values[current + 3];

                    values[value3] = value1 * value2;

                    current += 4;
                }
                else if (opCode == OpCode.Input)
                {
                    values[values[current + 1]] = register;

                    current += 2;
                }
                else if (opCode == OpCode.Output)
                {
                    register = values[values[current + 1]];

                    current += 2;
                }
                else if (opCode == OpCode.JumpIfTrue)
                {
                    var value1 = getParameterValue(values, current, 1);
                    var value2 = getParameterValue(values, current, 2);

                    if (value1 != 0)
                        current = value2;
                    else
                        current += 3;
                }
                else if (opCode == OpCode.JumpIfFalse)
                {
                    var value1 = getParameterValue(values, current, 1);
                    var value2 = getParameterValue(values, current, 2);

                    if (value1 == 0)
                        current = value2;
                    else
                        current += 3;
                }
                else if (opCode == OpCode.LessThan)
                {
                    var value1 = getParameterValue(values, current, 1);
                    var value2 = getParameterValue(values, current, 2);
                    var value3 = values[current + 3];

                    values[value3] = value1 < value2 ? 1 : 0;

                    current += 4;
                }
                else if (opCode == OpCode.Equals)
                {
                    var value1 = getParameterValue(values, current, 1);
                    var value2 = getParameterValue(values, current, 2);
                    var value3 = values[current + 3];

                    values[value3] = value1 == value2 ? 1 : 0;

                    current += 4;
                }
                else if (opCode == OpCode.Terminate)
                {
                    break;
                }
            }

            Console.WriteLine(register);
        }
    }
}
