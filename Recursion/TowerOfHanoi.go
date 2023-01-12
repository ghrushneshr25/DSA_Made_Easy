package Recursion

var steps int = 0
func TowerOfHanoi(n int, from , to , aux string){
	if n==1 {
		steps = steps + 1
		return 
	}
	TowerOfHanoi(n-1, from , aux, to)
	steps = steps + 1
	TowerOfHanoi(n-1, aux, to, from)
}
