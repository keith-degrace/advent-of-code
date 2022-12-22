import { log } from "./log";

let start: Date;
export const startTimer = () => {
    start = new Date();
};

export const stopTimer = () => {
    const end = new Date();
    log(`Time: ${end.valueOf() - start.valueOf()}ms`);
};
