import { getPositionKey, Position } from "./position";
import { log } from "./log";

export interface Grid {
    pixels: Record<string, string>;
    width: number;
    height: number;
}

export const setPixel = (grid: Grid, position: Position, value: string): void => {
    const positionKey = getPositionKey(position);

    if (value === " ") {
        delete grid.pixels[positionKey];
    } else {
        grid.pixels[positionKey] = value;
    }
};

export const getPixel = (grid: Grid, position: Position): string => {
    const positionKey = getPositionKey(position);
    return grid.pixels[positionKey] ?? " ";
};

export const printGrid = (grid: Grid): void => {
    for (let y = 0; y < grid.height; y++) {
        let line = "";

        for (let x = 0; x < grid.width; x++) {
            line += getPixel(grid, { x, y }) ?? " ";
        }

        log(line);
    }
};
