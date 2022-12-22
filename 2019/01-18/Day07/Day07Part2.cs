using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day07Part2
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

            int output;
            int current;
            bool halted;
            int[] program;

            public IntcodeComputer(int[] program)
            {
                this.output = 0;
                this.current = 0;
                this.halted = false;
                this.program = (int[])program.Clone();
            }

            public bool isHalted()
            {
                return halted;
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

            public int execute(int[] inputs)
            {
                var currentInput = 0;

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
                        if (currentInput == inputs.Length)
                            return output;

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
                        halted = true;
                        break;
                    }
                }

                return output;
            }
        }

        public static void permutate(List<int> sequence, List<int> remaining, Action<List<int>> apply)
        {
            if (remaining.Count == 0)
                apply(sequence);

            for (var i = 0; i < remaining.Count; i++)
            {
                List<int> subSequence = new List<int>(sequence);
                subSequence.Add(remaining[i]);

                List<int> subRemaining = new List<int>(remaining);
                subRemaining.RemoveAt(i);

                permutate(subSequence, subRemaining, apply);
            }

        }

        static bool allHalted(IntcodeComputer[] amplifiers)
        {
            foreach (var amplifier in amplifiers)
            {
                if (!amplifier.isHalted())
                    return false;
            }

            return true;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("07").Split(",");

            var program = new int[input.Length];
            for (int i = 0; i < input.Length; i++)
                program[i] = Int32.Parse(input[i]);

            var max = 0;

            permutate(new List<int> { }, new List<int> { 5, 6, 7, 8, 9 }, (List<int> sequence) =>
             {
                 var amplifiers = new IntcodeComputer[] {
                    new IntcodeComputer(program),
                    new IntcodeComputer(program),
                    new IntcodeComputer(program),
                    new IntcodeComputer(program),
                    new IntcodeComputer(program)
                };

                 var lastSignal = 0;

                 for (var i = 0; i < sequence.Count; i++)
                     lastSignal = amplifiers[i].execute(new int[] { sequence[i], lastSignal });

                 while (!allHalted(amplifiers))
                 {
                     for (var i = 0; i < amplifiers.Length; i++)
                         lastSignal = amplifiers[i].execute(new int[] { lastSignal });
                 }

                 max = Math.Max(max, lastSignal);
             });

            Console.WriteLine(max);
        }
    }
}
