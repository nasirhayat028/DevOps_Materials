def sum (a, b):
    return a + b

result = sum(5, 7)
print("Sum of 5 and 7 is:", result)

def multiply (a, b):
    return a * b    
result = multiply(5, 7)
print("Multiplication of 5 and 7 is:", result)  

def greet(name):
    return "Hello, " + name + "!"     

greeting = greet("Nasir")
print(greeting)

def factorial(n):
    if n == 0 or n == 1:
        return 1
    else:
        return n * factorial(n - 1)
    
fact = factorial(5)
print("Factorial of 5 is:", fact)