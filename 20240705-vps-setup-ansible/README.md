# Some Notes
- The ansible scripts only works in Debian family
linux distribution.

# Installation
Follow [installation guide](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)

# Running Playbook
```shell
ansible-playbook -i inventory.yaml --vault-password-file .vault-password playbooks/07-run-traefik.yaml
```
