export const lcm = (number1: number, number2: number) => {
    let min = Math.min(number1, number2);

    while (true) {
        if (min % number1 == 0 && min % number2 == 0) {
            break;
        }

        min++;
    }

    return min;
};
