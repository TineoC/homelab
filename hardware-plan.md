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
| B3 | Rack | 10" mini rack, 12U (e.g. Tecmojo 12U, 10.23" deep) + shelves | Amazon | ~140 |
| B4 | Cabling | Cat6 patch cables | Micro Center | ~30 |
| B5 | Rack PDU | 10" rack power strip (no daisy-chained power strips) | Amazon | 30-50 |
| B6 | 10G modules | 2x SFP+ to 10GBASE-T RJ45 (MikroTik S+RJ10 or similar), for the NAS and MS-S1 MAX | Amazon, eBay | 80-120 |
| | **Phase B subtotal** | | | **605-765** |
| **Phase C: k8s HA** | | | | |
| C1 | HP RAM | 2x 16 GB DDR4 SODIMM (HP 800 G3 to 32 GB) | Micro Center, eBay | 50-60 |
| C2 | k8s node 2 | Lenovo M920q (i5-8500T, 6 cores), 32 GB | eBay, Facebook Marketplace | 200-250 |
| C3 | k8s node 3 | The Phase D MS-S1 MAX doubles as node 3 | — | 0 |
| | **Phase C subtotal** | | | **250-310** |
| | **Phases A-C** | | | **~2,075-2,635** |
| | **Phase D (Minisforum MS-S1 MAX 128GB / 2TB, see below)** | | | **~3,959** |
| | **Grand total** | | | **~6,034-6,594** |

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
| Media, interim | Jellyfin on the laptop via [local/jellyfin/compose.yaml](local/jellyfin/compose.yaml) until the NAS arrives |
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

### Decision: Strix Halo 128GB (Ryzen AI Max+ 395)

Chosen 2026-11-22. One compact, quiet x86 box with 128 GB of shared memory. It serves models to the
whole cluster and doubles as k8s node 3. That keeps real 3-node HA (HP + M920q + Strix Halo, all x86,
so Proxmox, Talos and x86-only images like Harbor still work).

| Model | Price (verified 2026-11-22) | Notes |
|---|---|---|
| Bosgame M5 128GB | ~$3,324-3,599 | Cheapest when in stock |
| Framework Desktop 128GB | ~$3,449 (often out of stock); mainboard only $3,149 | Modular; repairable |
| NIMO AI Mini PC 128GB / 1TB | $3,514.99 on Amazon (Black Friday, list $4,299.99) | Small brand; few reviews |
| GMKtec EVO-X2 ("X3") 128GB / 2TB | $3,949.99 on Amazon | 2.5GbE only. 1-year warranty. Some dead-on-arrival reviews: test it inside Amazon's 30-day return window. |
| **Minisforum MS-S1 MAX 128GB / 2TB (picked)** | $3,959 on Amazon | 2x 10GbE RJ45, USB4 v2, internal PSU. Best networking and cooling of the list. |

Performance (qwen3-coder 30B, llama.cpp Vulkan): ~1,400 tok/s prompt, ~95-100 tok/s replies.
opencode's ~19,500-token prompt takes about 15-30 s. Prompt speed drops on long contexts.

What it gives up: memory is soldered (no upgrades), and it's slower than a CUDA card on long prompts.
The AMD GPU stack (ROCm/Vulkan) is less common in k8s than the NVIDIA GPU Operator. Use the AMD GPU
device plugin, or run Ollama/llama.cpp with Vulkan on the host.

Picked 2026-11-24: **Minisforum MS-S1 MAX**. Most reliable of the compact Strix Halo boxes that are in stock
(Framework is more repairable but sold out), and its 10GbE reaches the NAS at full speed. Test it in the
first week while the return window is open (BIOS shows 128 GB, `stress-ng` 2 h, `memtest86+`,
`smartctl`, a 30-minute model run).

Growth path: a 2nd Strix Halo, or an NVIDIA box later (DGX Spark / GB10 if it drops to ~$4,000).

### Alternatives considered

| Option | Memory | Price | Why not chosen |
|---|---|---|---|
| 4U server + used RTX 3090 | 24 GB (48 GB with 2 cards) | ~$2,280-2,700 | Fastest and extensible, but loud, hot, needs a deep floor rack, and caps at ~30B models on one card |
| NVIDIA DGX Spark / ASUS GX10 (GB10) | 128 GB shared | ~$6,000 | Faster prompts (CUDA), but ~2x the price today. ARM64. |
| Minisforum MS-01 + RTX 4000 SFF Ada 20GB | 20 GB | ~$2,200-3,000 | Tops out at ~30B models; expensive per GB |
| Mac Studio M4 Max 128GB | 128 GB shared | not verified | macOS; no Talos or CUDA |
| Jetson AGX Thor | 128 GB shared | not verified | Built for robotics; few LLM benchmarks |
| RTX 5060 Ti 16GB / Ryzen 8845HS mini PCs | 16 GB / shared | ~$676-1,100 | Too little memory or too slow for agentic work |

Visual pros/cons: https://claude.ai/artifact/K6GGkXrZ6WjnfA69XM1mNG (private).

### Rack layout (10" mini rack, ~11U)

Everything is compact now, so the deep 19" floor rack isn't needed.

```
1U  patch panel (10")
1U  MikroTik CRS310
1U  shelf: Tiny OPNsense | M920q node 2
1U  shelf: HP EliteDesk node 1
2U  shelf: Strix Halo node 3 (AI)
6U  shelf: UGREEN DXP4800 Plus (check it fits a 10" shelf; otherwise beside the rack)
1U  PDU (10")
--  UPS beside the rack
```

About 10" wide x ~12" deep. It fits on a desk, a shelf or the closet floor. Load: ~250-300 W peak.

### Safety requirements

- Every board inside its own closed case. No open-air frames.
- Shelves rated for the NAS with drives (~6-8 kg). The rack sits on the floor or a sturdy desk.
- Leave the rack's back open for airflow and keep its vents clear. Clean the dust every 1-3 months.
  The closet needs some ventilation: ~250-300 W of heat under load.
- UPS + rack PDU. No daisy-chained power strips.
- The Strix Halo's 330 W power brick goes on a shelf, not hanging from the cable.

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
- [x] GPU/AI node: Strix Halo 128GB in a 10" mini rack.
- [x] Pick the Strix Halo model: Minisforum MS-S1 MAX.
- [ ] Pick a 10" rack. Check the NAS and the MS-S1 MAX fit a 10" shelf.
