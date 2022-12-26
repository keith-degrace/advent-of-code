import { getPositionKey, Position } from "./position";
import { log } from "./log";

export interface Grid {
    pixels: Record<string, string>;
    width?: number;
    height?: number;
}

export const createGrid = (width?: number, height?: number): Grid => {
    return {
        pixels: {},
        width,
        height,
    };
};

export const setPixel = (grid: Grid, position: Position, value: string): void => {
    const positionKey = getPositionKey(position);

    if (value === " ") {
        delete grid.pixels[positionKey];
    } else {
        grid.pixels[positionKey] = value;
    }

    grid.width = Math.max(grid.width ?? 0, position.x + 1);
    grid.height = Math.max(grid.height ?? 0, position.y + 1);
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

export const findValueInGrid = (grid: Grid, value: string): Position => {
    for (let y = 0; y < grid.height; y++) {
        for (let x = 0; x < grid.width; x++) {
            if (getPixel(grid, { x, y }) === value) {
                return { x, y };
            }
        }
    }

    return undefined;
};
