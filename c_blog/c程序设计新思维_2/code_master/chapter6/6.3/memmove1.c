#include <stdio.h>
#include <string.h>

/*
需要对数组进行复制，就回到了内存操纵的语法
复制一个数组需要使用memmove，它已经过时，但是能够完成任务。 
*/

int main(){
	int abc[] = {0, 1, 2}; 
	int *copy1, copy2[3];
	copy1 = abc;
	memmove(copy2, abc, sizeof(int)*3);
	abc[0] = 3;
	printf("acb[0] = %d\n", abc[0]);
	printf("copy1[0] = %d\n", copy1[0]);
	printf("copy2[0] = %d\n", copy2[0]);
	return 0;
}
