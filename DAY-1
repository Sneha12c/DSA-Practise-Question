DAY 1 BINARY SORTING STRING

You are given a binary string s
 consisting of only characters 0 and/or 1.

You can perform several operations on this string (possibly zero). There are two types of operations:

    choose two consecutive elements and swap them. In order to perform this operation, you pay (10)^12 coins;
    choose any element from the string and remove it. In order to perform this operation, you pay (10)^12+1 coins.

Your task is to calculate the minimum number of coins required to sort the string s
 in non-decreasing order (i. e. transform so that s1 ≤ s2 ≤⋯≤ sm, where m
 is the length of the string after applying all operations). An empty string is also considered sorted in non-decreasing order.




SOLUTION =>

Intuition =  As we can obwerve that swapping consecutive elemens is beneficial only when "1" is followed by 1 "0" , 
             when it is followed by more than one zeroes than it is better to delete.
             
Algorithm = We will apply a loop from left to right and standing at any particular index , we will count how many zeroes
            ones we have seen. 
            If we require more than 1 swaps then it is beneficial to delete itself .
            Let we have to do n swap
             Cost of n swap = n* (10)^12
             Cost of 1 deletion = (10)^12 + 1
             
       There are two possibilties =
       1) delete all elements which are equal to '1'
       2) delete either all '1' to right or all '0' to right
       
       
       CODE OF SOLUTION in C++
       
#include<iostream>
#include <bits/stdc++.h>
using namespace std;      
       
void solve(){
  string s;
  cin>>s;
  int n = s.length();
  long long int z=0 , o=0;
  for(int i=0; i<n; i++){
    if(s[i]=='0'){
      z++;
    }
    else{
      o++;
    }
   }
long long int po=0 , pz=0;
long long int sw = 1e12 , sd = 1e12+1;
long long int res = sd*(min(o, z));
for(int i=0; i<n; i++){
 if(s[i]=='1'){
  po++;
 }
 else{
  pz++;
 }
 long long int curr=0;
 if(s[i]=='1' && i<n-1 && s[i+1]=='0'){
   long long int curr = (po-1)*sd + min((z-pz-1)*sd , (o-po+1)*sd) + sw;
   res = min(res , curr);
 }
  curr+= po*sd + min((z-pz)*sd , (o-po)*sd);
  res = min(res , curr);
}
 cout<<res<<endl;
}


int main(){
  int t;
  cin>>t;
  while(t--){
    solve();
  }
  return 0;
}


       
       
       
