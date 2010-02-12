include $(GOROOT)/src/Make.$(GOARCH)

TARG=linear
GOFILES=\
	utils.go\
	matrix.go\
	matrix_data.go\
	matrix_operations.go\


include $(GOROOT)/src/Make.pkg
