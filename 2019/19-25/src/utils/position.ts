export interface Position {
    x: number;
    y: number;
}

export const getPositionKey = (position: Position) => {
    return `${position.x},${position.y}`;
};

export const isPositionsEqual = (a: Position, b: Position): boolean => {
    if (!a || !b) {
        return !a == !b;
    }

    return a.x == b.x && a.y == b.y;
};

export const getManhattanDistance = (a: Position, b: Position): number => {
    return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
};
