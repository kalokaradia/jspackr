(() => {
  // tests/fixtures/imports/math.js
  function add(a, b) {
    return a + b;
  }
  function multiply(a, b) {
    return a * b;
  }

  // tests/fixtures/imports/message.js
  var message = "Bundled successfully";

  // tests/fixtures/imports/index.js
  console.log(message);
  console.log("add:", add(10, 20));
  console.log("multiply:", multiply(5, 6));
})();
