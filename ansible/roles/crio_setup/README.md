crio_installation
=========

This role is used for installation cri-o with help [package](#dependencies) and setting up cri-o service.
The role installs packages on Debian-based systems.

Role Variables
--------------

A description of the settable variables for this role should go here, including any variables that are in defaults/main.yml, vars/main.yml, and any variables that can/should be set via parameters to the role. Any variables that are read from other roles and/or the global scope (ie. hostvars, group vars, etc.) should be mentioned here as well.

| Variable       | Default Value                                      | Description                                                                 |
|----------------|----------------------------------------------------|-----------------------------------------------------------------------------|
| `repo_url`     | `"https://pkgs.k8s.io/addons:/cri-o:/stable:/v1.30/deb/"` | The URL of the repository where the CRI-O packages are hosted.             |
| `repo_key_url` | `"https://pkgs.k8s.io/addons:/cri-o:/stable:/v1.30/deb/Release.key"` | The URL of the GPG key for the repository.                                |
| `key_dest`     | `"/etc/apt/trusted.gpg.d/cri-o-apt-keyring.gpg"`     | The destination path where the GPG key for the repository will be stored.   |
| `packages`     | `["software-properties-common", "cri-o"]`             | A list of packages to be installed, including dependencies.                |

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
    role: crio_installation
```

License
-------

MIT

Author Information
------------------

`email`: sveboo3348@gmail.com
