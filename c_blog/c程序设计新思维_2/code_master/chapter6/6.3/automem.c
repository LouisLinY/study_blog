#include <stdio.h>

/*
    两个函数自动分配两块内存：第一个函数分配了一个结构，第二个函数分配了一个较短的数组。
对于自动分配的内存，在每个函数结束时，它们各自的内存块将被释放。
    指定return x为返回语句的函数会把x的值返回给调用函数。但是这个值必须被复制到调用函数中，
而后者的函数帧将要被销毁。对于结构、数值甚至指针类型，调用函数将得到返回值的一份复制。对于数组，
调用函数将得到指向这个数组的指针，而不是数组中数据的复制。
    因为被返回的指针可能指向一块自动分配的数组数据，而后者在函数退出时将要销毁。返回一个指向
一块可能已经被自动销毁的内存的指针时再糟糕不过的事情了 

3这条语句则是非法的，数组事实上是被当作指针看待，因此在退出时，会创建指向out的指针的一份复制。但是
一旦这块自动分配的内存被销毁，这个指针就指向一块坏数据。 
function returns address of local variable
*/

typedef struct power {
	double base,square, cube;
}powers; 

powers get_power(double in){
	powers out = {
	    .base = in,
	    .square = in * in,
	    .cube = in*in*in,
	};
	return out;
}

int *get_even(int count){
	int out[count];
	int i;
	for (i=0; i<count; i++){
		out[i] = 2*i;
	}
	return out;    // 3 -- function returns address of local variable [-Wreturn-local-addr]
}


int main(){
    powers threes = get_power(3);
    //int *evens = get_even(3);
    printf("threes : %g\t%g\t%g\n", threes.base, threes.cube, threes.square);
    //printf("evens : %i\t%i\t%i\n", evens[0], evens[1], evens[2]);  //evens是合法的指向int类型的指针，但是它所指向的数据已经被释放。 
    return 0;
}
