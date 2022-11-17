//https://www.hackerrank.com/challenges/kangaroo/problem?isFullScreen=true
// Farhod aka javobi
 if (x1 + v1 == x2 + v2) || ( x1<x2 && v1>v2) &&
    ((x2 + v2)%(x1 + v1) == 0 || (x2 - x1)%(v1 - v2) == 0){
        return "YES"
    } 
 return "NO"


// Abdulahad aka javobi Javobi

if v1>v2 && (x2-x1)%(v1-v2)==0{
        return "YES"
    }
    return "NO"

//https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Write your code here
    var ans int32
    
    for i:=0;i<len(ar);i++{
        for d:=0;d<len(ar);d++{
            if i<d && (ar[i]+ar[d])%k==0 {
                ans++
            }
        }
    }
    
    
    return ans
}


