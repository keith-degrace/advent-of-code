using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day17Part2
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

        static LinkedList<string> getPath(IntcodeComputer computer)
        {
            var outputs = computer.execute(new int[] { });

            var robot = Tuple.Create(0, 0);

            var map = new Dictionary<Tuple<int, int>, char>();

            var x = 0;
            var y = 0;

            foreach (var output in outputs)
            {
                var pos = Tuple.Create(x, y);

                map[pos] = (char)output;

                if (output == '^')
                    robot = pos;

                if (output == 10)
                {
                    x = 0;
                    y++;
                }
                else
                {
                    x++;
                }
            }

            var speed = Tuple.Create(1, 0);
            var dir = 'R';
            var steps = 0;

            var path = new LinkedList<string>();
            while (true)
            {
                var next = Tuple.Create(robot.Item1 + speed.Item1, robot.Item2 + speed.Item2);

                if (map.ContainsKey(next) && map[next] == '#')
                {
                    steps++;
                    robot = next;
                }
                else
                {
                    path.AddLast("" + dir + steps);

                    // Left?
                    var leftSpeed = Tuple.Create(speed.Item2, -speed.Item1);
                    var left = Tuple.Create(robot.Item1 + leftSpeed.Item1, robot.Item2 + leftSpeed.Item2);

                    if (map.ContainsKey(left) && map[left] == '#')
                    {
                        speed = leftSpeed;
                        dir = 'L';
                        steps = 0;
                    }
                    else
                    {
                        // Right?
                        var rightSpeed = Tuple.Create(-speed.Item2, speed.Item1);
                        var right = Tuple.Create(robot.Item1 + rightSpeed.Item1, robot.Item2 + rightSpeed.Item2);

                        if (map.ContainsKey(right) && map[right] == '#')
                        {
                            speed = rightSpeed;
                            dir = 'R';
                            steps = 0;
                        }
                        else
                        {
                            break;
                        }
                    }
                }
            }

            return path;
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("17").Split(",");

            // Wake up the robot.
            input[0] = "2";

            var computer = new IntcodeComputer(input);

            var path = getPath(computer);

            ////////////////////////////////////////////
            // Cheaped out and did the next part by hand.
            ////////////////////////////////////////////

            foreach (var entry in path)
                Console.Write(entry + ",");

            Console.WriteLine();

            // A - R10,L12,R6,
            // A - R10,L12,R6,
            // B - R6,R10,R12,R6,
            // C - R10,L12,L12,
            // B - R6,R10,R12,R6,
            // C - R10,L12,L12,
            // B - R6,R10,R12,R6,
            // C - R10,L12,L12,
            // B - R6,R10,R12,R6,
            // A - R10,L12,R6,

            var instructions = "A,A,B,C,B,C,B,C,B,A\nR,10,L,12,R,6\nR,6,R,10,R,12,R,6\nR,10,L,12,L,12\nn\n";

            var computerInput = new int[instructions.Length];
            for (var i = 0; i < instructions.Length; i++)
                computerInput[i] = instructions[i];

            var outputs = computer.execute(computerInput);
            Console.WriteLine(outputs[outputs.Length - 1]);
        }
    }
}
