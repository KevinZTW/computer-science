# Course 1

## step to developing a useful algorithms

1. Model the problem ( think about what step need to take)
2. Find a algorithm to solve it
3. Fast enough? Fit memory?
4. If not, figure out why

## 1.5 Union Find

### Quick Find

So called eager algorithms.

Data structure : - Integer array id [ ] of size N - intepretation : p and q are connected iff (if and only if) they have the same id e.g. [1,3,1, 3, 3] means that 0=2 1=3=4
Find : - To check if they have same id

Union : - To connect p and q, change all entries whose id equals id[p] and id[q]
