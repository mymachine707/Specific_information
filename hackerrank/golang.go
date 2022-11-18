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
// my answer true
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



// https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true
// kabisa yili
func dayOfProgrammer(year int32) string {
    // Write your code here
    var ans string
    var a7 int32 = 215
    var res int32
    var f int32 = 28
    if year<1918 {
        if year%4==0{
            f=29
        }
    }
    if year==1918 {
        return "26.09.1918"
    }
    if year>1918 {
        if year%400==0 || (year%4==0 && year%100!=0){
            f=29
        }
    }
    
    res=a7+f
    res=256-res
    
    ans = strconv.Itoa(int(res))+".09."+strconv.Itoa(int(year))
    return ans
}
