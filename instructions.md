# install nodes
-  Fedora media writer : https://getfedora.org/en/workstation/download/ (You can get the image there as well)
- During installation of fedora - make sure partition uses full disks, defaults to 15 gb if set to auto
- dnf check-update
- sudo dnf upgrade
- Set up NTP - it was set to manual by default
- Set NIC name to the same for all nodes - task ansible:prepare
```
sudo vim /etc/systemd/network/25-eth0.link

variant: fcos
version: 1.4.0
storage:
  files:
    - path: /etc/systemd/network/25-eth0.link
      mode: 0644
      contents:
        inline: |
          [Match]
          MACAddress=4c:cc:6a:65:1b:28
          [Link]
          Name=eth0
```
- ssh-copy-id root@192.168.1.x

# Install cluster
- Checkout repo (this one) - based on https://github.com/onedr0p/flux-cluster-template/issues/311
- Change interface in templates:

```
# ./provision/ansible/playbooks/templates/calico-installation.yaml.j2
---
apiVersion: operator.tigera.io/v1
kind: Installation
metadata:
  name: default
spec:
  calicoNetwork:
    nodeAddressAutodetectionV4:
      interface: "eth0"
```
and
```
# ./provision/ansible/playbooks/templates/kube-vip-daemonset.yaml.j2

          env:
            - name: vip_interface
              value: "eth0"

```
and
```
# ./tmpl/cluster/kube-vip-daemonset.yaml

          env:
            - name: vip_interface
              value: "eth0"

```

# install disks on nodes
https://www.techotopia.com/index.php/Adding_a_New_Disk_Drive_to_a_Fedora_Linux_System
lsblk -f

fdisk /dev/sd*

vim /etc/fstab
UUID=<INSERT UUID> /mnt/disk2              ext3    defaults,errors=remount-ro  0 1

### Expand volume
`lvextend -l +100%FREE --resizefs fedora/root`

# Updates from template
Changes from the template was updated with
`git `
`git merge v3.14.0 --allow-unrelated-histories`

# Taints
kubectl taint nodes k3s-node-4 coral:NoExecute
kubectl taint nodes k3s-node-5 zigbee:NoExecute



# Bootstrap secrets
- See home-assistant crd for password ref

# Migrate volumes from Longhorn
- Set backup target in longhorn to `nfs://192.168.1.49:/volume1/longhorn-backup`
- Restore backup volumes


# Adding nodes

# Replacing nodes
- kubectl delete node
- run the nuke Ansible playbook ONLY on that node
- follow https://github.com/onedr0p/flux-cluster-template/discussions/589


# Pi's
https://docs.fedoraproject.org/en-US/quick-docs/raspberry-pi/
https://docs.fedoraproject.org/en-US/quick-docs/raspberry-pi/#resizing-the-main-partition-of-the-microsd-card-after-setup_rpi
