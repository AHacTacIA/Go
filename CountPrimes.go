func countPrimes(n int) int {
    var count int = 0
    var prime bool
    if(n>2){
        count++
    }
    for i:=3;i<n;i+=2{
        prime = true
        for k:=3;(k*k)<i+1;k+=2{
            if i%k==0{
                prime = false
            }
        }
        if prime{
            count++
        }
    }
    return count    
}
