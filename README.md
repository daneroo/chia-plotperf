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

### Shannon (Space)

- `-b 3400 -r 2` - staggered 90m = 5400

```bash
chia plots create -k 32 -b 3400 -r 2 -t /Volumes/Space/ChiaTemp/ -d /Volumes/ChiaPlots/
sleep 5400; chia plots create -k 32 -b 3400 -r 2 -t /Volumes/Space/ChiaTemp/ -d /Volumes/ChiaPlots/

plot-k32-2021-05-10-01-21
expected : plot-k32-2021-05-10-02-51 (sleep 5400)
```

### DaVinci (DaVinciTM20)

- staggered by 1hr on HDD - /Volumes/DaVinciTM20/ChiaTemp/

```bash
# in screen
chia plots create -k 32 -b 4000 -r 2 -t /Volumes/DaVinciTM20/ChiaTemp/ -d /Volumes/ChiaPlots/
# in second screen
sleep 3600; chia plots create -k 32 -b 4000 -r 2 -t /Volumes/DaVinciTM20/ChiaTemp/ -d /Volumes/ChiaPlots/

plot-k32-2021-05-10-00-01-... start at 00:01 stopped /is junk ?
plot-k32-2021-05-10-00-07-... start at
sleep 3600 
plot-k32-2021-05-10-01-22
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
