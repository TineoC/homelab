# Homelab Hardware Purchase Plan

Status: approved, full HA plan, buying in phases. Last updated 2026-11-22.

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

Decision: **full HA plan** (3 k8s nodes). Prices are snapshots from 2026-11-22. Re-check them before
buying. RAM, M920q and cable prices are estimates that haven't been verified.

| # | Item | Choice | Where to buy | ~USD |
|---|---|---|---|---|
| **Phase A: storage and media** | | | | |
| A1 | Drives | 2x 12 TB enterprise, recertified (2x 16 TB adds $200-420) | ServerPartDeals, eBay (2-yr+ warranty) | 380-600 |
| A2 | NAS | UGREEN DXP4800 Plus, running TrueNAS | UGREEN refurb, eBay open box, Micro Center open box | 550-620 |
| A3 | App SSD | 1 TB NVMe, new | Micro Center, Amazon | 70 |
| A4 | NAS RAM | +8 GB DDR5 SODIMM | Micro Center | 40-50 |
| A5 | UPS | CyberPower CP1500PFCLCD, refurb | Woot, Verdi (new battery) | 180-220 |
| | **Phase A subtotal** | | | **1,220-1,560** |
| **Phase B: network** | | | | |
| B1 | Switch | MikroTik CRS310-8G+2S+IN, used | eBay | 165 |
| B2 | Firewall | Lenovo M720q (i5-8500T, 8 GB) running OPNsense, router-on-a-stick | eBay, Facebook Marketplace | 160-260 |
| B3 | Rack | 15-18U floor-standing 4-post open frame, adjustable depth to ~30", casters (fits the 25"-deep 4U GPU server) | Facebook Marketplace, eBay, Amazon | 80-150 |
| B4 | Cabling | Cat6 patch cables | Micro Center | ~30 |
| B5 | Rack PDU | 1U rack power strip (no daisy-chained power strips) | Amazon, Micro Center | 40-60 |
| | **Phase B subtotal** | | | **475-665** |
| **Phase C: k8s HA** | | | | |
| C1 | HP RAM | 2x 16 GB DDR4 SODIMM (HP 800 G3 to 32 GB) | Micro Center, eBay | 50-60 |
| C2 | k8s node 2 | Lenovo M920q (i5-8500T, 6 cores), 32 GB | eBay, Facebook Marketplace | 200-250 |
| C3 | k8s node 3 | The Phase D GPU server doubles as node 3 | — | 0 |
| | **Phase C subtotal** | | | **250-310** |
| | **Phases A-C** | | | **~1,945-2,535** |
| | **Phase D (GPU/AI node, see below)** | | | **~2,280-2,700** |
| | **Grand total** | | | **~4,225-5,235** |

Already owned: HP EliteDesk 800 G3, ISP router (used as the Wi-Fi access point), TV.

The original $2,000 budget was dropped once local AI became a goal. Phases A + B (~$1,695-2,225) still
cover storage, media and the network lab on their own. Phases C and D can follow as prices allow.

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

## GPU / AI node (Phase D)

Goal: serve local models to the whole cluster for agentic coding (opencode), private chat (Open WebUI),
automation (Postiz, n8n), and practice with k8s GPU operations. Must be **extensible**. The $2,000
budget doesn't apply to this phase.

Why not CPU: on the laptop (i7-1165G7), prompt processing runs at ~12 tokens/s. opencode sends ~19,500
tokens per request, so each reply takes ~27 minutes.

### Design

- One AI node, not a GPU in every node. VRAM doesn't pool across nodes, and splitting a model over
  2.5GbE is slow. Nodes 1-2 stay cheap CPU nodes.
- The AI node is also k8s node 3, so HA keeps 3 nodes.
- NVIDIA GPU Operator plus device-plugin time-slicing lets several pods share one GPU. A taint on the
  node keeps ordinary pods off it.
- Ollama or vLLM runs as a Deployment behind a Service (`ollama.ai.svc`), and is also exposed as
  `ai.home.tineochristopher.com` via Envoy Gateway. Every node and the laptop use it over the network.

### Chosen build: 4U GPU server from used parts + used RTX 3090

This is the most cost-effective option that fits in the rack. It's fast on long agentic prompts
(CUDA), extensible to a 2nd GPU, and doubles as k8s node 3. AM4/DDR4 avoids the worst of the DDR5
shortage.

