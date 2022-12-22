import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split(/\n/);

let mostCalories: number = 0;
let currentCalories: number = 0;

for (const line of input) {
    if (line === "") {
        if (currentCalories > mostCalories) {
            mostCalories = currentCalories;
        }

        currentCalories = 0;
    } else {
        currentCalories += Number.parseInt(line);
    }
}

console.log(mostCalories);
