import * as fs from "fs";
import * as path from "path";
import { createIntCode, IntCode, runIntCode } from "../utils/intcode";
import { log } from "../utils/log";

export const run = () => {
    let program = fs
        .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
        .split(",")
        .map((value) => parseInt(value));

    const computers: IntCode[] = [];
    for (let i = 0; i < 50; i++) {
        computers[i] = createIntCode(`Computer ${i}`, { ...program }, [i]);
    }

    while (true) {
        const activeComputers = computers.filter((computer) => computer.state !== "Terminated");
        if (activeComputers.length === 0) {
            break;
        }

        for (const computer of activeComputers) {
            runIntCode(computer);
            if (computer.state === "Terminated") {
                log(`Computer ${computer.id} terminated`);
            }

            if (computer.output.length > 0) {
                if (computer.output.length % 3 !== 0) {
                    throw "Length: " + computer.output.length;
                }

                for (let i = 0; i < computer.output.length / 3; i++) {
                    const targetComputerIndex = computer.output[i * 3];
                    const x = computer.output[i * 3 + 1];
                    const y = computer.output[i * 3 + 2];

                    if (targetComputerIndex === 255) {
                        log(y);
                        return;
                    }

                    const targetComputer = computers[targetComputerIndex];
                    if (!targetComputer) {
                        throw `Computer ${targetComputerIndex} does not exist`;
                    }

                    if (targetComputer.state === "Terminated") {
                        throw "huh?";
                    }

                    targetComputer.input.push(x, y);
                }

                computer.output = [];
            }
        }
    }
};
