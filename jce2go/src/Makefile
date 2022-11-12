TOPDIR    := ./

#-------------------------------------------------------------------------------

YACC_Y    := jce.y
LEX_L    := jce.l

YACC_H     := jce.tab.hpp
YACC_C     := jce.tab.cpp
YACC_O    := jce.tab.o
LEX_C    := lex.yy.cpp
LEX_O    := lex.yy.o

#-------------------------------------------------------------------------------

VERSION_MAJOR  := 1
VERSION_MINOR  := 6
VERSION_PATCH  := 0
TAF_VERSION    := $(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_PATCH)

#-------------------------------------------------------------------------------

INCLUDE     += -I./
LIB          += -L./util -lutil
LOCAL_OBJ    := ${LEX_O} ${YACC_O} $(LOCAL_OBJ) 
CLEANFILE    +=${YACC_H} ${YACC_C} ${LEX_C}

YACC            := bison
LEX             := flex
CC              = gcc
CXX             = g++
CFLAGS          = -g -fPIC -Wno-deprecated -Wall -DTAF_VERSION=\"$(TAF_VERSION)\"

#TARGET        :=${PARSELIB}

TARGET        :=  jce2go

all : ${YACC_H} ${LEX_C} $(TARGET)

#-------------------------------------------------------------------------------

${YACC_H} ${YACC_C} : ${YACC_Y}
	rm -vf $(YACC_H) $(YACC_C) $(YACC_O)
	${YACC} $^ -o ${YACC_C}


${LEX_C} : ${LEX_L}
	rm -vf ${LEX_C} ${LEX_O}
	${LEX} $^
#maybe....
./util/libutil.a:
	cd ./util && make
lex.yy.o: lex.yy.cpp
	$(CXX) $(INCLUDE) $(CFLAGS) -c $< -o $@

jce.tab.o: jce.tab.cpp
	$(CXX) $(INCLUDE) $(CFLAGS) -c $< -o $@

parse.o: parse.cpp
	$(CXX) $(INCLUDE) $(CFLAGS) -c $< -o $@

element.o: element.cpp
	$(CXX) $(INCLUDE) $(CFLAGS) -c $< -o $@

main.o: main.cpp
	$(CXX) $(INCLUDE) $(CFLAGS) -c $< -o $@

jce2go.o: jce2go.cpp
	$(CXX) $(INCLUDE) $(CFLAGS) -c $< -o $@

jce2go : main.o jce2go.o element.o parse.o jce.tab.o ${LEX_O} ./util/libutil.a
	$(CXX) $(CFLAGS) $^ -o $@ $(LIB)
#-------------------------------------------------------------------------------

install:
	mkdir -p ../bin
	install ${TARGET} ../bin

clean:
	rm -f *.o ${YACC_H} ${YACC_C} ${LEX_C} $(TARGET)
        
#-------------------------------------------------------------------------------
