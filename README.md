# Chia (XCH) Plotting Performance

```bash
go run main.go
go build
```

## Syncing logs

```bash
# from dirac
rsync -av --progress davinci:Code/iMetrical/chia-plotperf/logs/ ~/Code/iMetrical/chia-plotperf/logs/
rsync -av --progress shannon:Code/iMetrical/chia-plotperf/logs/ ~/Code/iMetrical/chia-plotperf/logs/
# from shannon
rsync -av --progress davinci:Code/iMetrical/chia-plotperf/logs/ ~/Code/iMetrical/chia-plotperf/logs/
```

## Move plots from drobo to darwin

```txt
# lsblk --fs # for /sd[bc]
sdb                                                                                            
└─sdb1                ext4        ex14-0 9b38d716-5fe4-48b9-9285-6f5696ff29d7      8.6T    27% /mnt/ex14-0
sdc                                                                                            
└─sdc1                ext4        ex14-1 c1977674-78f1-4ea4-9729-f00b61d3da60       12T     0% /mnt/ex14-1
```

```bash

# mount ex14-0, ex14-1, n300-0 by uuid or label...
# see https://linuxhint.com/mount_partition_uuid_label_linux/
sudo mount /dev/sdb1 /mnt/ex14-0
sudo mount /dev/sdc1 /mnt/ex14-1
sudo mount /dev/sdd1 /mnt/n300-0


sudo mount -t cifs -o user=daniel //drobo.imetrical.com/ChiaPlots /mnt/drobo-chiaplots/

rsync -av --progress /mnt/drobo-chiaplots/plot-k32-*.plot /mnt/ex14-0/plots
```

### Export CIFS/Samba from darwin

- <https://ubuntuhandbook.org/index.php/2020/07/share-folder-smb-ubuntu-20-04/>

Will share with no auth @ smb://darwin.imetrical.com/ChiaPlots[01] nd mount from shannon and davinci

```bash
sudo emacs /etc/samba/smb.conf
sudo systemctl restart smbd nmbd
```

```txt
[ChiaPlots0]
   path =  /mnt/ex14-0/plots/
   writable = yes
   guest ok = yes
   guest only = yes
   create mode = 0777
   directory mode = 0777

[ChiaPlots1]
   path =  /mnt/ex14-1/plots/
   writable = yes
   guest ok = yes
   guest only = yes
   create mode = 0777
   directory mode = 0777

[ChiaPlots2]
   path =  /mnt/n300-0/plots/
   writable = yes
   guest ok = yes
   guest only = yes
   create mode = 0777
   directory mode = 0777
```

## `chia` CLI

### Setup PATH on OSX

```bash
export PATH=/Applications/Chia.app/Contents/Resources/app.asar.unpacked/daemon:$PATH
export CHIA_ROOT=~/.chia/mainnet/
```

### Install on darwin (ubunt0)

- Followed instructions on <https://github.com/Chia-Network/chia-blockchain/wiki/INSTALL#ubuntudebian>
  - Update instruction below install section above
- Installed in darwin:~/chia-blockchain

```bash
cd chia-blockchain
. ./activate

chia init
#Add a private key by mnemonic
chia keys add 


chia plots add -d /mnt/ex14-0/plots
chia plots add -d /mnt/ex14-1/plots

chia start farmer

$ chia start farmer
chia_harvester: Already running, use `-r` to restart
chia_farmer: Already running, use `-r` to restart
chia_full_node: Already running, use `-r` to restart
chia_wallet: Already running, use `-r` to restart

chia farm summary
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

### Disk Speed

```bash
# https://blog.purestorage.com/purely-technical/io-plumbing-tests-with-fio/
# us posixaio instead of libaio to work on macos
fio --name=seqwrite --rw=write --direct=1 --ioengine=posixaio --bs=32k --numjobs=4 --size=2G --runtime=600 --group_reporting
fio --name=seqread --rw=read --direct=1 --ioengine=posixaio --bs=32k --numjobs=4 --size=2G --runtime=600 --group_reporting