| Part | Choice | ~USD |
|---|---|---|
| Case | Rosewill RSV-L4000U 4U (16.8" W x 7" H x 25" D; GPU length up to 320 mm with middle fan, 400 mm without) | 210 (verified) |
| Rails | 4U slide rails, or a full-depth rack shelf | 30-50 |
| Motherboard | Used X570 with x8/x8 x16 slots, 3-4 slots apart (room for 2 GPUs) | 120-180 |
| CPU | Used Ryzen 7 5700X | 100-130 |
| RAM | 64 GB DDR4 | 150-250 |
| PSU | 1,000 W 80+ Gold (2x 3090 with power limit ~280 W each) | 140-180 |
| Boot/model disk | 1 TB NVMe | 70 |
| GPU | Used RTX 3090 24GB, at most 320 mm long (Founders Edition: 313 mm) | 1,430-1,630 (verified) |
| **Phase D total** | | **~2,280-2,700** |

Only the case and GPU prices are verified (2026-11-22). Board, CPU, RAM, PSU and rails are estimates.

Later: add a 2nd RTX 3090 (~$1,500) for 48 GB of VRAM and 70B models.

### Alternatives considered

| Option | Memory | Price | Why not chosen |
|---|---|---|---|
| Strix Halo 64 / 128 GB (Ryzen AI Max+ 395) | 64-128 GB shared | ~$2,200 / $3,300-3,800 | Not extensible (soldered memory). Prompt processing ~1,400 tok/s on qwen3-coder 30B, slower than CUDA on long contexts. AMD stack is less common in k8s. |
| NVIDIA DGX Spark / ASUS Ascent GX10 (GB10) | 128 GB shared | ~$6,000 | Best compact option (CUDA, ~3x Strix Halo prefill, 2 units can be linked), but too expensive at current prices. Revisit if it drops. |
| Minisforum MS-01 + RTX 4000 SFF Ada 20GB | 20 GB | ~$2,200-3,000 | Card alone costs $1,333-2,300. Needs a single-slot low-profile card. Tops out at ~30B models. |
| RTX 5060 Ti 16GB in the same 4U server | 16 GB | ~$676 used / ~$800 new + platform | Cheaper start, but 16 GB limits agentic-quality models |
| Ryzen 8845HS mini PCs | Shared RAM | ~$820-1,100 | Integrated GPU too slow for agentic prompts |

### Rack layout (15-18U, floor-standing, 4-post)

```
1U  patch panel
1U  MikroTik CRS310
1U  shelf: Tiny OPNsense | Tiny node 2 (M920q)
1U  shelf: HP EliteDesk (node 1)
1U  PDU
6U  shelf: UGREEN DXP4800 Plus
1U  empty (air gap)
4U  GPU server (node 3) on rails, low in the rack (heaviest)
--  UPS on the floor beside the rack
```

Footprint: about 24" W x 30-36" D, plus ~6" behind for cables and airflow. 18U is ~3 ft tall.
Total weight when loaded: ~100-150 lb. Use the floor, not the wall: wall racks are too shallow for a
25" case, and the home is rented (no drilling).

### Safety requirements

- Motherboard only inside a closed, grounded metal chassis on standoffs. No open-air or mining frames.
- Rails or a shelf rated for 15-20 kg. Don't hang the case from its front ears alone.
- Front-to-back airflow, dust filters on intake fans (clean every 1-3 months), rack back left open.
  The closet needs ventilation: ~600 W of heat under load.
- UPS + rack PDU. No daisy-chained power strips.
- Load: ~600-700 W peak with 1 GPU, ~1,000 W with 2. Fine on a 15 A circuit (~1,800 W) without
  other heavy loads.
- Touch the chassis metal before handling parts. Switch the PSU off before opening the case.

Sources: [RTX 3090 sold prices](https://gpusold.com/gpu/rtx-3090-24),
[RTX 3090 history](https://gpudojo.com/rtx-3090),
[Rosewill RSV-L4000U](https://www.newegg.com/rosewill-rsv-l4000u-black/p/N82E16811147327),
[Strix Halo comparison](https://computingforgeeks.com/ryzen-ai-max-395-mini-pc-comparison/),
[Strix Halo Qwen3-Coder benchmarks](https://github.com/hogeheer499-commits/strix-halo-guide/blob/main/BENCHMARKS.md),
[DGX Spark vs Strix Halo](https://petronellatech.com/blog/dgx-spark-vs-strix-halo-vs-gorgon-halo-local-llm-benchmarks/),
[ASUS GX10 at Micro Center](https://www.microcenter.com/product/703497/asus-ascent-gx10-ai-supercomputer),
[RTX 4000 SFF Ada](https://gpuprix.com/us/gpus/rtx-4000-sff-ada-generation),
[mini PC prices](https://computingforgeeks.com/mini-pc-prices/),
[RTX 5060 Ti 16GB history](https://bestvaluegpu.com/history/new-and-used-rtx-5060-ti-16gb-price-history-and-specs/).

## Open items

- [ ] Decide between 12 TB and 16 TB drives, based on restock prices.
- [ ] Verify RAM and M920q prices.
- [x] Go with the full HA plan (3 k8s nodes).
- [ ] Set up price alerts.
- [ ] Choose backup option.
- [x] GPU/AI node: 4U server with used RTX 3090, in a 15-18U floor rack.
- [ ] Verify GPU server platform prices (board, CPU, DDR4, PSU, rails) and pick a 15-18U rack.
