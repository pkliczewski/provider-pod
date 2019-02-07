# provider-pod

Before the container is released we need to build it using:

```bash
make build
```

or we can pull it from quay.io by running:

```bash
docker pull quay.io/pkliczewski/provider-pod:latest
```

Now when container image is available locally we need to customize the [manifest](manifests/manifest.yml):

We need to customize data in our secret. Each value needs to be base64 encoded

```yaml
  url: {{vmware address}}
  username: {{admin username}}
  password: {{admin password}}
  token: {{api authentication token}}
```

We can customize pods port by changing `SERVER_PORT` and ports section in the service.

next can start the pod by:

```bash
kubectl create -f manifests/manifest.yml
```

We can call the pod by running

```bash
curl localhost:8080/healthcheck
{"result":"OK"}
```

To call protected endpoints like list of VMs we need to use the token from the secret:

```bash
curl -H "Authorization: password" localhost:8080/vms
{"result":["RHEL7_10_NICs","test_v2v"]}
```

To get the details about specific VM we need to:

```bash
curl -H "Authorization: password" localhost:8080/vms/test_v2v
{"result":{"Name":"test_v2v","Template":false,"VmPathName":"[nsimsolo_vmware1] test_v2v/test_v2v.vmx","MemorySizeMB":2048,"CpuReservation":0,"MemoryReservation":0,"NumCpu":2,"NumEthernetCards":1,"NumVirtualDisks":1,"Uuid":"42352ec7-97c8-12a8-a8bc-39a1cd603cd0","InstanceUuid":"503539d1-9968-c8a1-a5a7-e451767a53bd","GuestId":"rhel7_64Guest","GuestFullName":"Red Hat Enterprise Linux 7 (64-bit)","Annotation":"","Product":null,"InstallBootRequired":false,"FtInfo":null,"ManagedBy":null,"TpmPresent":null,"NumVmiopBackings":0}}
```