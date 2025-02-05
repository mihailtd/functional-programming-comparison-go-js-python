function applyOperation(operation, a, b) {
  return operation(a, b);
}

function add(a, b) {
  return a + b;
}

function subtract(a, b) {
  return a - b;
}

let result = applyOperation(add, 10, 5);
console.log(result); // Output: 15

result = applyOperation(subtract, 10, 5);
console.log(result); // Output: 5