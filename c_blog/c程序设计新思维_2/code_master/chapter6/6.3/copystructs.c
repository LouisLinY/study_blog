#include <assert.h>
#include <stdio.h>

typedef struct demo_mode {
    int a, b;
    double c, d;
    int *efg;
}demo_s;

/*
需要知道赋值操作是创建了数据的一份复制还是在创建一个新的别名。
该例子，修改了d1.b和d1.c，但是d2中的b、c值并没有发生变化，因此它是
创建了一份复制。但是一个指针的复制仍然指向原来的数据，因此当我们修改了
d1.efg[0]时，这个修改还会影响到d2.efg的复制，如果需要对指针的内容进行复制，就
需要一个结构复制函数。如果我们并不需要追踪任何指针，那么使用复制函数就有点小题大作，只需要
使用等号就可以。 
*/


int main(){
	demo_s d1 = {.b=1, .c=2, .d=3, .efg=(int []){4, 5, 6}};
	demo_s d2 = d1;
	
	d1.b = 14;
	d1.c = 41;
	d1.efg[0] = 7;
	
//	assert(d2.a == 0);
    printf("d2.a = %d\n", d2.a);				//自定数据结构内的数据会自动初始化 
	printf("d2.efg[0] = %d\n", d2.efg[0]);		// 是一个指针的复制仍然指向原来的数据，因此当我们修改了d1.efg[0]时，这个修改还会影响到d2.efg的复制
//	printf(assert(d2.b == 1));
//	printf(assert(d2.c == 2));
//	printf(assert(d2.efg[0] == 7));
    return 0;
}
