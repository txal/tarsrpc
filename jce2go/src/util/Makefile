TOPDIR  := ../
HEADERS := $(wildcard *.h)

.SUFFIXES: .o .c

#CXX = clang
CXX = gcc
CFLAGS  = -m64 -g -O -Wall -D_FILE_OFFSET_BITS=64 -D_LARGEFILE_SOURCE

OBJS = tc_common.o tc_ex.o tc_file.o tc_md5.o tc_option.o
		
all: libutil.a

libutil.a:$(OBJS) $(HEADERS)
	ar -rs $@ $(OBJS)

.cpp.o:
	$(CXX) -I$(TOPDIR) $(CFLAGS) -c $^
.c.o:
	$(CXX) -I$(TOPDIR) $(CFLAGS) -c $^
clean:
	rm -f *.o *.a
