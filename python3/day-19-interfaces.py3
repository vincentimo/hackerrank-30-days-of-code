# https://www.hackerrank.com/challenges/30-interfaces

class AdvancedArithmetic(object):
  def divisorSum(n):
    raise NotImplementedError

class Calculator(AdvancedArithmetic):
  def divisorSum(self, n):
    divisor_list = [i + 1 for i in range(n) if n % (i + 1) == 0]
    return sum(divisor_list)


n = int(input())
my_calculator = Calculator()
s = my_calculator.divisorSum(n)
print("I implemented: " + type(my_calculator).__bases__[0].__name__)
print(s)