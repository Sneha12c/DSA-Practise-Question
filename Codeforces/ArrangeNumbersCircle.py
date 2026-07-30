import sys

def solve():
    data = sys.stdin.read().split()

    t = int(data[0])
    out = []
    idx = 1

    for _ in range(t):
        p = int(data[idx])
        idx+= 1
       
        cntone = 0
        arr = []
        ans = 0

        for _ in range(p):
            num = int(data[idx])
            idx += 1
            if num == 1:
                cntone += 1
            else:
                arr.append(num)
                ans+= num

        if cntone == 0:
            out.append(ans)
        else:
            if p==cntone:
                out.append(0)
            elif (p-cntone) == 1:
                add = (arr[0]/2)
                ans += min(add, cntone)
                out.append(ans)
            else:
                for v in arr:
                    if cntone <= 0:
                        break
                    if v >=4:
                        ans += 1
                        cntone -= 1
                        more = (v-4)/2
                        ans += min(more ,cntone)
                        cntone -= more
                out.append(ans)

    for v in out:
        print(v)


if __name__ == "__main__":
    solve()
