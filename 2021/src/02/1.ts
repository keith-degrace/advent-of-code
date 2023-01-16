import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

export const run = () => {
    let values = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    let x = 0;
    let y = 0;

    for (const instruction of values) {
        let action = instruction.split(" ")[0];
        let value = parseInt(instruction.split(" ")[1]);

        switch (action) {
            case "forward":
                x += value;
                break;

            case "down":
                y += value;
                break;

            case "up":
                y -= value;
                break;
        }
    }

    log(x * y);
};
