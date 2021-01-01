BinariesDirectory = ./bin
WindowsBinariesDirectory = bin
MainDirectory = ./cmd/main
ConfigDirectory = ./configs
ConfigDirectoryForWindows = configs

all: create-windows-bin win-copy-config build run
run-l: run-linux
runl: run-linux
run-u: run-unix
run-linux: create-bin copy-config build linux-run
run-unix: create-bin copy-config-mac build linux-run
run-ps: create-windows-bin win-copy-config build run-direct

create-windows-bin:
	if not exist "$(BinariesDirectory)" mkdir "$(BinariesDirectory)"

create-bin:
	mkdir -p "$(BinariesDirectory)"

ps-create-bin:
	New-Item -ItemType Directory -Force -Path bin

copy-config:
	cp -rfRT "$(ConfigDirectory)" "$(BinariesDirectory)/"

copy-config-mac:
	cp -rf "$(ConfigDirectory)" "$(BinariesDirectory)/"

ps-copy-config:
	COPY-ITEM "$(ConfigDirectory)/*.*" "./bin/" -Force

win-copy-config:
	xcopy "$(ConfigDirectoryForWindows)" "$(WindowsBinariesDirectory)" /e /h /c /y /s

build:
	go build -o "$(BinariesDirectory)" "$(MainDirectory)/main.go"

run:
	cd "$(BinariesDirectory)" && main

run-direct:
	"$(BinariesDirectory)/main"

linux-run:
	cd "$(BinariesDirectory)" && ./main

run-tests:
	cd tests && go test -v
	
cat-ssh:
	cat ~/.ssh/id_rsa.pub

ssh-sample:
	echo "ssh-keygen -t rsa -b 4096 -C 'Your email'"
	
modify-authorized-keys:
	sudo vim ~/.ssh/authorized_keys
	
git-clean-get:
	git reset --hard
	git clean -df
	git status
	git pull