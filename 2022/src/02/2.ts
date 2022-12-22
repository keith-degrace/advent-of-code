import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split(/\n/);

type Choice = "Rock" | "Paper" | "Scissors";

const getChoice = (value: string): Choice => {
    if (value === "A") {
        return "Rock";
    }

    if (value === "B") {
        return "Paper";
    }

    if (value === "C") {
        return "Scissors";
    }

    throw `Unknown choice: ${value}`;
};

type Outcome = "Win" | "Lose" | "Draw";

const getOutcome = (value: string): Outcome => {
    if (value === "X") {
        return "Lose";
    }

    if (value === "Y") {
        return "Draw";
    }

    if (value === "Z") {
        return "Win";
    }

    throw `Unknown outcome: ${value}`;
};

const getMyChoice = (theirChoice: Choice, outcome: Outcome): Choice => {
    if (outcome === "Draw") {
        return theirChoice;
    }

    if (outcome === "Win") {
        switch (theirChoice) {
            case "Rock":
                return "Paper";

            case "Paper":
                return "Scissors";

            case "Scissors":
                return "Rock";
        }
    }

    if (outcome === "Lose") {
        switch (theirChoice) {
            case "Rock":
                return "Scissors";

            case "Paper":
                return "Rock";

            case "Scissors":
                return "Paper";
        }
    }

    throw `Huh?!?`;
};

const ChoiceValue: { [key in Choice]: number } = {
    Rock: 1,
    Paper: 2,
    Scissors: 3,
};

const getScore = (myChoice: Choice, outcome: Outcome): number => {
    switch (outcome) {
        case "Draw":
            return ChoiceValue[myChoice] + 3;

        case "Win":
            return ChoiceValue[myChoice] + 6;
        case "Lose":
            return ChoiceValue[myChoice];
    }

    throw `Huh?!?`;
};

let score: number = 0;
for (const round of input) {
    const choices: string[] = round.split(" ");

    const theirChoice: Choice = getChoice(choices[0]);
    const outcome: Outcome = getOutcome(choices[1]);

    const myChoice: Choice = getMyChoice(theirChoice, outcome);

    score += getScore(myChoice, outcome);
}

console.log(score);
