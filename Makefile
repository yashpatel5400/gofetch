main: main.go player.go background.go
	go build main.go player.go background.go

player: player.go player_tests.go background.go 
	go build player.go player_tests.go background.go 

background: background.go background_tests.go
	go build background.go background_tests.go

clean:
	rm player enemy background keyboard