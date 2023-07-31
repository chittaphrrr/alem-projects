all: 
	go build main_quadA/quadA.go
	go build main_quadB/quadB.go
	go build main_quadC/quadC.go
	go build main_quadD/quadD.go
	go build main_quadE/quadE.go
	go build main_quadchecker/quadchecker.go
	mv main_quadchecker/quadchecker.go main.go
	rm -rf main_quadA
	rm -rf main_quadB
	rm -rf main_quadC
	rm -rf main_quadD
	rm -rf main_quadE
	rm -rf main_quadchecker
	rm -rf go.mod
	rm -rf Makefile
	rm -rf quadAndChecker.tar.xz

#	tar -cf - * | xz -9 -c - > quadAndChecker.tar.xz
#	tar quadAndChecker foo.tar.xz