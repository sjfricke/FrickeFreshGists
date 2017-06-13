OPENCL_HEADERS=/usr/local/cuda-8.0/lib64/
LIBOPENCL=/usr/local/cuda-8.0/include/
CC=gcc

main: hello.o
	$(CC) hello.o -o hello -I$(LIBOPENCL) -L$(OPENCL_HEADERS) -lOpenCL

