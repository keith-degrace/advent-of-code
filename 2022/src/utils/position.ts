export interface Position {
    x: number;
    y: number;
}

export const getPositionKey = (position: Position) => {
    return `${position.x},${position.y}`;
};
