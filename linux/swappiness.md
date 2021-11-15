# Swappiness

## Description
Swappiness is how frequently Linux OS swaps.  
The higher the value is, the more frequently OS swaps.  
For example, swappiness `20` means OS starts swapping when memory usage is
around 80%.  

## Operation

### View
```sh
cat /proc/sys/vm/swappiness
```

### Change swappiness permanetly
add `vm.swappiness = 10` in `/etc/sysctl.conf`

#### For chrome OS
Before doing that above, remove `Read Only` attribute from mount by running the command below.
```sh
sudo /usr/share/vboot/bin/make_dev_ssd.sh --remove_rootfs_verification
```

### Change swappiness temporarily
```sh
sudo sysctl vm.swappiness=10
```
