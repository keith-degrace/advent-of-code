import * as fs from "fs";
import * as path from "path";
import { countValueInGrid, createGrid, getPixel, Grid, printGrid, setPixel } from "../utils/grid";
import { log } from "../utils/log";

const load = (): [string, Grid] => {
    const input: string[] = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8").split("\r\n");

    const algorithm = input[0];

    const grid = createGrid();

    let y = 0;
    for (const line of input.slice(2)) {
        let x = 0;
        for (const char of line) {
            setPixel(grid, { x: x++, y }, char);
        }

        y++;
    }

    return [algorithm, grid];
};

const getValue = (grid: Grid, x: number, y: number, algorithm: string): string => {
    let binary = "";
    binary += getPixel(grid, { x: x - 1, y: y - 1 }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x, y: y - 1 }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x + 1, y: y - 1 }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x - 1, y: y }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x, y: y }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x + 1, y: y }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x - 1, y: y + 1 }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x, y: y + 1 }) === "#" ? 1 : 0;
    binary += getPixel(grid, { x: x + 1, y: y + 1 }) === "#" ? 1 : 0;

    const index = parseInt(binary, 2);

    return algorithm[index];
};

const enhance = (grid: Grid, algorithm: string): Grid => {
    const enhancedGrid = createGrid();

    // It's infinite so we calculate just enough outside the output to account for infinity.
    //
    // Note: Found 100 by trial and error.. not really sure why
    //
    for (let x = grid.min.x - 100; x <= grid.max.x + 100; x++) {
        for (let y = grid.min.y - 100; y <= grid.max.y + 100; y++) {
            const value = getValue(grid, x, y, algorithm);
            setPixel(enhancedGrid, { x, y }, value);
        }
    }

    // The grid should only grow by 1 on each side for each iteration, so snap it back.
    enhancedGrid.min.x = grid.min.x - 1;
    enhancedGrid.min.y = grid.min.y - 1;
    enhancedGrid.max.x = grid.max.x + 1;
    enhancedGrid.max.y = grid.max.y + 1;

    return enhancedGrid;
};

export const run = () => {
    let [algorithm, grid] = load();

    for (let i = 0; i < 50; i++) {
        grid = enhance(grid, algorithm);
    }

    const litCount = countValueInGrid(grid, "#");

    log(litCount);
};
