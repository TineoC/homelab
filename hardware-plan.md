# Homelab Hardware Purchase Plan

Status: planning. Last updated 2026-11-22.

## Goals

- Replace Google Drive and Google Photos with self-hosted apps, then cancel the Google One plan.
- Stream movies and shows, including 4K, to the living-room TV with Jellyfin.
- Own every network device and config, to practice networking (DNS, firewall, VLANs, subnetting),
  automation, Kubernetes, blogs and social media automation.
- Keep subscriptions to a minimum.

## Constraints

| Item | Value |
|---|---|
| Budget | Under $2,000, drives included |
| Users | 1 |
| Google data today | ~174 GB (Drive + Photos) |
| Media library | Fewer than 540 titles, 4K wanted |
| TV | TCL with Google TV, Jellyfin app from the Play Store |
| ISP | Verizon Fios, 300 Mbps, no Fios TV (no MoCA needed) |
| Location | Rack goes in the same room as the ISP router. Rented home, no drilling. |
| Offsite backup site | None |
| Existing compute | HP EliteDesk 800 G3 DM (see [HOMELAB.md](HOMELAB.md)) |

## Storage sizing

- Google data: 174 GB x 2 for growth / 0.8 for ZFS headroom = ~0.5 TB.
- Media: about 6-8 TB if half the library is 4K HEVC (WEB-DL, ~20 GB per title) and the rest is 1080p.
  An all-4K-remux library would need 30 TB or more, so keep 4K as compressed HEVC.
- Start with a 4-bay chassis holding 2 drives in a ZFS mirror. Add a second mirror pair when the pool
  reaches 80%.

## Shopping list

Prices are snapshots from 2026-11-22. Re-check them before buying.

| # | Item | Choice | New | Used / refurb |
|---|---|---|---|---|
| 1 | Drives | 2x 12 TB or 2x 16 TB enterprise, recertified | — | 12 TB $190-300 each, 16 TB ~$399 each (WD HC550) |
| 2 | NAS | UGREEN DXP4800 Plus, running TrueNAS | ~$584-657 | Official refurb (2-yr warranty), open box ~$550 |
| 2 | App SSD | 1 TB NVMe | ~$70 | Buy new |
| 3 | UPS | CyberPower CP1500PFCLCD | ~$240 | Refurb $180-220 |
| 4 | Switch | MikroTik CRS310-8G+2S+IN | ~$192 | Used ~$165 |
| 5 | Firewall | Lenovo ThinkCentre M720q (i5-8500T, 8 GB) running OPNsense, router-on-a-stick | — | $160-260 |
| 6 | Rack | 12U open frame | ~$75 | Used from $36 |
| 7 | RAM upgrades | HP 800 G3: 2x 16 GB DDR4 SODIMM. NAS: +8 GB DDR5 SODIMM. | ~$70-110 | — |
| 8 | Second k8s node | Lenovo M920q (i5-8500T, 6 cores), 32 GB | — | $200-250 |
| 9 | Third k8s node (HA) | Lenovo M920q (i5-8500T, 6 cores), 32 GB | — | $200-250 |

Estimated total for items 1-6: ~$1,600-1,800 with 12 TB drives, ~$2,000 with 16 TB drives.
Items 7-8 add ~$270-360. Item 9 adds ~$200-250 and can wait. RAM and M920q prices are estimates that
haven't been verified.

### Rejected options

- **UGREEN DH4300 Plus**: ARM (RK3588), 8 GB of RAM that can't be expanded, UGOS only, no TrueNAS,
  no VMs. Too locked down for the lab goals.
- **Synology**: vendor OS, and its drive-compatibility policy has been restrictive.
- **N100 4x 2.5GbE firewall box**: $250-420 today. At 300 Mbps, a used Tiny with one NIC trunking VLANs
  is enough and costs less.

## Buying order

1. **Drives.** They are the scarcest item. Buy as soon as they restock near target.
2. **NAS + NVMe.** Install TrueNAS (disable the BIOS watchdog first). Deploy Jellyfin, Immich and
   Nextcloud. Migrate the Google data with Takeout, verify it, then cancel Google One.
3. **UPS.** Install NUT so the NAS shuts down cleanly on power loss.
4. **Switch.** Gives the NAS 10GbE or 2.5GbE. Needed before creating VLANs.
5. **OPNsense on the Lenovo Tiny.**
6. **Rack.**
7. **RAM upgrades + second k8s node.** Runs the full stack, but without HA.
8. **Third k8s node.** Gives real HA. Then move data volumes from NFS to Longhorn.

Before trusting data to used gear: SMART long test and `badblocks` on drives, battery self-test on the
UPS, boot test on mini PCs. All inside the return window.

## Target prices (for alerts)

