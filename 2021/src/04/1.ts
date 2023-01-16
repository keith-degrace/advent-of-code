import * as fs from "fs";
import * as path from "path";
import { log } from "../utils/log";

interface BoardNumber {
    value: number;
    called: boolean;
}

type Board = BoardNumber[];

const load = (): [number[], Board[]] => {
    const lines = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\n");

    let index = 0;

    const numbers: number[] = lines[index].split(",").map((value) => parseInt(value));
    index += 2;

    const readNextBoardLine = () =>
        lines[index++]
            .split(" ")
            .map((value) => value.trim())
            .filter((value) => value !== "")
            .map((value) => ({ value: parseInt(value), called: false }));

    const boards: Board[] = [];
    const boardCount = Math.floor((lines.length - 1) / 6);
    for (let i = 0; i < boardCount; i++) {
        boards.push([...readNextBoardLine(), ...readNextBoardLine(), ...readNextBoardLine(), ...readNextBoardLine(), ...readNextBoardLine()]);

        index++;
    }

    return [numbers, boards];
};

const isWinning = (board: Board): boolean => {
    // Check vertical lines.
    for (let x = 0; x < 5; x++) {
        let allCalled = true;
        for (let y = 0; y < 5; y++) {
            if (!board[y * 5 + x].called) {
                allCalled = false;
                break;
            }
        }

        if (allCalled) {
            return true;
        }
    }

    // Check horizontal lines.
    for (let y = 0; y < 5; y++) {
        let allCalled = true;
        for (let x = 0; x < 5; x++) {
            if (!board[y * 5 + x].called) {
                allCalled = false;
                break;
            }
        }

        if (allCalled) {
            return true;
        }
    }

    return false;
};

const getSum = (board: Board): number => {
    let calledSum = 0;

    for (let x = 0; x < 5; x++) {
        for (let y = 0; y < 5; y++) {
            if (!board[y * 5 + x].called) {
                calledSum += board[y * 5 + x].value;
            }
        }
    }

    return calledSum;
};

export const run = () => {
    const [numbers, boards] = load();

    for (const number of numbers) {
        for (const board of boards) {
            const boardNumber = board.find((boardNumber) => boardNumber.value == number);
            if (boardNumber) {
                boardNumber.called = true;
            }

            if (isWinning(board)) {
                const score = getSum(board) * number;
                log(score);
                return;
            }
        }
    }

    log("No winner found");
};
