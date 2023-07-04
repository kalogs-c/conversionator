#include "./turbosizenator.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) {
  if (argc < 2) {
    return 0;
  }

  char *fpath = argv[1];
  char *format = argv[2];

  printf("%s - %s\n", fpath, format);

  char *destination = (char *)malloc(100);
  strcpy(destination, fpath);
  strcat(destination, ".");
  strcat(destination, format);

  ResizeImageFromFilePath(fpath, destination, format);
  free(destination);

  return 1;
}
