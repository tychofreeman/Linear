include $(GOROOT)/src/Make.$(GOARCH)

TARG=linear
GOFILES=\
	matrix.go\


include $(GOROOT)/src/Make.pkg
