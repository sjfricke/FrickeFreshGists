OPENCL_HEADERS=/usr/local/cuda-8.0/lib64/
LIBOPENCL=/usr/local/cuda-8.0/include/
CC=gcc

main:
	gcc vecAdd.c -o vecAdd -I$(LIBOPENCL) -L$(OPENCL_HEADERS) -lOpenCL -lm

