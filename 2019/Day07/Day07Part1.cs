using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day07Part1
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
                Terminate = 99,
            };

            enum ParameterMode
            {
                Position = 0,
                Immediate = 1,
            }

            private static OpCode getOpCode(int input)
            {
                return (OpCode)(input % 100);
            }

            private static ParameterMode getParameterMode(int input, int which)
            {
                return (ParameterMode)((input / (100 * Math.Pow(10, which - 1))) % 10);
            }

            private static int getParameterValue(int[] program, int current, int parameter)
            {
                var paramMode = getParameterMode(program[current], parameter);
                var paramValue = program[current + parameter];

                return (paramMode == ParameterMode.Position) ? program[paramValue] : paramValue;
            }

            public int execute(int[] program, int[] inputs)
            {
                var currentInput = 0;
                var output = 0;

                var current = 0;
                while (true)
                {
                    var opCode = getOpCode(program[current]);

                    if (opCode == OpCode.Add)
                    {
                        var value1 = getParameterValue(program, current, 1);
                        var value2 = getParameterValue(program, current, 2);
                        var value3 = program[current + 3];

                        program[value3] = value1 + value2;

                        current += 4;
                    }
                    else if (opCode == OpCode.Mul)
                    {
                        var value1 = getParameterValue(program, current, 1);
                        var value2 = getParameterValue(program, current, 2);
                        var value3 = program[current + 3];

                        program[value3] = value1 * value2;

                        current += 4;
                    }
                    else if (opCode == OpCode.Input)
                    {
                        program[program[current + 1]] = inputs[currentInput++];

                        current += 2;
                    }
                    else if (opCode == OpCode.Output)
                    {
                        output = program[program[current + 1]];

                        current += 2;
                    }
                    else if (opCode == OpCode.JumpIfTrue)
                    {
                        var value1 = getParameterValue(program, current, 1);
                        var value2 = getParameterValue(program, current, 2);

                        if (value1 != 0)
                            current = value2;
                        else
                            current += 3;
                    }
                    else if (opCode == OpCode.JumpIfFalse)
                    {
                        var value1 = getParameterValue(program, current, 1);
                        var value2 = getParameterValue(program, current, 2);

                        if (value1 == 0)
                            current = value2;
                        else
                            current += 3;
                    }
                    else if (opCode == OpCode.LessThan)
                    {
                        var value1 = getParameterValue(program, current, 1);
                        var value2 = getParameterValue(program, current, 2);
                        var value3 = program[current + 3];

                        program[value3] = value1 < value2 ? 1 : 0;

                        current += 4;
                    }
                    else if (opCode == OpCode.Equals)
                    {
                        var value1 = getParameterValue(program, current, 1);
                        var value2 = getParameterValue(program, current, 2);
                        var value3 = program[current + 3];

                        program[value3] = value1 == value2 ? 1 : 0;

                        current += 4;
                    }
                    else if (opCode == OpCode.Terminate)
                    {
                        break;
                    }
                }

                return output;
            }
        }
    
        public static int getMaxSignal(int[] program, List<int> sequence, int lastSignal)
        {
            // Assert sequence.Length > 0

            if (sequence.Count == 1) {
                var computer = new IntcodeComputer();
                return computer.execute(program, new int[] { sequence[0], lastSignal });
            }

            var maxSignal = 0;

            for (var i = 0; i<sequence.Count; i++)
            {
                var computer = new IntcodeComputer();
                var signal = computer.execute(program, new int[] { sequence[i], lastSignal });

                var subSequence = new List<int>(sequence);
                subSequence.RemoveAt(i);

                var subSignal = getMaxSignal(program, subSequence, signal);

                maxSignal = Math.Max(maxSignal, subSignal);
            }

            return maxSignal;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("07").Split(",");

            var program = new int[input.Length];
            for (int i = 0; i < input.Length; i++)
                program[i] = Int32.Parse(input[i]);

            var maxSignal = getMaxSignal(program, new List<int> { 0, 1, 2, 3, 4 }, 0);

            Console.WriteLine(maxSignal);
        }
    }
}
