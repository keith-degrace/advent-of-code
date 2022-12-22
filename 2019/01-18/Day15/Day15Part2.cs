using System;
using System.Collections.Generic;
using Advent_of_Code_2019.Utils;

namespace Advent_of_Code_2019
{
    class Day15Part2
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

        enum Direction
        {
            North=1,
            South=2,
            West=3,
            East=4,
        }

        class Map
        {
            Dictionary<Tuple<int, int>, char> data = new Dictionary<Tuple<int, int>, char>();
            int minX = 0;
            int minY = 0;
            int maxX = 0;
            int maxY = 0;

            public void draw(Tuple<int, int> currentPosition)
            {
                for (var y = minY; y <= maxY; y++)
                {
                    for (var x = minX; x <= maxX; x++)
                    {
                        if (x == currentPosition.Item1 && y == currentPosition.Item2)
                            Console.Write('O');
                        else
                            Console.Write(data.GetValueOrDefault(Tuple.Create(x, y), ' '));
                    }

                    Console.WriteLine();
                }
            }

            public bool isVisited(Tuple<int, int> position)
            {
                return data.ContainsKey(position);
            }

            public bool isVisited(Tuple<int, int> position, Direction direction)
            {

                return isVisited(move(position, direction));
            }

            public char get(Tuple<int, int> position)
            {
                return data.GetValueOrDefault(position, ' ');
            }

            public void set(Tuple<int, int> position, char value)
            {
                data[position] = value;

                minX = Math.Min(minX, position.Item1);
                minY = Math.Min(minY, position.Item2);

                maxX = Math.Max(maxX, position.Item1);
                maxY = Math.Max(maxY, position.Item2);
            }
        }

        static Direction getOppositeDirection(Direction direction)
        {
            switch (direction)
            {
                case Direction.North:
                    return Direction.South;
                case Direction.South:
                    return Direction.North;
                case Direction.West:
                    return Direction.East;
                default:
                    return Direction.West;
            }
        }

        static Tuple<int, int> move(Tuple<int, int> position, Direction direction)
        {
            switch (direction)
            {
                case Direction.North:
                    return Tuple.Create(position.Item1, position.Item2 - 1);
                case Direction.South:
                    return Tuple.Create(position.Item1, position.Item2 + 1);
                case Direction.West:
                    return Tuple.Create(position.Item1 - 1, position.Item2);
                case Direction.East:
                    return Tuple.Create(position.Item1 + 1, position.Item2);
            }

            return null;
         }

        class Path
        {
            public Direction direction;
            public Path previous;
            public int output = -1;
            public int length = 0;

            public Path(Direction direction, Path previous)
            {
                this.direction = direction;
                this.previous = previous;
                length = previous != null ? previous.length + 1 : 1;
            }

            public Tuple<Tuple<int, int>, int> apply(IntcodeComputer computer, Map map)
            {
                var position = Tuple.Create(0, 0);

                var directions = new LinkedList<Direction>();
                for (var current = this; current != null; current = current.previous)
                    directions.AddFirst(current.direction);

                for (var current = directions.First; current != null; current = current.Next)
                {
                    output = (int) computer.execute(new int[] { (int)current.Value })[0];
                    if (output == 0)
                    {
                        map.set(move(position, current.Value), '#');
                    }
                    else if (output == 1)
                    {
                        position = move(position, current.Value);

                        map.set(position, '.');
                    }
                }

                return Tuple.Create(position, output);
            }

            public void revert(IntcodeComputer computer)
            {
                var current = this;
                if (output == 0)
                    current = current.previous;

                for (; current != null; current = current.previous)
                {
                    var oppositeDirection = getOppositeDirection(current.direction);

                    var outputs = computer.execute(new int[] { (int)oppositeDirection });
                    if (outputs[0] != 1)
                        throw new Exception("" + outputs[0]);
                }
            }
        }

        public static void solve()
        {
            var input = InputLoader.loadAsString("15").Split(",");

            var computer = new IntcodeComputer(input);

            var map = new Map();

            map.set(Tuple.Create(0, 0), 'X');

            var pathes = new LinkedList<Path>();
            pathes.AddFirst(new Path(Direction.North, null));
            pathes.AddFirst(new Path(Direction.East, null));
            pathes.AddFirst(new Path(Direction.South, null));
            pathes.AddFirst(new Path(Direction.West, null));

            var oxygenStation = Tuple.Create(0, 0);
            while (pathes.Count > 0)
            {
                var currentPath = pathes.Last.Value;
                pathes.RemoveLast();

                var state = currentPath.apply(computer, map);

                if (state.Item2 != 0)
                {
                    if (!map.isVisited(state.Item1, Direction.North))
                        pathes.AddFirst(new Path(Direction.North, currentPath));

                    if (!map.isVisited(state.Item1, Direction.East))
                        pathes.AddFirst(new Path(Direction.East, currentPath));

                    if (!map.isVisited(state.Item1, Direction.South))
                        pathes.AddFirst(new Path(Direction.South, currentPath));

                    if (!map.isVisited(state.Item1, Direction.West))
                        pathes.AddFirst(new Path(Direction.West, currentPath));
                }
                
                if (state.Item2 == 2)
                    oxygenStation = state.Item1;

                currentPath.revert(computer);
            }


            var oxygens = new LinkedList<Tuple<int, int>>();
            oxygens.AddFirst(oxygenStation);

            var minutes = 0;
            while (oxygens.Count > 0)
            {
                var newList = new LinkedList<Tuple<int, int>>();
                for (var current = oxygens.First; current != null; current = current.Next)
                {
                    map.set(current.Value, 'O');

                    if (map.get(Tuple.Create(current.Value.Item1 + 1, current.Value.Item2)) == '.') {
                        newList.AddFirst(Tuple.Create(current.Value.Item1 + 1, current.Value.Item2));
                    }

                    if (map.get(Tuple.Create(current.Value.Item1 - 1, current.Value.Item2)) == '.')
                    {
                        newList.AddFirst(Tuple.Create(current.Value.Item1 - 1, current.Value.Item2));
                    }

                    if (map.get(Tuple.Create(current.Value.Item1, current.Value.Item2 + 1)) == '.')
                    {
                        newList.AddFirst(Tuple.Create(current.Value.Item1, current.Value.Item2 + 1));
                    }

                    if (map.get(Tuple.Create(current.Value.Item1, current.Value.Item2 - 1)) == '.')
                    {
                        newList.AddFirst(Tuple.Create(current.Value.Item1, current.Value.Item2 - 1));
                    }
                }

                oxygens = newList;

                minutes++;
            }

            Console.WriteLine(minutes);
        }
    }
}
