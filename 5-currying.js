function curry(fn) {
  return function curried(...args) {
    if (args.length >= fn.length) {
      return fn.apply(this, args);
    } else {
      return (...args2) => curried.apply(this, args.concat(args2));
    }
  };
}

function add(a, b, c) {
  return a + b + c;
}

const result = curry(add)(1)(2)(3);
console.log(result); // Output: 6
