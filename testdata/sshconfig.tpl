Host {{.Host}}
  HostName {{.HostName}}
  User {{.SshUser}}
  Port {{.SshPort}}
  UserKnownHostsFile /dev/null
  StrictHostKeyChecking no
  PasswordAuthentication no
  IdentityFile {{.IdentityFile}}
  IdentitiesOnly yes
  LogLevel FATAL