## hdparm just read
sudo hdparm -Tt /dev/sdXXX

## 4G just Write
dd if=/dev/zero of=./largefile bs=1M count=4096
```

### Shannon (Space)

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

Dell G5 Gaming PC - Abyss Black (Intel Core i7-10700F/1TB SSD/16GB RAM/RTX 2060 Super)

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

[ 2021-05-09T17:10:00 - 2021-05-10T04:09:54 ]: 10h59m54s
```

## Example Output

```bash
2021-05-10T20:34:46.138Z - Extracting Chia Plotting Performance From `.plot`s
[ 2021-05-05T15:19:00 - 2021-05-06T13:35:59 ]: 22h16m59s
[ 2021-05-05T15:19:00 - 2021-05-06T13:30:49 ]: 22h11m49s
[ 2021-05-06T18:30:00 - 2021-05-07T04:29:56 ]: 9h59m56s
[ 2021-05-06T23:52:00 - 2021-05-07T09:51:08 ]: 9h59m8s
[ 2021-05-07T04:46:00 - 2021-05-07T12:33:39 ]: 7h47m39s
[ 2021-05-07T15:07:00 - 2021-05-08T04:42:16 ]: 13h35m16s
[ 2021-05-07T15:07:00 - 2021-05-08T04:39:38 ]: 13h32m38s
[ 2021-05-07T20:25:00 - 2021-05-08T14:56:51 ]: 18h31m51s
[ 2021-05-08T05:06:00 - 2021-05-08T12:46:07 ]: 7h40m7s
[ 2021-05-08T16:27:00 - 2021-05-09T02:11:43 ]: 9h44m43s
[ 2021-05-09T02:38:00 - 2021-05-09T12:10:05 ]: 9h32m5s
[ 2021-05-09T12:10:00 - 2021-05-09T21:38:02 ]: 9h28m2s
[ 2021-05-09T17:10:00 - 2021-05-10T04:09:54 ]: 10h59m54s
```

## Plotman

- [Plotman](https://github.com/ericaltendorf/plotman)

Plotman was installed in the global python 3 space (modified PATH for python3 on Shannon)

This should be in a venv...

```bash
pip install --force-reinstall git+https://github.com/ericaltendorf/plotman@main
# make a config for the machine: 
# Wrote default plotman.yaml to: /Users/daniel/Library/Application Support/plotman/plotman.yaml
plotman config generate 
plotman config path # shows where the file is located (~/Library/Application Support/plotman/..)
```

- Experiment with -2 == -d, might pack more parallell plots into tmp dir
- resize partition
- netdata failed
- [nvme monitoring](https://chiadecentral.com/nuc-small-form-factor-plotting-build/)

## Raid0 Experiment (Euler)

[LVM Striping](https://www.theurbanpenguin.com/striped-lvm-volumes/)

[striped vg](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/logical_volume_manager_administration/stripe_create_ex)

```bash
# gdisk use partition type Linux LVM = 8E00

pvcreate /dev/sdb1 /dev/sdc1
pvcreate /dev/sdb1 /dev/sdc1 /dev/sdd1 /dev/sde1

vgcreate vgstripe /dev/sdb1 /dev/sdc1
vgcreate vgstripe /dev/sdb1 /dev/sdc1 /dev/sdd1 /dev/sde1
# 1TB stripe/2 - Using default stripesize 64.00 KiB
lvcreate  -i 2 -L 1T  -n lvstripe vgstripe
lvconvert --type thin-pool -Zn vgstripe/lvstripe

# undo
lvremove vgstripe # removes lvstripe
vgremove vgstripe
```

[Partition with gdisk](https://www.tecklyfe.com/how-to-partition-format-and-mount-a-disk-on-ubuntu-20-04/)
