import * as fs from "fs";
import * as path from "path";
import { createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";

interface Fold {
    where: number;
    axis: "x" | "y";
}

const load = (): [Grid, Fold[]] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const grid = createGrid();

    let index = 0;
    for (; input[index].length > 0; index++) {
        const x: number = parseInt(input[index].split(",")[0]);
        const y: number = parseInt(input[index].split(",")[1]);

        setPixel(grid, { x, y }, "#");
    }

    index++;

    const folds: Fold[] = [];
    for (; index < input.length; index++) {
        folds.push({
            where: parseInt(input[index].split("=")[1]),
            axis: input[index].split("=")[0].slice(-1) as Fold["axis"],
        });
    }

    return [grid, folds];
};

const foldGrid = (grid: Grid, fold: Fold): Grid => {
    const newGrid = createGrid();

    if (fold.axis === "x") {
        let x1 = fold.where - 1;
        let x2 = fold.where + 1;

        while (x1 >= grid.min.x || x2 <= grid.max.x) {
            for (let y = grid.min.y; y <= grid.max.y; y++) {
                const value1 = getPixel(grid, { x: x1, y });
                const value2 = getPixel(grid, { x: x2, y });

                if (value1 === "#" || value2 === "#") {
                    setPixel(newGrid, { x: x1, y }, "#");
                }
            }

            x1--;
            x2++;
        }
    } else {
        let y1 = fold.where - 1;
        let y2 = fold.where + 1;

        while (y1 >= grid.min.y || y2 <= grid.max.y) {
            for (let x = grid.min.x; x <= grid.max.x; x++) {
                const value1 = getPixel(grid, { x, y: y1 });
                const value2 = getPixel(grid, { x, y: y2 });

                if (value1 === "#" || value2 === "#") {
                    setPixel(newGrid, { x, y: y1 }, "#");
                }
            }

            y1--;
            y2++;
        }
    }

    return newGrid;
};

export const run = () => {
    let [grid, folds] = load();

    for (const fold of folds) {
        grid = foldGrid(grid, fold);
    }

    printGrid(grid, " ");
};
