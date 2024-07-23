# install nodes
-  Fedora media writer : https://getfedora.org/en/workstation/download/ (You can get the image there as well)
- ssh-copy-id root@192.168.1.x
- dnf check-update
- sudo dnf upgrade
- Set up NTP - it was set to manual by default in the webui for fedora
- Set NIC name to the same for all nodes like so:
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

- During installation of fedora - make sure partition full disk if we need to, defaults to 15 gb if set to auto
  - We are attempting to stick to 15 gb til we need a new partion, if we need it we create a new one and leave OS to 15

- Add disk 2
  - mkdir /mnt/disk2
  - find UUID - `lsblk -f`
  - vim etc/fstab
    - Append `UUID=xxxx-xxxx /mnt/disk2              ext3    defaults,errors=remount-ro  0 1`
  - reboot - check disk is mounted


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

# Taints and labels
kubectl label nodes k3s-node-4 accelerator=coral
kubectl taint nodes k3s-node-4 coral=true:NoSchedule
kubectl taint nodes k3s-node-4 coral=true:NoSchedule-

kubectl label nodes k3s-node-2 ups=true



# When the initial cert expires after a year...
  - `kubectl create serviceaccount my-admin-sa -n default`
  - `kubectl create clusterrolebinding my-admin-sa-binding --clusterrole=cluster-admin --serviceaccount=default:my-admin-sa`

```
apiVersion: v1
kind: Secret
metadata:
  name: my-admin-sa-token
  namespace: my-namespace
  annotations:
    kubernetes.io/service-account.name: my-admin-sa
type: kubernetes.io/service-account-token
```

  - `kubectl apply -f sa-secret.yaml`
  - `TOKEN=$(kubectl get secret my-admin-sa-token -n default -o jsonpath='{.data.token}' | base64 --decode)`
  - `kubectl get secret $(kubectl get sa default -o jsonpath='{.secrets[0].name}') -o jsonpath='{.data.ca\.crt}' | base64 --decode > ca.crt`
  - `SERVER_URL=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}') # This will probably give you the internal ip, make sure you use the cluster ip`

```
cat <<EOF > ./admin-sa-kubeconfig
apiVersion: v1
kind: Config
clusters:
- cluster:
    certificate-authority: ./ca.crt
    server: $SERVER_URL
  name: my-cluster
contexts:
- context:
    cluster: my-cluster
    user: my-admin-sa
  name: my-admin-context
current-context: my-admin-context
users:
- name: my-admin-sa
  user:
    token: $TOKEN
EOF
```

# Bootstrap secrets
- See home-assistant crd for password ref

# Migrate volumes from Longhorn
- Set backup target in longhorn to `nfs://192.168.1.49:/volume1/longhorn-backup`
- Restore backup volumes
- Set backup-routines
- Set


# Adding nodes

# Replacing nodes
1.  kubectl delete node
2. SSH into node and run
  - `/usr/local/bin/k3s-killall.sh`
  - `/usr/local/bin/k3s-uninstall.sh`
3.  Reboot node
4. Update ansible inventory `ansible/inventory/hosts.yaml`
5. Check that version of k3s_release_version still matches the cluster in `ansible/inventory/group_vars/kubernets/k3s.yaml`
6. Run `task ansible:ping`
7. Run `task ansible:prepare`
8. Run `task ansible:install`

- Based on
  - https://discord.com/channels/673534664354430999/1066771341212127242/1172721839324270622
  - https://discord.com/channels/673534664354430999/1066771341212127242/1078088800577785876
  - https://github.com/onedr0p/flux-cluster-template/discussions/589




# Postgres & TeslaMate
create database x;
create user x with encrypted password '';
grant all privileges on database x to x;