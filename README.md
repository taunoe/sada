# Sada

Ten times ten LED matrix display.

## Hardware

1. Self-made Arduino (ATmega328) board.
2. Self-made 10*10 LED matrix.
3. Self-made driver board with three shift registers.

|       Func | Arduino Pin | ATmega328 | 74HC595 |  name |
|------------|-------------|-----------|---------|-------|
|       Latch|         D12 |        18 |      12 | ST_CP |
|       Clock|         D11 |        17 |      11 | SH_CP |
| Serail Data|         D10 |        16 |      14 |    DS |

## Build

### Init modul

On folder create Go module:

    go mod init

Ad used depencies to `go.mod` file:

    go mod tidy

### Flash

    tinygo flash -target=arduino main.go

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)
