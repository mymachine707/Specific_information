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
