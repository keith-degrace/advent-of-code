export const log = (message?: any): void => {
    console.log(`[${new Date().toLocaleString()}] ${message}`);
};

let start: Date;
export const startTimer = () => {
    start = new Date();
};

export const stopTimer = () => {
    const end = new Date();
    log(`Time: ${end.valueOf() - start.valueOf()}ms`);
};
