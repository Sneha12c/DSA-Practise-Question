import sys

def solve():
    data = sys.stdin.read().split()

    t = int(data[0])
    out = []
    idx = 1

    for _ in range(t):
        p = int(data[idx])
        k = int(data[idx+1])
        idx += 2
        mp = {}
        ans = 0
        for i in range(p):
            mp[data[idx]] = mp.get(data[idx], 0) + 1
            idx += 1
        arr = []
        sum = 0
        for _ , v in mp.items():
            sum += v
            arr.append(v)

        m = len(arr)
        arr.sort()
        prev = 0

        for j, val in enumerate(arr):
            if val==prev:
                continue

            diff = (k-sum)
            maxallow = val-1-prev
            if (diff % (m-j))== 0:
                step = diff//(m-j)
                if step >= -maxallow:
                    ans += 1

            sum -= ((val-prev)*(m-j))
            prev = val


        out.append(ans)

    for i in out:
        print(i)

if __name__ == "__main__" :
    solve()

