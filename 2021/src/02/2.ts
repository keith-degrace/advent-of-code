import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    let values = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    let x = 0;
    let y = 0;
    let aim = 0;

    for (const instruction of values) {
        let action = instruction.split(" ")[0];
        let value = parseInt(instruction.split(" ")[1]);

        switch (action) {
            case "forward":
                x += value;
                y += value * aim;
                break;

            case "down":
                aim += value;
                break;

            case "up":
                aim -= value;
                break;
        }
    }

    log(x * y);
};
