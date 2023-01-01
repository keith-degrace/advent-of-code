import { log } from "./log";

type IntCodeState = "Ready" | "Terminated" | "WaitingForInput";

export interface IntCode {
    id: string;
    program: number[];
    state: IntCodeState;
    input: number[];
    output: number[];
    instructionPointer: number;
    inputIndex: number;
    relativeBase: number;
}

export const createIntCode = (id: string, program: number[], input?: number[]): IntCode => {
    return {
        id,
        program,
        state: "Ready",
        input: input ? [...input] : [],
        output: [],
        instructionPointer: 0,
        inputIndex: 0,
        relativeBase: 0,
    };
};

enum OpCode {
    Add = 1,
    Multiply = 2,
    Input = 3,
    Output = 4,
    JumpIfTrue = 5,
    JumpIfFalse = 6,
    LessThan = 7,
    Equals = 8,
    AdjustRelativeBase = 9,
    Terminate = 99,
}

enum ParameterMode {
    Position = 0,
    Immediate = 1,
    Relative = 2,
}

export const runIntCode = (intCode: IntCode): void => {
    if (intCode.state === "Terminated") {
        return;
    }

    intCode.state = "Ready";

    const getValue = (pointer: number): number => {
        return intCode.program[pointer] ?? 0;
    };

    const getParameterMode = (instructionPointer: number, parameter: 1 | 2 | 3): ParameterMode => {
        const fullOpCode = getValue(instructionPointer);
        return Math.floor(fullOpCode / (100 * Math.pow(10, parameter - 1))) % 10;
    };

    const getParameterValue = (instructionPointer: number, parameter: 1 | 2 | 3) => {
        const value = getValue(instructionPointer + parameter);
        const mode = getParameterMode(instructionPointer, parameter);

        switch (mode) {
            case ParameterMode.Position:
                return getValue(value);

            case ParameterMode.Immediate:
                return value;

            case ParameterMode.Relative:
                return getValue(intCode.relativeBase + value);
        }
    };

    const getWriteParameterIndex = (instructionPointer: number, parameter: 1 | 2 | 3): number => {
        const mode = getParameterMode(instructionPointer, parameter);
        const value = getValue(instructionPointer + parameter);
        return mode === ParameterMode.Position ? value : intCode.relativeBase + value;
    };

    while (true) {
        const opcode = intCode.program[intCode.instructionPointer] % 100;

        switch (opcode) {
            case OpCode.Add: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);
                const value2 = getParameterValue(intCode.instructionPointer, 2);
                const storeAt = getWriteParameterIndex(intCode.instructionPointer, 3);

                intCode.program[storeAt] = value1 + value2;

                intCode.instructionPointer += 4;
                break;
            }

            case OpCode.Multiply: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);
                const value2 = getParameterValue(intCode.instructionPointer, 2);
                const storeAt = getWriteParameterIndex(intCode.instructionPointer, 3);

                intCode.program[storeAt] = value1 * value2;

                intCode.instructionPointer += 4;
                break;
            }

            case OpCode.Input: {
                let value;
                if (intCode.inputIndex >= intCode.input.length) {
                    value = -1;
                    intCode.state = "WaitingForInput";
                } else {
                    value = intCode.input[intCode.inputIndex++];
                }

                const storeAt = getWriteParameterIndex(intCode.instructionPointer, 1);

                intCode.program[storeAt] = value;

                intCode.instructionPointer += 2;

                if (intCode.state === "WaitingForInput") {
                    return;
                }

                break;
            }

            case OpCode.Output: {
                const value = getParameterValue(intCode.instructionPointer, 1);

                intCode.output.push(value);

                intCode.instructionPointer += 2;
                break;
            }

            case OpCode.JumpIfTrue: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);
                const value2 = getParameterValue(intCode.instructionPointer, 2);

                if (value1 !== 0) {
                    intCode.instructionPointer = value2;
                } else {
                    intCode.instructionPointer += 3;
                }

                break;
            }

            case OpCode.JumpIfFalse: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);
                const value2 = getParameterValue(intCode.instructionPointer, 2);

                if (value1 === 0) {
                    intCode.instructionPointer = value2;
                } else {
                    intCode.instructionPointer += 3;
                }

                break;
            }

            case OpCode.LessThan: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);
                const value2 = getParameterValue(intCode.instructionPointer, 2);
                const storeAt = getWriteParameterIndex(intCode.instructionPointer, 3);

                intCode.program[storeAt] = value1 < value2 ? 1 : 0;

                intCode.instructionPointer += 4;
                break;
            }

            case OpCode.Equals: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);
                const value2 = getParameterValue(intCode.instructionPointer, 2);
                const storeAt = getWriteParameterIndex(intCode.instructionPointer, 3);

                intCode.program[storeAt] = value1 === value2 ? 1 : 0;

                intCode.instructionPointer += 4;
                break;
            }

            case OpCode.AdjustRelativeBase: {
                const value1 = getParameterValue(intCode.instructionPointer, 1);

                intCode.relativeBase += value1;

                intCode.instructionPointer += 2;
                break;
            }

            case OpCode.Terminate:
                intCode.state = "Terminated";
                return;

            default:
                throw `Invalid opcode: ${opcode}`;
        }
    }
};

export const runSimpleIntCode = (program: number[], input?: number[]): number[] => {
    const intCode = createIntCode("id", program, input);
    runIntCode(intCode);
    return intCode.output;
};
