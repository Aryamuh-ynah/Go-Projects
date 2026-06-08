sum = 0
s=0
sign = 1
n = int(input("Enter an odd number: "))

for i in range(1, n + 1):
    sum += i

print("The sum of numbers from 1 to", n, "is:", sum)    


for i in range(1, n + 1, 2):
    s += i*i*sign
    sign *= -1

print("The sum of the series 1 - 3^2 + 5^2 - 7^2 + ... is:", s) 


def seriedsum(a, n, sign):
    
    if n < a:
        return 0
    square = a * a
    return sign * square + seriedsum(a+2 , n, -sign)

# n = int(input("Enter an odd number: "))
result = seriedsum(1,n, 1)
print("The sum of the series 1 - 3^2 + 5^2 - 7^2 + ... is:", result)


#for n number of terms

def seriedsum(a, n, sign):
    
    if n == 0:
        return 0
    square = a * a
    return sign * square + seriedsum(a+2 , n-1, -sign)


n = int(input("Enter the number of terms: "))
result = seriedsum(1,n, 1)
print("The sum of the series 1 - 3^2 + 5^2 - 7^2 + ... is:", result)



c =0
for i in range(n):
    for j in range(n):
        c += 1
print("The number of iterations is:", c)