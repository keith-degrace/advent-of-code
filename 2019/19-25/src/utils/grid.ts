import { getPositionKey, Position } from "./position";
import { log } from "./log";

export interface Grid<T = string> {
    pixels: Record<string, T>;
    width?: number;
    height?: number;
    min?: Position;
    max?: Position;
}

export const createGrid = <T = string>(): Grid<T> => {
    return {
        pixels: {},
        min: { x: Infinity, y: Infinity },
        max: { x: -Infinity, y: -Infinity },
    };
};

export const setPixel = <T>(grid: Grid<T>, position: Position, value: T): void => {
    const positionKey = getPositionKey(position);

    if (value === " ") {
        delete grid.pixels[positionKey];
    } else {
        grid.pixels[positionKey] = value;
    }

    grid.min.x = Math.min(grid.min.x, position.x);
    grid.max.x = Math.max(grid.max.x, position.x);
    grid.min.y = Math.min(grid.min.y, position.y);
    grid.max.y = Math.max(grid.max.y, position.y);

    grid.width = grid.max.x - grid.min.x;
    grid.height = grid.max.y - grid.min.y;
};

export const getPixel = <T>(grid: Grid<T>, position: Position): T => {
    const positionKey = getPositionKey(position);
    return grid.pixels[positionKey] ?? (" " as T);
};

export const hasPixel = <T>(grid: Grid<T>, position: Position): boolean => {
    const positionKey = getPositionKey(position);
    return grid.pixels[positionKey] !== undefined;
};

export const printGrid = <T>(grid: Grid<T>): void => {
    for (let y = grid.min.y; y <= grid.max.y; y++) {
        let line = "";

        for (let x = grid.min.x; x <= grid.max.x; x++) {
            line += getPixel(grid, { x, y });
        }

        log(line);
    }
};

export const findValueInGrid = <T>(grid: Grid<T>, value: T): Position => {
    for (let y = grid.min.y; y <= grid.max.y; y++) {
        for (let x = grid.min.x; x <= grid.max.x; x++) {
            if (getPixel(grid, { x, y }) === value) {
                return { x, y };
            }
        }
    }

    return undefined;
};
