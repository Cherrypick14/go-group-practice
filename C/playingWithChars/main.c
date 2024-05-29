#include <stdio.h>

int main() {
    char ch;
    char s[100];
    char sen[100];

    // Input character
    scanf("%c", &ch);
    printf("%c\n", ch);

    // Input string
    scanf("%s", s);
    printf("%s\n", s);

    // Input sentence (including spaces)
    scanf(" %[^\n]", sen);
    printf("%s\n", sen);

    return 0;
}
