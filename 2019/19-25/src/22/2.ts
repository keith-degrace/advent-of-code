import bigInt from "big-integer";
import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

/**
 * Note to future self. You're not this smart. Sadly, this is the first and only
 * puzzle so far that required knowledge (of math) that's way above your head.
 *
 * The math logic here was snatched from someone else's solution and implemented in Typescript.
 */
export const run = () => {
    const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const deckSize = 119315717514047;
    const times = 101741582076661;

    let incMultiplier = bigInt(1);
    let offsetDiff = bigInt(0);

    for (const instruction of input) {
        if (instruction === "deal into new stack") {
            incMultiplier = incMultiplier.mod(deckSize).times(-1);
            offsetDiff = offsetDiff.add(incMultiplier).mod(deckSize);
        } else if (instruction.startsWith("cut")) {
            const cutSize = bigInt(instruction.split(" ")[1]);
            offsetDiff = offsetDiff.add(cutSize.multiply(incMultiplier)).mod(deckSize);
        } else if (instruction.startsWith("deal with")) {
            const increment = bigInt(instruction.split(" ")[3]);
            incMultiplier = incMultiplier.multiply(increment.modInv(deckSize)).mod(deckSize);
        }
    }

    const inc = incMultiplier.modPow(times, deckSize);
    const offset = offsetDiff.multiply(bigInt(1).minus(inc)).multiply(bigInt(1).minus(incMultiplier).mod(deckSize).modInv(deckSize)).mod(deckSize);

    log(offset.add(inc.multiply(2020)).mod(deckSize));
};
