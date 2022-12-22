export const log = (message?: any): void => {
    console.log(`[${new Date().toLocaleString()}] ${message}`);
};
