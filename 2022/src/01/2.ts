import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split(/\n/);

let calories: number[] = [];

let currentCalories: number = 0;

for (const line of input) {
    if (line === "") {
        calories.push(currentCalories);
        currentCalories = 0;
    } else {
        currentCalories += Number.parseInt(line);
    }
}

calories.sort((a, b) => b - a);

console.log(calories[0] + calories[1] + calories[2]);
