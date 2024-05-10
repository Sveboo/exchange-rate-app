package
=========

This role is used to install package from default repository or from additional repository mentioned in variables. 
The role installs packages on Debian-based systems.

Role Variables
--------------

| Variable           | Description                                                      | Default Value |
|--------------------|------------------------------------------------------------------|---------------|
| repository_url     | Optional. The URL of the repository to fetch packages from.      | ""            |
| repository_key_url | Optional. The URL of the repository key for package validation. | ""            |
| keyring            | Optional. The destination path for the repository key.           | ""            |
| install_packages   | A list of packages to be installed.                              | []            |

These variables allow for additional flexibility in configuring package installation and repository management within the `package` role.

Example Playbook
----------------

```yaml
- name: Example
  hosts: hosts
  become: true
  roles:
    role: package
    vars:
      install_packages:
        - curl
```

License
-------

MIT

Author Information
------------------

`email`: sveboo3348@gmail.com