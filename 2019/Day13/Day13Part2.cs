using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day13Part2
    {
        class IntcodeComputer
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
                AdjustRelativeBase = 9,
                Terminate = 99,
            };

            enum ParameterMode
            {
                Position = 0,
                Immediate = 1,
                Relative = 2,
            }

            Dictionary<Int64, Int64> program;
            Int64 current = 0;
            Int64 relativeBase = 0;
            bool halted = false;


            public IntcodeComputer(string[] input)
            {
                program = new Dictionary<Int64, Int64>();
                for (int i = 0; i < input.Length; i++)
                    program[i] = Int64.Parse(input[i]);
            }

            public bool isHalted()
            {
                return halted;
            }

            private static OpCode getOpCode(Int64 input)
            {
                return (OpCode)(input % 100);
            }

            private static ParameterMode getParameterMode(Int64 input, int which)
            {
                return (ParameterMode)((input / (100 * Math.Pow(10, which - 1))) % 10);
            }

            private Int64 getValue(Int64 position)
            {
                return program.ContainsKey(position) ? program[position] : 0;
            }

            private Int64 getParameterValue(Int64 position, int parameter)
            {
                var value = getValue(position);
                var paramMode = getParameterMode(value, parameter);
                var paramValue = getValue(position + parameter);

                Int64 finalValue;

                if (paramMode == ParameterMode.Position)
                    finalValue = getValue(paramValue);
                else if (paramMode == ParameterMode.Relative)
                    finalValue = getValue(relativeBase + paramValue);
                else
                    finalValue = paramValue;

                return finalValue;
            }

            private Int64 getWriteParameterIndex(Int64 position, int parameter)
            {
                var value = getValue(position);
                var paramMode = getParameterMode(value, parameter);
                var paramValue = getValue(position + parameter);

                return paramMode == ParameterMode.Position ? paramValue : relativeBase + paramValue;
            }

            public Int64[] execute(int[] inputs)
            {
                Int64 currentInput = 0;
                List<Int64> outputs = new List<Int64>();

                while (true)
                {
                    var currentValue = getValue(current);

                    var opCode = getOpCode(currentValue);

                    if (opCode == OpCode.Add)
                    {
                        var param1 = getParameterValue(current, 1);
                        var param2 = getParameterValue(current, 2);
                        var param3 = getWriteParameterIndex(current, 3);

                        program[param3] = param1 + param2;

                        current += 4;
                    }
                    else if (opCode == OpCode.Mul)
                    {
                        var param1 = getParameterValue(current, 1);
                        var param2 = getParameterValue(current, 2);
                        var param3 = getWriteParameterIndex(current, 3);

                        program[param3] = param1 * param2;

                        current += 4;
                    }
                    else if (opCode == OpCode.Input)
                    {
                        var param1 = getWriteParameterIndex(current, 1);

                        if (currentInput == inputs.Length)
                            break;

                        program[param1] = inputs[currentInput++];

                        current += 2;
                    }
                    else if (opCode == OpCode.Output)
                    {
                        var param1 = getParameterValue(current, 1);

                        outputs.Add(param1);

                        current += 2;
                    }
                    else if (opCode == OpCode.JumpIfTrue)
                    {
                        var param1 = getParameterValue(current, 1);
                        var param2 = getParameterValue(current, 2);

                        if (param1 != 0)
                            current = param2;
                        else
                            current += 3;
                    }
                    else if (opCode == OpCode.JumpIfFalse)
                    {
                        var param1 = getParameterValue(current, 1);
                        var param2 = getParameterValue(current, 2);

                        if (param1 == 0)
                            current = param2;
                        else
                            current += 3;
                    }
                    else if (opCode == OpCode.LessThan)
                    {
                        var param1 = getParameterValue(current, 1);
                        var param2 = getParameterValue(current, 2);
                        var param3 = getWriteParameterIndex(current, 3);

                        program[param3] = param1 < param2 ? 1 : 0;

                        current += 4;
                    }
                    else if (opCode == OpCode.Equals)
                    {
                        var param1 = getParameterValue(current, 1);
                        var param2 = getParameterValue(current, 2);
                        var param3 = getWriteParameterIndex(current, 3);

                        program[param3] = param1 == param2 ? 1 : 0;

                        current += 4;
                    }
                    else if (opCode == OpCode.AdjustRelativeBase)
                    {
                        var param1 = getParameterValue(current, 1);

                        relativeBase += param1;

                        current += 2;
                    }
                    else if (opCode == OpCode.Terminate)
                    {
                        halted = true;
                        break;
                    }
                }

                return (Int64[])outputs.ToArray();
            }
        }

        public static bool hasBlocksLeft(Dictionary<Tuple<Int64, Int64>, Int64> screen)
        {
            foreach (var entry in screen)
            {
                if (entry.Value == 2)
                    return true;
            }

            return false;
        }

        public static void draw(Dictionary<Tuple<Int64, Int64>, Int64> screen)
        {
            var minX = Int64.MaxValue;
            var minY = Int64.MaxValue;
            var maxX = Int64.MinValue;
            var maxY = Int64.MinValue;

            foreach (var entry in screen)
            {
                minX = Math.Min(minX, entry.Key.Item1);
                minY = Math.Min(minY, entry.Key.Item2);
                maxX = Math.Max(maxX, entry.Key.Item1);
                maxY = Math.Max(maxY, entry.Key.Item2);
            }

            for (var y = minY; y <= maxY; y++)
            {
                for (var x = minX; x <= maxX; x++)
                {
                    var pos = Tuple.Create(x, y);

                    if (!screen.ContainsKey(pos) || screen[pos] == 0)
                        Console.Write(' ');
                    else if (screen[pos] == 1)
                        Console.Write('#');
                    else if (screen[pos] == 2)
                        Console.Write('D');
                    else if (screen[pos] == 3)
                        Console.Write('-');
                    else if (screen[pos] == 4)
                        Console.Write('o');
                }
                Console.WriteLine();
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("13").Split(",");

            input[0] = "2";

            var screen = new Dictionary<Tuple<Int64, Int64>, Int64>();
            var computer = new IntcodeComputer(input);

            var joystick = 0;
            Int64 score = 0;
            do
            {
                var output = computer.execute(new int[] { joystick });

                Int64 ballX = 0;
                Int64 paddleX = 0;
                for (int i = 0; i < output.Length / 3; i++)
                {
                    var offset = i * 3;

                    var x = output[offset];
                    var y = output[offset + 1];
                    var position = Tuple.Create(x, y);
                    var value = output[offset + 2];

                    if (x == -1 && y == 0)
                        score = value;
                    else
                        screen[position] = value;

                    if (value == 3)
                        paddleX = x;

                    if (value == 4)
                        ballX = x;
                }

                if (ballX > paddleX)
                    joystick = 1;
                else if (ballX < paddleX)
                    joystick = -1;
                else
                    joystick = 0;
            }
            while (hasBlocksLeft(screen));

            Console.WriteLine(score);
        }
    }
}
