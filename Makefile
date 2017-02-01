all: player background
	make player background

player: player.go player_test.go
	go build player.go player_test.go

enemy: enemy.go enemy_test.go
	go build enemy.go enemy_test.go

background: background.go background_test.go
	go build background.go background_test.go

keyboard: keyboard.go keyboard_test.go
	go build keyboard.go keyboard_test.go

clean:
	rm player background