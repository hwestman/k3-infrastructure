# install nodes
-  Fedora media writer : https://getfedora.org/en/workstation/download/ (You can get the image there as well)
- During installation of fedora - make sure partition uses full disks, defaults to 15 gb if set to auto
- dnf check-update
- sudo dnf upgrade
- Set up NTP - it was set to manual by default
- Set NIC name to the same for all nodes - https://docs.fedoraproject.org/en-US/fedora-coreos/customize-nic/
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
          MACAddress=XX:XX:XX:XX:XX:XX:XX:XX
          [Link]
          Name=eth0
```
- https://github.com/onedr0p/flux-cluster-template/issues/311
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
- ssh-copy-id root@192.168.1.x

# Install cluster
- https://github.com/onedr0p/flux-cluster-template


# Migrate volumes from Longhorn
- Set backup target in longhorn to `nfs://192.168.1.49:/volume1/longhorn-backup`
