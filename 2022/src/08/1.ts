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

const isVisibleOnLeft = (x: number, y: number): boolean => {
    const height: number = getHeight(x, y);

    for (let xx = x + 1; xx < gridWidth; xx++) {
        if (getHeight(xx, y) >= height) {
            return false;
        }
    }

    return true;
};

const isVisibleOnRight = (x: number, y: number): boolean => {
    const height: number = getHeight(x, y);

    for (let xx = 0; xx < x; xx++) {
        if (getHeight(xx, y) >= height) {
            return false;
        }
    }

    return true;
};

const isVisibleBelow = (x: number, y: number): boolean => {
    const height: number = getHeight(x, y);

    for (let yy = y + 1; yy < gridHeight; yy++) {
        if (getHeight(x, yy) >= height) {
            return false;
        }
    }

    return true;
};

const isVisibleAbove = (x: number, y: number): boolean => {
    const height: number = getHeight(x, y);

    for (let yy = 0; yy < y; yy++) {
        if (getHeight(x, yy) >= height) {
            return false;
        }
    }

    return true;
};

const isVisible = (x: number, y: number): boolean => {
    return isVisibleOnLeft(x, y) || isVisibleOnRight(x, y) || isVisibleBelow(x, y) || isVisibleAbove(x, y);
};

const getVisibleCount = (): number => {
    let visibleCount: number = 0;

    for (let x = 0; x < gridWidth; x++) {
        for (let y = 0; y < gridWidth; y++) {
            if (isVisible(x, y)) {
                visibleCount++;
            }
        }
    }

    return visibleCount;
};

console.log(getVisibleCount());
