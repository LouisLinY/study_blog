#include <stdio.h>

/*
    ���������Զ����������ڴ棺��һ������������һ���ṹ���ڶ�������������һ���϶̵����顣
�����Զ�������ڴ棬��ÿ����������ʱ�����Ǹ��Ե��ڴ�齫���ͷš�
    ָ��return xΪ�������ĺ������x��ֵ���ظ����ú������������ֵ���뱻���Ƶ����ú����У�
�����ߵĺ���֡��Ҫ�����١����ڽṹ����ֵ����ָ�����ͣ����ú������õ�����ֵ��һ�ݸ��ơ��������飬
���ú������õ�ָ����������ָ�룬���������������ݵĸ��ơ�
    ��Ϊ�����ص�ָ�����ָ��һ���Զ�������������ݣ��������ں����˳�ʱ��Ҫ���١�����һ��ָ��
һ������Ѿ����Զ����ٵ��ڴ��ָ��ʱ����ⲻ���������� 

3����������ǷǷ��ģ�������ʵ���Ǳ�����ָ�뿴����������˳�ʱ���ᴴ��ָ��out��ָ���һ�ݸ��ơ�����
һ������Զ�������ڴ汻���٣����ָ���ָ��һ�黵���ݡ� 
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
    //printf("evens : %i\t%i\t%i\n", evens[0], evens[1], evens[2]);  //evens�ǺϷ���ָ��int���͵�ָ�룬��������ָ��������Ѿ����ͷš� 
    return 0;
}
