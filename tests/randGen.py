import random

random.seed(69)

TOTALGEN = int(1e8)
filename = "nums1e7.txt"
a=[]
for _ in range(TOTALGEN):
    a.append(random.randint(1,10*TOTALGEN))
    
op = " ".join(str(x)  for x in a)

with open(filename,"w") as f:
    f.write(op)