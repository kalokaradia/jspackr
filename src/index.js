function fibonacci(n) {
	return n <= 1 ? n : fibonacci(n - 1) + fibonacci(n - 2);
}

console.log(fibonacci(10));

function greet(name) {
	console.log(`Hello, ${name}`);
}

greet("Kaloka");
