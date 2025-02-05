import copy

points1 = [{"x": 1, "y": 2}, {"x": 3, "y": 4}]

# Shallow Copy
points2 = copy.copy(points1)
points2.append({"x": 5, "y": 6})  # Modifies only points2
points2[0]["x"] = 10  # Modifies the object in both lists

print(points1)  # Output: [{'x': 10, 'y': 2}, {'x': 3, 'y': 4}] (affected!)
print(points2)  # Output: [{'x': 10, 'y': 2}, {'x': 3, 'y': 4}, {'x': 5, 'y': 6}]

# Deep Copy
points3 = copy.deepcopy(points1)
points3[0]["x"] = 20  # Modifies only points3

print(points1)  # Output: [{'x': 10, 'y': 2}, {'x': 3, 'y': 4}] (unaffected)
print(points3)  # Output: [{'x': 20, 'y': 2}, {'x': 3, 'y': 4}]
