using System;
using System.Collections.Generic;
using System.Linq;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day11Part1
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

            public IntcodeComputer(Dictionary<Int64, Int64> program)
            {
                this.program = new Dictionary<Int64, Int64>(program);
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

        enum Direction
        {
            Up = 0,
            Right = 1,
            Down = 2,
            Left = 3,
        }

        static Direction rotateLeft(Direction direction)
        {
            switch (direction)
            {
                case Direction.Up:
                    return Direction.Left;
                case Direction.Left:
                    return Direction.Down;
                case Direction.Down:
                    return Direction.Right;
                default:
                    return Direction.Up;
            }
        }

        static Direction rotateRight(Direction direction)
        {
            switch (direction)
            {
                case Direction.Up:
                    return Direction.Right;
                case Direction.Right:
                    return Direction.Down;
                case Direction.Down:
                    return Direction.Left;
                default:
                    return Direction.Up;
            }
        }

        static Tuple<int, int> move(Tuple<int, int> pos, Direction direction)
        {
            switch (direction)
            {
                case Direction.Up:
                    return Tuple.Create(pos.Item1, pos.Item2 - 1);
                case Direction.Right:
                    return Tuple.Create(pos.Item1 + 1, pos.Item2);
                case Direction.Down:
                    return Tuple.Create(pos.Item1, pos.Item2 + 1);
                default:
                    return Tuple.Create(pos.Item1 - 1, pos.Item2);
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("11").Split(",");

            var program = new Dictionary<Int64, Int64>();
            for (int i = 0; i < input.Length; i++)
                program[i] = Int64.Parse(input[i]);

            var computer = new IntcodeComputer(program);

            var panels = new Dictionary<Tuple<int, int>, int>();

            var robotPos = Tuple.Create(0, 0);
            var robotDir = Direction.Up;

            while (!computer.isHalted())
            {
                var currentColor = 0;
                if (panels.ContainsKey(robotPos))
                    currentColor = panels[robotPos];

                // Compute
                var outputs = computer.execute(new int[] { currentColor });

                // Rotate
                if (outputs[1] == 0)
                    robotDir = rotateLeft(robotDir);
                else
                    robotDir = rotateRight(robotDir);

                // Paint
                panels[robotPos] = (int) outputs[0];

                // Move
                robotPos = move(robotPos, robotDir);
            }

            Console.WriteLine(panels.Count);
        }
    }
}
