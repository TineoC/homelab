# Kubernetes Homelab

This project sets up and maintains a Kubernetes homelab environment using Ansible. The configuration includes regular system updates, NTP synchronization, UFW firewall configuration, and K3s installation.

## Project Structure

- `ansible/main.yml`: Main playbook that includes roles for system updates, NTP, UFW, and K3s installation.
- `ansible/ansible.cfg`: Ansible configuration file specifying the inventory and vault password file.
- `roles/apt-maintenance`: Role for ensuring regular system updates.
- `roles/ntp`: Role for configuring NTP synchronization.
- `roles/ufw`: Role for configuring UFW firewall.
- `roles/k3s`: Role for installing and configuring K3s.
- `roles/configure-kubeconfig`: Role for configuring the kubeconfig file.

## Prerequisites

- Ansible installed on the control machine.
- SSH access to the target machines.
- Inventory file configured with the target machines.

## Usage

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/homelab.git
    cd homelab/ansible
    ```

2. **Configure the inventory:**

    Edit the `inventory/inventory.ini` file to include your target machines.

3. **Run the playbook:**

    ```bash
    ansible-playbook main.yml
    ```

## Roles

### apt-maintenance

Ensures regular system updates are applied.

### ntp

Configures NTP synchronization to keep system time accurate.

### ufw

Configures UFW firewall to secure the system.

### k3s

Installs and configures K3s, a lightweight Kubernetes distribution.

### configure-kubeconfig

Configures the kubeconfig file for accessing the Kubernetes cluster.

## Configuration

### ansible.cfg

The `ansible.cfg` file specifies the inventory location and vault password file:

```cfg
[defaults]
inventory=./inventory/inventory.ini
vault_password_file = ~/.ansible_vault_password