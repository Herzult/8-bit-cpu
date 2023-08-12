# 8 Bit CPU


## Controls

| ID  | Name               | Description                                                |
|-----|--------------------|------------------------------------------------------------|
| HLT | Halt               | Halt the clock                                             |
| MI  | Memory Input       | Update RAM address with bus value                          |
| RI  | RAM Input          | Write RAM current address value with bus value             |
| RO  | RAM Output         | Output RAM current address value to the bus                |
| IO  | Instruction Output | Output address pointer of current instruction to the bus   |
| II  | Instruction Input  | Read current instruction from the bus                      |
| AI  | A Register Input   | Update A register value with bus value                     |
| AO  | A Register Output  | Output A register value to the bus                         |
| ∑O  | Sum Output         | Output the ALU value to the bus                            |
| SU  | Subtract           | Activate ALU subtraction mode                              |
| BI  | B Register Input   | Update B register value with bus value                     |
| OI  | Output Input       | Update output display value with value from the bus        |
| CE  | Counter Enable     | Enable program counter                                     |
| CO  | Counter Output     | Output the current value of the program counter to the bus |
| J   | Jump               | Update program counter to value from the bus               |

## Instructions

### Fetch

Every instruction starts with the same two instructions to fetch instruction and increment program counter:

0. `MI` + `CO`
1. `RO` + `II` + `CE`

### LDA (0001)

2. `MI` + `IO`
3. `RO` + `AI`
4. *Nothing*

### ADD (0010)

2. `MI` + `IO`
3. `RO` + `BI`
4. `AI` + `∑O`

### OUT (1110)

2. `AO` + `OI`
3. *Nothing*
4. *Nothing*

### HLT (0011)

2. `HLT`
3. *Nothing*
4. *Nothing*

## Instruction Register

The instruction register works with two EEPROM chips (AT28C64B), flash the first one on the left with `instruction_register_1.bin` and the second one on the right with `instruction_register_2.bin`.

### Flashing using TommyPROM

You can flash the EEPROM chips using [TommyPROM](https://tomnisbet.github.io/TommyPROM/).

On OSX you will need to install xmodem:

```
$ brew install lrzsz
```

Then connect to the device with screen:

```
screen /dev/tty.usbserial-xxxxxx 115200
```

Use the `W` command and send the file using xmodem:

```
Press CTRL+a
:exec !! lsx -b -X /path/to/instruction_register_x.bin
```