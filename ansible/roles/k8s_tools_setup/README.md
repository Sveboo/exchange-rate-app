kubernetes_tools_installation
=========

This role is used for installation k8s tools kubelet, kubectl and kubeadm with help [package](#dependencies) and setting up kubelet
The role installs packages on Debian-based systems.

Role Variables
--------------

|Variable      | Default Value                                     | Description                                                                 |
|--------------|---------------------------------------------------|-----------------------------------------------------------------------------|
|`repo_url`    |`"https://pkgs.k8s.io/core:/stable:/v1.30/deb/"`   | The URL of the repository where the k8s packages are hosted.             |
|`repo_key_url`|`"https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key"`| The URL of the GPG key for the repository.                                |
|`key_dest`    |`"/etc/apt/trusted.gpg.d/kubernetes-apt-keyring.gpg"`| The destination path where the GPG key for the repository will be stored.   |
|`packages`    |`["kubelet", "kubectl", "kubeadm"]`                | A list of packages to be installed.               |

## Dependencies

This role depends on the `package` role for managing package installations and repository configurations. The `package` role should be included using the `import_role` module with the following variables:

```yaml
  import_role:
    name: package
  vars:
    repository_url: "{{ repo_url }}"
    repository_key_url: "{{ repo_key_url }}"
    keyring: "{{ key_dest }}"
    install_packages: "{{ packages }}"
```

Ensure that the `repo_url`, `repo_key_url`, `key_dest`, and `packages` variables are properly defined to configure the package installation process within the included `package` role.

Example Playbook
----------------

```yaml
- name: Example
  hosts: hosts
  become: true
  roles:
    role: kubernetes_tools_installation
```

License
-------

MIT

Author Information
------------------

`email`: sveboo3348@gmail.com
