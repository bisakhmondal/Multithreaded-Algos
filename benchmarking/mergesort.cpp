#include<bits/stdc++.h>
#define loop(i,a,b) for(int i=a;i<b;i++)
#define case int t; cin>>t; while(t--)
using namespace std;

void merge(int arr[], int l, int m, int r) 
{ 
    int i, j, k; 
    int n1 = m - l + 1; 
    int n2 = r - m; 
  
    /* create temp arrays */
    int L[n1], R[n2]; 
  
    /* Copy data to temp arrays L[] and R[] */
    for (i = 0; i < n1; i++) 
        L[i] = arr[l + i]; 
    for (j = 0; j < n2; j++) 
        R[j] = arr[m + 1 + j]; 
  
    /* Merge the temp arrays back into arr[l..r]*/
    i = 0; // Initial index of first subarray 
    j = 0; // Initial index of second subarray 
    k = l; // Initial index of merged subarray 
    while (i < n1 && j < n2) { 
        if (L[i] <= R[j]) { 
            arr[k] = L[i]; 
            i++; 
        } 
        else { 
            arr[k] = R[j]; 
            j++; 
        } 
        k++; 
    } 
  
    /* Copy the remaining elements of L[], if there 
       are any */
    while (i < n1) { 
        arr[k] = L[i]; 
        i++; 
        k++; 
    } 
  
    /* Copy the remaining elements of R[], if there 
       are any */
    while (j < n2) { 
        arr[k] = R[j]; 
        j++; 
        k++; 
    } 
} 
  
/* l is for left index and r is right index of the 
   sub-array of arr to be sorted */
void mergeSort(int arr[], int l, int r) 
{ 
    if (l < r) { 
        // Same as (l+r)/2, but avoids overflow for 
        // large l and h 
        int m = l + (r - l) / 2; 
  
        // Sort first and second halves 
        mergeSort(arr, l, m); 
        mergeSort(arr, m + 1, r); 
  
        merge(arr, l, m, r); 
    } 
}
int len=0;
int* preprocess(string s){
    int ss=0;
    vector<int> v;
    for(auto i: s){
        if(i==' '){
            v.push_back(ss);
            ss=0;
        }else
            ss = ss*10+(i-'0');
    }
    if(ss!=0)
        v.push_back(ss);

    int *arr = new int[v.size()];
    for(int i=0;i<v.size();i++)
    arr[i]=v[i];
    len=v.size();
    return arr;
}
int main()
{
    ifstream ip("../tests/nums1e6.txt",ios::in);
    string s;
    getline(ip,s);
    int *v = preprocess(s);
    cout<<v[len-1]<<endl;
    auto start = std::chrono::high_resolution_clock::now();
    mergeSort(v,0,len-1);
    auto stop = std::chrono::high_resolution_clock::now(); 

    auto duration = std::chrono::duration_cast<std::chrono::milliseconds>(stop-start);
    cout<<duration.count()<<endl;
return 0;
}