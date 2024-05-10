Jenkins
=========

This role is used to install and set up jenkins.

Role Variables
--------------

| Variable           | Description                                                      | Default Value |
|--------------------|------------------------------------------------------------------|---------------|
| repository_url     | The URL of the jenkins repository to fetch packages from.        | ""            |
| repository_key_url | The URL of the jenkins repository key for package validation.    | ""            |
| install_packages   | A list of packages to be installed. (java, jenkins)              | []            |

Example Playbook
----------------

```yml
--- 
- name: Playbook for setting up Jenkins
  hosts: jenkins
  become: yes
  - name: Set up Jenkins
    include_role:
      name: jenkins
    vars:
      repository_url: "{{ jenkins_repo }}"
      repository_key_url: "{{ jenkins_repo_key }}"
      install_packages: "{{ jenkins_packages }}"
```

License
-------

MIT

Author Information
------------------

`email`: sveboo3348@gmail.com
