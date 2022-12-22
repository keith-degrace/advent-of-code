import * as fs from "fs";
import * as path from "path";

let input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").trim().split("\n");

const gridWidth: number = input[0].length;
const gridHeight: number = input.length;

const getHeight = (x: number, y: number): number => {
    if (x < 0 || x > gridWidth - 1 || y < 0 || y > gridHeight - 1) {
        return 0;
    }

    return Number.parseInt(input[y][x]);
};

const getVisibleOnRight = (x: number, y: number): number => {
    const height: number = getHeight(x, y);

    let visibleCount: number = 0;

    for (let xx = x + 1; xx < gridWidth; xx++) {
        visibleCount++;

        if (getHeight(xx, y) >= height) {
            break;
        }
    }

    return visibleCount;
};

const getVisibleOnLeft = (x: number, y: number): number => {
    const height: number = getHeight(x, y);

    let visibleCount: number = 0;

    for (let xx = x - 1; xx >= 0; xx--) {
        visibleCount++;

        if (getHeight(xx, y) >= height) {
            break;
        }
    }

    return visibleCount;
};

const getVisibleBelow = (x: number, y: number): number => {
    const height: number = getHeight(x, y);

    let visibleCount: number = 0;

    for (let yy = y + 1; yy < gridHeight; yy++) {
        visibleCount++;

        if (getHeight(x, yy) >= height) {
            break;
        }
    }

    return visibleCount;
};

const getVisibleAbove = (x: number, y: number): number => {
    const height: number = getHeight(x, y);

    let visibleCount: number = 0;

    for (let yy = y - 1; yy >= 0; yy--) {
        visibleCount++;

        if (getHeight(x, yy) >= height) {
            break;
        }
    }

    return visibleCount;
};

const getScore = (x: number, y: number): number => {
    return getVisibleOnRight(x, y) * getVisibleOnLeft(x, y) * getVisibleBelow(x, y) * getVisibleAbove(x, y);
};

const getBestScore = (): number => {
    let bestScore: number = 0;

    for (let x = 0; x < gridWidth; x++) {
        for (let y = 0; y < gridWidth; y++) {
            bestScore = Math.max(bestScore, getScore(x, y));
        }
    }

    return bestScore;
};

console.log(getBestScore());
