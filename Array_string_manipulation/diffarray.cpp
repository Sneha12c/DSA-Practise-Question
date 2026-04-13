// Difference Array Problem 
// You have an array of students where each index represents points each student has. We get 2D array of queries[L, R, X], 
// where all the students in range L to R get X more points. Then we are asked to find final points of any student.
// Approach to do this.
// Create a difference array of N+1 size named diff.
// For each query do diff[L]+=X and diff[R+1]-=X.
// Create a prefix array of N+1 and update each index with pre[i-1]+diff[i]
// Use this prefix array to return value of points at asked index in O(1)
// This way the total complexity reduces it O(N+Q), O(Q) for creating the diff array and O(N) for prefix array. 

// https://leetcode.com/problems/zero-array-transformation-ii/description/

class Solution {
public:
    bool check(vector<int>& nums, vector<vector<int>>& queries, int k) {
        int n = nums.size();
        // We need ONE difference array of size n+1
        vector<long long> diff(n + 1, 0);
        
        // Apply only the first 'k' queries
        for (int i = 0; i < k; i++) {
            int l = queries[i][0];
            int r = queries[i][1];
            int val = queries[i][2];
            diff[l] += val;
            if (r + 1 < n) diff[r + 1] -= val;
        }
        
        long long current_reduction = 0;
        for (int i = 0; i < n; i++) {
            current_reduction += diff[i];
            // If the original value is greater than the total reduction, 
            // it can't be zeroed out.
            if (nums[i] > current_reduction) {
                return false;
            }
        }
        return true;
    }

    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size(), m = queries.size();
        
        int l = 0, r = m, ans = -1;
        
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (check(nums, queries, mid)) {
                ans = mid;
                r = mid - 1; // Try to find a smaller k
            } else {
                l = mid + 1;
            }
        }
        
        return ans;
    }
};
