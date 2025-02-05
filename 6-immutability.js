const points1 = [{ x: 1, y: 2 }, { x: 3, y: 4 }];
const points2 = [...points1]; // Creates a shallow copy

points2.push({ x: 5, y: 6 }); // Modifies only points2
points2[0].x = 10; // Modifies the object in both arrays

console.log(points1); // Output: [{ x: 10, y: 2 }, { x: 3, y: 4 }] (affected!)
console.log(points2); // Output: [{ x: 10, y: 2 }, { x: 3, y: 4 }, { x: 5, y: 6 }]