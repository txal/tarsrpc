
TOPDIR	:= ../..
include ${TOPDIR}/make.rules

#------------------------------------------------

all : $(JCE_H)

ifneq ($(JCE_SRC),)
$(JCE_H) : $(JCE_SRC)
	rm -vf $(JCE_H) $(JCE_CPP)
	$(JCE2CPP_H) --with-taf $(JCE_SRC)
	cp -rf ${JCE_H} ${NOTIFY_H}
endif

#------------------------------------------------

clean:
	rm -vf $(JCE_H) $(JCE_CPP)

cleanall:
	rm -vf $(JCE_H) $(JCE_CPP)

install:
	echo "do nothing"
