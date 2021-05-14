# Chia (XCH) Plotting Performance

```bash
go run main.go

go build
```

## `chia` CLI

### Setup PATH

```bash
export PATH=/Applications/Chia.app/Contents/Resources/app.asar.unpacked/daemon:$PATH
export CHIA_ROOT=~/.chia/mainnet/
```

### Setup New Worker

```bash
# apt install ....

chia init
#Add a private key by mnemonic
chia keys add 
```

## Experiments

### TODO

- single on Chromebook - in progress
- staggered on DaVinci - in progress
- staggered on Shannon  - /Volumes/Space
- staggered on Shannon  - /Volumes/Rocket
- Rocket on Chromebook

### Drives

Performance measured by AJA/BlackMagic/fio

- Drobo CIFS/SMB Mount `/Volumes/ChiaPlots/` : W:70/R:100
- /Volumes/DaVinciTM20/ChiaTemp/ : W:180/R:160
- /Volumes/Rocket/ChiaTemp/ : W:422/R:416 (MacOS Journaled) W:414/R:412 (ExFat)

### Shannon

- 2X - `-b 3400 -r 2` - staggered 90m = 5400 /Volumes/Space/ChiaTemp/ 

```bash
chia plots create -k 32 -b 3400 -r 2 -t /Volumes/Space/ChiaTemp/ -d /Volumes/ChiaPlots/
sleep 5400; chia plots create -k 32 -b 3400 -r 2 -t /Volumes/Space/ChiaTemp/ -d /Volumes/ChiaPlots/

[ 2021-05-10T01:21:00 - 2021-05-11T01:49:20 ]: 24h28m20s (1h01 hour to move to Drobo)
[ 2021-05-10T02:52:00 - 2021-05-11T03:25:05 ]: 24h33m5s
```

- 3X4 `-b 3400 -r 2 -n 4` - staggered 120m,240m = 7200,14400, /Volumes/Rocket/ChiaTemp/

```bash
chia plots create -k 32 -b 3400 -r 2 -n 4 -t /Volumes/Rocket/ChiaTemp -d /Volumes/ChiaPlots/
sleep 7200; chia plots create -k 32 -b 3400 -r 2 -n 4 -t /Volumes/Rocket/ChiaTemp -d /Volumes/ChiaPlots/
sleep 14400; chia plots create -k 32 -b 3400 -r 2 -n 4 -t /Volumes/Rocket/ChiaTemp -d /Volumes/ChiaPlots/

# expecting 03:30, 05:30, 07:30
[ 2021-05-11T03:30:00 - 2021-05-12T05:25:40 ]: 25h55m40s
[ 2021-05-11T05:31:00 - 2021-05-12T18:44:45 ]: 37h13m45s
[ 2021-05-11T07:31:00 - 2021-05-12T16:19:27 ]: 32h48m27s

single on Shannon from GUI -b 3390 -r 4 /Volumes/Space/ChiaTemp
[ 2021-05-12T22:13:00 - 2021-05-13T10:24:43 ]: 12h11m43s
2 sequential on Shannon from GUI -b 3400 -r 4 /Volumes/Space/ChiaTemp
expect 2021-05-13T16:55, 2021-05-14T05:05 done at 2021-05-14T17:15

# plotman could save the trasport time (1hr)
```

### Feynman (/Volumes/Rocket) - as SS-USB 400MB/s

```bash
chia plots create -k 32 -b 3400 -r 4 -t /Volumes/Rocket/ChiaTemp/ -d /Volumes/ChiaPlots/

[ 2021-05-12T22:02:00 - 2021-05-13T06:38:02 ]: 8h36m2s (transport excluded) front USB-C
expect 2021-05-13T16:44 done 2021-05-14T01:20 (unless new usb port is faster!)
Then stagger and perhaps use plotman ... 
```

### DaVinci (DaVinciTM20)

- staggered by 1hr on HDD - /Volumes/DaVinciTM20/ChiaTemp/

```bash
# in screen
chia plots create -k 32 -b 4000 -r 2 -t /Volumes/DaVinciTM20/ChiaTemp/ -d /Volumes/ChiaPlots/
# in second screen
sleep 3600; chia plots create -k 32 -b 4000 -r 2 -t /Volumes/DaVinciTM20/ChiaTemp/ -d /Volumes/ChiaPlots/

[ 2021-05-10T00:07:00 - 2021-05-10T19:48:29 ]: 19h41m29s
[ 2021-05-10T01:22:00 - 2021-05-10T20:22:56 ]: 19h0m56s

```

### Chromebook

```bash
chia plots create -k 32 -b 4000 -r 2 -t ~/ChiaTemp -d ~/ChiaPlots
```

## Example Output

```bash
2021-05-09T00:21:27.883Z - Extracting Chia Plotting Performance From `.plot`s
[ 2021-05-05T15:19:00 - 2021-05-06T13:35:59 ]: 22h16m59s
[ 2021-05-05T15:19:00 - 2021-05-06T13:30:49 ]: 22h11m49s
[ 2021-05-06T18:30:00 - 2021-05-07T04:29:56 ]: 9h59m56s
[ 2021-05-06T23:52:00 - 2021-05-07T09:51:08 ]: 9h59m8s
[ 2021-05-07T04:46:00 - 2021-05-07T12:33:39 ]: 7h47m39s
[ 2021-05-07T15:07:00 - 2021-05-08T04:42:16 ]: 13h35m16s
[ 2021-05-07T15:07:00 - 2021-05-08T04:39:38 ]: 13h32m38s
[ 2021-05-07T20:25:00 - 2021-05-08T14:56:51 ]: 18h31m51s
[ 2021-05-08T05:06:00 - 2021-05-08T12:46:07 ]: 7h40m7s
```

## Plotman

- [Plotman](https://github.com/ericaltendorf/plotman)

THis should be in a venv...

```bash
pip install --force-reinstall git+https://github.com/ericaltendorf/plotman@main
# make a config for the machine: 
# Wrote default plotman.yaml to: /Users/daniel/Library/Application Support/plotman/plotman.yaml
plotman config generate 
plotman config path # shows where the file is located (~/Library/Application Support/plotman/..)
```
