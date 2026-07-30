import sys

def solve():
    input_data = sys.stdin.read().split()
    if not input_data:
        return

    INF = float('inf')
    out = []

    t = int(input_data[0])
    idx = 1
    
    for _ in range(t):
        n = int(input_data[idx])
        k = int(input_data[idx + 1])
        s = input_data[idx + 2]
        idx += 3

        flag = [0] * n
        
        for _ in range(k):
            flag1 = -1
            flag2 = -1
            cnt1 = 0
            cnt2 = 0

            for i in range(n):
                if s[i] == '(' and flag[i] == 0:
                    flag1 = i
                    break

            if flag1 == -1:
                break

            for i in range(flag1 + 1, n):
                if flag[i] != 1 and s[i] == ')':
                    cnt1 += 1

            for i in range(n-1, -1, -1):
                if s[i] == ')' and flag[i] == 0:
                    flag2 = i
                    break
            
            if flag2 == -1:
                break

            for i in range(flag2):
                if flag[i] != 1 and s[i] == '(':
                    cnt2 += 1

            if cnt1 <= cnt2:
                flag[flag2] = 1
            else:
                flag[flag1] = 1

        out.append(''.join(map(str, flag)))


if __name__ == '__main__':
    solve()
