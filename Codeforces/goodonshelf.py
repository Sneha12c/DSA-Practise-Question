import sys

def solve():
    data = sys.stdin.read().split()
    if not data:
        return
    
    INF = float('inf')
    out = []
    
    t = int(data[0])
    idx = 1

    for _ in range(t):
        n = int(data[idx])
        idx += 1
    
        arrmp = {}
        for i in range(n):
            p = int(data[idx])
            idx += 1
            arrmp[p].append(i)

        

        

    

if __name__ == '__main__':
    solve()

