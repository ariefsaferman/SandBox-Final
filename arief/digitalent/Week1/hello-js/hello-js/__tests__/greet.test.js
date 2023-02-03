const greet = require('../src/greet');

describe('#greet', function () {
  it('should return Hello world!', function () {
    const result = greet();

    expect(result).toBe('Hello world!');
  });
});