| Item | Alert at or below |
|---|---|
| UGREEN DXP4800 Plus | $560 |
| 12 TB drive (recertified) | $200 each |
| 16 TB drive (recertified) | $260 each |
| MikroTik CRS310 | $185 |
| CyberPower CP1500PFCLCD | $200 |
| Lenovo Tiny (M720q or newer) | $160 |
| 1 TB NVMe | $60 |

Free alerts: camelcamelcamel (Amazon), ServerPartDeals "Notify me", Slickdeals keyword alerts. Later,
self-host changedetection.io on the NAS.

## Network design

### Phase 1: lab behind the ISP router

```
ONT -- ISP router (Wi-Fi for the house and the TV)
          |
       OPNsense (Lenovo Tiny, single NIC, VLAN trunk)
          |
       MikroTik CRS310 -- NAS, HP, k8s nodes
```

- The house internet keeps working while the lab is being built. Double NAT is fine for a lab.
- The TV stays on the ISP Wi-Fi. An OPNsense port forward exposes Jellyfin (`8096`) on the NAS.

### Phase 2: OPNsense at the edge

- Connect the ONT Ethernet straight to OPNsense WAN. Set the ISP router to access-point mode.
- Rolling back means moving one cable.

### VLANs

| VLAN | Purpose | Subnet |
|---|---|---|
| 10 | Management (switch, NAS admin) | 10.0.10.0/24 |
| 20 | Trusted clients | 10.0.20.0/24 |
| 30 | TV / IoT: internet + Jellyfin only | 10.0.30.0/24 |
| 40 | k8s lab (matches [dod-homelab-network-exercise.md](dod-homelab-network-exercise.md)) | 10.0.40.0/24 |

## Software

| Need | App |
|---|---|
| Photos | Immich |
| Files | Nextcloud (or Seafile) |
| Media | Jellyfin (hardware transcoding with Intel QuickSync) |
| Blog | Ghost or Hugo |
| Social media scheduling | Postiz or Mixpost |
| k8s storage | NFS from TrueNAS via `nfs-subdir-external-provisioner` or `democratic-csi` |

Rule: apps holding irreplaceable data (Immich, Nextcloud) run on the NAS. Experiments run on the k8s
cluster.

## Kubernetes capacity

[HOMELAB.md](HOMELAB.md) sizes the full stack at **16 cores, 36 GB RAM and 300 GB+ NVMe**. That stack is
Talos, Cilium, Kyverno, 3 vclusters, Harbor, Tekton, Forgejo, Keycloak, OpenBao, Flux and
Prometheus/Grafana/Loki.

| Stage | Nodes | Cores | RAM | Meets sizing? |
|---|---|---|---|---|
| HP as bought + M720q 8 GB | 2 | 10 | 24 GB | No (RAM short) |
| HP 32 GB + M920q 32 GB | 2 | 10 | 64 GB | RAM yes, cores tight. Drop the staging vcluster if needed. |
| HP 32 GB + 2x M920q 32 GB | 3 | 16 | 96 GB | Yes |

The OPNsense Tiny and the NAS are not counted. The NAS stays busy with ZFS, Immich, Nextcloud and
Jellyfin.

## High availability

| Nodes | etcd quorum | If one node dies |
|---|---|---|
| 1 | 1/1 | Cluster down |
| 2 | 2/2 | Cluster down. Worse than 1 node, since either failure breaks quorum. |
| 3 | 2/3 | Cluster survives |

Requirements for k8s HA:

- 3 nodes, each running both control plane and workloads (Talos).
- Replicated storage: Longhorn or Rook-Ceph on each node's NVMe, 3 replicas. NFS volumes from the NAS
  are not HA.
- 2 or more replicas per app, with a PodDisruptionBudget and anti-affinity: Keycloak, Envoy Gateway,
  Harbor, Forgejo. CloudNativePG with 3 instances for databases.
- If Proxmox is used: 3 nodes, or 2 nodes plus a Raspberry Pi QDevice.

Accepted single points of failure (normal for a homelab):

| Component | Impact | HA option, if ever wanted |
|---|---|---|
| NAS | Media, photos, files and NFS go down | Second NAS with ZFS replication |
| Switch | Everything goes down | Two switches with LACP/MLAG |
| OPNsense | Lab loses internet | Second box with CARP. Good future networking exercise (~$180). |
| UPS / power | Everything goes down | Dual PSUs |
| ISP | No internet | 5G failover |

Data protection comes from ZFS snapshots and backups, not from infrastructure redundancy.

## Backup

RAID is not a backup. With no offsite location, choose one:

- No subscription: a rotated USB drive stored away from home.
- Minimal subscription: restic to Backblaze B2 for photos and documents only (~174 GB, ~$1/month).

Media is not backed up, since it can be re-ripped.

## Open items

- [ ] Decide between 12 TB and 16 TB drives, based on restock prices.
- [ ] Verify RAM and M920q prices.
- [ ] Decide when to buy the third k8s node (HA).
- [ ] Set up price alerts.
- [ ] Choose backup option.
