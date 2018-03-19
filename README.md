# This branch for simple rhvh pxe install

## require go version >= 1.9

## Usage

```
go build cmd/simplerhvhprobision/app.go
```

**beaker client must be properly setup, and the machine must be reserved by `bkr whoami`**

Run

```
just for example, use the default case
$> ./app -k "inst.ks=http://10.66.8.175:5060/crawled.rhvh4x_iso/RHVH-4.2-20180305.0/rhvh.ks"
```

will do the pxe autoinstall job

```
# this must run during the post install period to prevent infinte loop
$> ./app -remove hp-z220-16.qe.lab.eng.nay.redhat.com
```
