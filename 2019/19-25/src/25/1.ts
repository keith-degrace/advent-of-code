import * as fs from "fs";
import * as path from "path";
import { runSimpleIntCode } from "../utils/intcode";
import { log } from "../utils/log";

const runProgram = (program: number[], instructions: string[]): string => {
    let input: number[] = [];
    for (const instruction of instructions) {
        for (let i = 0; i < instruction.length; i++) {
            input.push(instruction.codePointAt(i));
        }

        input.push(10);
    }

    return String.fromCodePoint(...runSimpleIntCode({ ...program }, input));
};

export const run = () => {
    let program = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    const getCoin = ["west", "south", "south", "west", "take coin", "east", "north", "north", "east"];
    const getMoltenLava = ["west", "south", "east", "take molten lava", "west", "north", "east"];
    const getInfiniteLoop = ["west", "north", "take infinite loop", "south", "east"];
    const getSpaceHeader = ["east", "south", "south", "take space heater", "north", "north", "west"];
    const getAstrolabe = ["east", "south", "south", "south", "take astrolabe", "north", "north", "north", "west"];
    const getWreath = ["east", "north", "take wreath", "south", "west"];
    const getDehydratedWater = ["east", "north", "north", "west", "take dehydrated water", "east", "south", "south", "west"];
    const getPointer = ["west", "south", "take pointer", "north", "east"];
    const getPrimeNumber = ["west", "south", "south", "take prime number", "north", "north", "east"];
    const goToExit = ["east", "north", "north", "west", "north", "east", "south"];

    log(
        runProgram(program, [
            // ...getCoin,
            ...getPointer,
            // ...getPrimeNumber,
            ...getSpaceHeader,
            // ...getAstrolabe,
            ...getWreath,
            ...getDehydratedWater,
            ...goToExit,
        ])
    );
};
