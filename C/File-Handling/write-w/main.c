#include <stdio.h>


int main(){
		
		FILE *file;

		file = fopen("main.txt","w");
		
		fputs("Hello World\n",file);

		fclose(file);

		return 0;
}
