// Map: Square each number
const numbers = [1, 2, 3, 4, 5];
const squares = numbers.map(n => n * n);
console.log("Squares:", squares);

// Filter: Get even numbers
const evens = numbers.filter(n => n % 2 === 0);
console.log("Evens:", evens);

// Reduce: Sum all numbers
const sum = numbers.reduce((acc, n) => acc + n, 0); // 0 is the initial value
console.log("Sum:", sum);