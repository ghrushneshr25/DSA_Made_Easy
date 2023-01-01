import java.util.*;

class TowerOfHanoi {
    public int ans;

    public TowerOfHanoi(int n){
        this.ans = 0 ;
        function(n, 'a', 'b', 'c');
        System.out.println(this.ans);
    }

    public void function(int n , char src, char aux, char dest){
        
        if(n==1) { // Base Condition
            this.ans++;
            return;
        }

        function(n-1, src, dest, aux);
        this.ans++;
        function(n-1,aux, src , dest);
    }

    public static void main(String[] args){
        TowerOfHanoi obj1 = new TowerOfHanoi(3);
        TowerOfHanoi obj2 = new TowerOfHanoi(8);
        TowerOfHanoi obj3 = new TowerOfHanoi(18);
    }
}