import sys

def solve():
    input = sys.stdin.read
    data = input().split()
     
    if not data:
        return 
    
    t = int(data[0])
    idx = 1

    out = []
    
    for _ in range(t):
        n = int(data[idx])
        idx += 1

        adj = [[] for _ in range(n+1)]
        for i in range(2, n+1):
            p = int(data[idx])
            idx += 1
            adj[p].append(i)

        depth = [0]*(n+1)
        max_d = [0]*(n+1)

        order = []
        stack = [1]
        depth[1] = 0

        while stack:
            u = stack.pop()
            order.append(u)
            for v in adj[u]:
                depth[v] = depth[u] + 1
                stack.append(v)

        for u in reversed(order):
            max_d[u] = depth[u]
            for v in adj[u]:
                if max_d[v] > max_d[u]:
                    max_d[u] = max_d[v]
        
        total = n

        for u in range(1, n+1):
            if len(adj[u]) >= 2:
                max1 = -1
                max2 = -1
                for v in adj[u]:
                    md = max_d[v]
                    if md > max1:
                        max2 = max1
                        max1 = md
                    elif md > max2:
                        max2 = md
                
                total += (max2 - depth[u])

        out.append(str(total))

    print('\n'.join(out))

if __name__ == '__main__':
    solve()
