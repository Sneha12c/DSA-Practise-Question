import sys

def solve():
    data = sys.stdin.read().split()
    if not data:
        return data

    out = []
    t = int(data[0])
    idx = 1
    
    for _ in range(t):
        n = int(data[idx])
        idx += 1
        prev = 0
        for _ in range(n):
            p = int(data[idx])
            idx+= 1
            if prev > p:
                prev += p
            else:
                prev = p

        out.append(prev)

    for v in out:
        print(v)
    

if __name__ == '__main__':
    solve()

