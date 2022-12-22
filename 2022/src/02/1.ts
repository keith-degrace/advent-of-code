import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\n/);

type Choice = "Rock" | "Paper" | "Scissors";

const getChoice = (value: string): Choice => {
    if (value === "A" || value === "X") {
        return "Rock";
    }

    if (value === "B" || value === "Y") {
        return "Paper";
    }

    if (value === "C" || value === "Z") {
        return "Scissors";
    }

    throw `Unknown choice: ${value}`;
};

const ChoiceValue: { [key in Choice]: number } = {
    Rock: 1,
    Paper: 2,
    Scissors: 3,
};

const isWin = (myChoice: Choice, theirChoice: Choice): boolean => {
    return (myChoice === "Rock" && theirChoice === "Scissors") || (myChoice === "Scissors" && theirChoice === "Paper") || (myChoice === "Paper" && theirChoice === "Rock");
};

const getScore = (myChoice: Choice, theirChoice: Choice): number => {
    let outcome: number;
    if (myChoice === theirChoice) {
        outcome = 3;
    } else {
        if (isWin(myChoice, theirChoice)) {
            outcome = 6;
        } else {
            outcome = 0;
        }
    }

    return ChoiceValue[myChoice] + outcome;
};

let score: number = 0;
for (const round of input) {
    const choices: string[] = round.split(" ");

    const theirChoice: Choice = getChoice(choices[0]);
    const myChoice: Choice = getChoice(choices[1]);

    score += getScore(myChoice, theirChoice);
}

console.log(score);
