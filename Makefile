all: player background
	make player background

player: player.go player_tests.go
	go build player.go player_tests.go

enemy: enemy.go enemy_tests.go
	go build enemy.go enemy_tests.go

background: background.go background_tests.go
	go build background.go background_tests.go

keyboard: keyboard.go keyboard_tests.go
	go build keyboard.go keyboard_tests.go

clean:
	rm player background