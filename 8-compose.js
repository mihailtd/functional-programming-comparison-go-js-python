const compose = (...funcs) => x => funcs.reduceRight((acc, f) => f(acc), x);

const addOne = x => x + 1;
const square = x => x * x;
const multiplyByTwo = x => x * 2;

const composed = compose(square, multiplyByTwo, addOne); // Order is correct
const result = composed(5); // ((5 + 1) * 2)^2 = 144
console.log("Composed result:", result);