#include <stdio.h>
#include <string.h>

/*
��Ҫ��������и��ƣ��ͻص����ڴ���ݵ��﷨
����һ��������Ҫʹ��memmove�����Ѿ���ʱ�������ܹ�������� 
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
