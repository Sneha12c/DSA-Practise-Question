# focus is on 0 if 0 is not present in subarray , mex will be 0 
import sys

def solve():
    data = sys.stdin.read().split()

    t = int(data[0])
    idx = 1
    out = []

    for _ in range(t):
        p = int(data[idx])
        idx += 1
        zero = []
        arr = []
        for i in range(2*p):
            num = int(data[idx])
            idx += 1
            arr.append(num)
            if num == 0:
                zero.append(i)
        ans = 1
        if zero[0] == (zero[1]-1):
            i = zero[0]
            j = zero[1]
            mp1arr = [0]*(p+1)
            mp1arr[0]=1
            while(i>=0 and j<2*p):
                if arr[i] != arr[j]:
                    break
                mp1arr[arr[i]] += 1
                i-=1
                j+=1
            for i, val in enumerate(mp1arr):
                if val==0:
                    ans = max(ans, i)
                    break
            
        else:
            ind = zero[0]
            i = 1
            mp1arr = [0]*(p+1)
            mp1arr[0]=1
            while((ind-i)>=0 and (ind+i)<2*p):
                if arr[ind-i] != arr[ind+i]:
                    break
                mp1arr[arr[ind+i]] += 1
                i+=1

            for i, val in enumerate(mp1arr):
                if val==0:
                    ans = max(ans, i)
                    break
            
            ind = zero[1]
            j = 1
            mp2arr = [0]*(p+1)
            mp2arr[0] = 1
            while((ind-j)>=0 and (ind+j)<2*p):
                if arr[ind-j] != arr[ind+j]:
                    break
                mp2arr[arr[ind+j]] += 1
                j+=1

            for i, val in enumerate(mp2arr):
                if val==0:
                    ans = max(ans , i)
                    break

            k = zero[0]
            l = zero[1]
            mp3arr = [0]*(p+1)
            mp3arr[0] = 1
            ind3 = 1
            while((k+ind3)<=(l-ind3)):
                if arr[k+ind3] != arr[l-ind3]:
                   mp3arr = [0]*(p+1)
                   break
                mp3arr[arr[ind3+k]] += 1
                ind3 += 1

            if mp3arr[0] == 1:
                ind3 = 1
                while((k-ind3)>=0 and (l+ind3)<2*p):
                    if arr[k-ind3] != arr[l+ind3]:
                        break
                    mp3arr[arr[k-ind3]] += 1
                    ind3 += 1

                for i, val in enumerate(mp3arr):
                    if val==0:
                        ans = max(ans , i)
                        break
            
        out.append(ans)

    for i in out:
        print(i)
    
if __name__ == "__main__" :
    solve()
    
    