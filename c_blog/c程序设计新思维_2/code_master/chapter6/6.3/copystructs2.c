#include <stdio.h>
/*
对于数组，等号将会复制一个别名，而不是数据本身，创建复制测试，修改原先的数据，并检查复制值。
结构被复制，但是吧一个数组设置为另一个数组只是创建了一个别名
当复制被修改时，原数据也被修改 
*/

int main(){
	int src[] = {0, 1, 2};
	int *copy = src;
	copy[0] = 3;
	printf("src[0] = %d\n", src[0]);
	printf("copy[0] = %d\n", copy[0]);
} 
