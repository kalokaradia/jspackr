(() => {
  // tests/fixtures/nested/utils/numbers.js
  var numbers = [1, 2, 3, 4, 5];
  function sum(values) {
    return values.reduce((total2, value) => total2 + value, 0);
  }

  // tests/fixtures/nested/utils/format.js
  function formatNumber(value) {
    return new Intl.NumberFormat("en-US").format(value);
  }

  // tests/fixtures/nested/index.js
  var total = sum(numbers);
  console.log("Numbers:", numbers);
  console.log("Total:", formatNumber(total));
})();
