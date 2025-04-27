<powershell>

Install-WindowsFeature -name AD-Domain-Services -IncludeManagementTools

Install-ADDSForest `
-CreateDnsDelegation:$false `
-DatabasePath "C:\Windows\NTDS" `
-DomainMode "WinThreshold" `
-DomainName "${DomainName}" `
-DomainNetbiosName "${NetBiosName}" `
-ForestMode "WinThreshold" `
-InstallDns:$true `
-LogPath "C:\Windows\NTDS" `
-SafeModeAdministratorPassword "$SafeModePassword" `
-NoRebootOnCompletion:$false `
-SysvolPath "C:\Windows\SYSVOL" `
-Force:$true

Restart-Computer -Force

<powershell>

