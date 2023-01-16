import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

interface Player {
    score: number;
    position: number;
}

const rollDice = (dice: number): number => {
    dice++;
    if (dice > 100) {
        dice = 1;
    }

    return dice;
};

const advance = (position: number, steps: number): number => {
    for (let i = 0; i < steps; i++) {
        position++;
        if (position > 10) {
            position = 1;
        }
    }

    return position;
};

interface Roll {
    value: number;
    combinations: number;
}

const possibleRolls: Roll[] = [
    { value: 3, combinations: 1 },
    { value: 4, combinations: 3 },
    { value: 5, combinations: 6 },
    { value: 6, combinations: 7 },
    { value: 7, combinations: 6 },
    { value: 8, combinations: 3 },
    { value: 9, combinations: 1 },
];

// 1, 1, 1  3
// 1, 2, 1  4
// 1, 3, 1  5
// 1, 1, 2  4
// 1, 2, 2  5
// 1, 3, 2  6
// 1, 1, 3  5
// 1, 2, 3  6
// 1, 3, 3  7

// 2, 1, 1  4
// 2, 2, 1  5
// 2, 3, 1  6
// 2, 1, 2  5
// 2, 2, 2  5
// 2, 3, 2  7
// 2, 1, 3  6
// 2, 2, 3  7
// 2, 3, 3  8

// 3, 1, 1  5
// 3, 2, 1  6
// 3, 3, 1  7
// 3, 1, 2  6
// 3, 2, 2  7
// 3, 3, 2  8
// 3, 1, 3  7
// 3, 2, 3  8
// 3, 3, 3  9

const play = (current: number, players: Player[]): [number, number] => {
    const totalWins: [number, number] = [0, 0];

    for (const roll of possibleRolls) {
        const previousPosition = players[current].position;
        const previousScore = players[current].score;

        players[current].position = advance(players[current].position, roll.value);
        players[current].score += players[current].position;

        if (players[current].score >= 21) {
            totalWins[current] += roll.combinations;
        } else {
            const wins = play(current == 1 ? 0 : 1, players);
            totalWins[0] += wins[0] * roll.combinations;
            totalWins[1] += wins[1] * roll.combinations;
        }

        players[current].position = previousPosition;
        players[current].score = previousScore;
    }

    return totalWins;
};

export const run = () => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const players: Player[] = [
        {
            score: 0,
            position: parseInt(input[0].split(": ")[1]),
        },
        {
            score: 0,
            position: parseInt(input[1].split(": ")[1]),
        },
    ];

    const wins = play(0, players);

    log(wins[0]);
    log(wins[1]);
};
