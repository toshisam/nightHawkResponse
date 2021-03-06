/*
 *@package 	nightHawk
 *@file 	module.go
 *@author	roshan maskey <roshanmaskey@gmail.com>
 *@version	0.0.1
 *@updated	2016-06-15
 *
 *@description	nightHawk Response module
 */

package nightHawk


import (
	"math/rand"
	"time"
	"encoding/xml"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func NewSessionDir(szdir int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, szdir)
	for i := range b {
		b[i] = chars[rand.Int63() % int64(len(chars))]
	}
	return string(b)
}

func FixEmptyTimestamp() string {
	return "1970-01-01T01:01:01Z"
}


func (rl *RlAgentStateInspector) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

}


func (rl *RlPersistence)ParseAuditData(computername string, caseinfo CaseInformation, auditinfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseinfo
	rl.AuditType = auditinfo
	xml.Unmarshal(xmlData, &rl)

	// Elastic does not like Timestamp field with "" 
	// Implementing workaround to fix it. Workaround timestamp is 1970-01-01T01:01:01Z
	for i := range rl.PersistenceList {
		if rl.PersistenceList[i].FileModified == "" {
			rl.PersistenceList[i].FileModified = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].FileCreated == "" {
			rl.PersistenceList[i].FileCreated = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].FileAccessed == "" {
			rl.PersistenceList[i].FileAccessed = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].FileChanged == "" {
			rl.PersistenceList[i].FileChanged = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].File.Created == "" {
			rl.PersistenceList[i].File.Created = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].File.JobCreated == "" {
			rl.PersistenceList[i].File.JobCreated = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].File.Modified == "" {
			rl.PersistenceList[i].File.Modified = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].File.Accessed == "" {
			rl.PersistenceList[i].File.Accessed = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].File.Changed == "" {
			rl.PersistenceList[i].File.Changed = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].Registry.JobCreated == "" {
			rl.PersistenceList[i].Registry.JobCreated = FixEmptyTimestamp()
		}

		if rl.PersistenceList[i].Registry.Modified == "" {
			rl.PersistenceList[i].Registry.Modified = FixEmptyTimestamp()
		}

		// Fixing for timeline search in Elastic Search
		rl.PersistenceList[i].TlnTime = rl.PersistenceList[i].FileCreated
		rl.PersistenceList[i].File.TlnTime = rl.PersistenceList[i].File.Created
		rl.PersistenceList[i].Registry.TlnTime = rl.PersistenceList[i].Registry.Modified
	}
}



func (rl *RlService) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)	
}

func (rl *RlPort) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)
}

func (rl *RlUserAccount) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.UserList {
		if rl.UserList[i].LastLogin == "" {
			rl.UserList[i].LastLogin = FixEmptyTimestamp()
		}
		rl.UserList[i].TlnTime = rl.UserList[i].LastLogin
	}
}

func (rl *RlTask) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.TaskList {
		if rl.TaskList[i].CreationDate == "" {
			rl.TaskList[i].CreationDate = FixEmptyTimestamp()
		}

		if rl.TaskList[i].MostRecentRunTime == "" {
			rl.TaskList[i].MostRecentRunTime = FixEmptyTimestamp()
		}

		if rl.TaskList[i].NextRunTime == "" {
			rl.TaskList[i].NextRunTime = FixEmptyTimestamp()
		}

		rl.TaskList[i].TlnTime = rl.TaskList[i].CreationDate
	}
}

func (rl *RlProcessMemory) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.ProcessList {
		if rl.ProcessList[i].StartTime == "" {
			rl.ProcessList[i].StartTime = FixEmptyTimestamp()
		}

		if rl.ProcessList[i].KernelTime == "" {
			rl.ProcessList[i].KernelTime = FixEmptyTimestamp()
		}

		if rl.ProcessList[i].UserTime == "" {
			rl.ProcessList[i].UserTime = FixEmptyTimestamp()
		}

		rl.ProcessList[i].TlnTime = rl.ProcessList[i].StartTime
	}
}

func (rl *RlPrefetch) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.PrefetchList {
		if rl.PrefetchList[i].Created == "" {
			rl.PrefetchList[i].Created = FixEmptyTimestamp()
		}

		if rl.PrefetchList[i].LastRun == "" {
			rl.PrefetchList[i].LastRun = FixEmptyTimestamp()
		}

		rl.PrefetchList[i].TlnTime = rl.PrefetchList[i].Created
	}
}


func (rl *RlRegistryRaw) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.RegistryList {
		if rl.RegistryList[i].Modified == "" {
			rl.RegistryList[i].Modified = FixEmptyTimestamp()
		}

		rl.RegistryList[i].TlnTime = rl.RegistryList[i].Modified
	}
}


func (rl *RlSystemInfo) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)
}

func (rl *RlVolume) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.VolumeList {
		if rl.VolumeList[i].CreationTime == "" {
			rl.VolumeList[i].CreationTime = FixEmptyTimestamp()
		}
	}
}


func (rl *RlUrlHistory) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.UrlHistoryList {
		if rl.UrlHistoryList[i].LastVisitDate == "" {
			rl.UrlHistoryList[i].LastVisitDate = FixEmptyTimestamp()
		}
		rl.UrlHistoryList[i].TlnTime = rl.UrlHistoryList[i].LastVisitDate
	}
}


func (rl *RlFileDownloadHistory) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.FileDownloadHistoryList {
		if rl.FileDownloadHistoryList[i].LastModifiedDate == "" {
			rl.FileDownloadHistoryList[i].LastModifiedDate = FixEmptyTimestamp()
		}

		if rl.FileDownloadHistoryList[i].LastCheckedDate == "" {
			rl.FileDownloadHistoryList[i].LastCheckedDate = FixEmptyTimestamp()
		}

		rl.FileDownloadHistoryList[i].TlnTime = rl.FileDownloadHistoryList[i].LastModifiedDate
	}
}


func (rl *RlNetworkDns) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

}


func (rl *RlNetworkRoute) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i:= range rl.RouteEntryList {
		if rl.RouteEntryList[i].JobCreated == "" {
			rl.RouteEntryList[i].JobCreated = FixEmptyTimestamp()
		}
	}
}


func (rl *RlNetworkArp) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i:= range rl.ArpList {
		if rl.ArpList[i].JobCreated == "" {
			rl.ArpList[i].JobCreated = FixEmptyTimestamp()
		}
	}
}


func (rl *RlApiFile) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.FileList {
		if rl.FileList[i].JobCreated == "" {
			rl.FileList[i].JobCreated = FixEmptyTimestamp()
		}

		if rl.FileList[i].Created == "" {
			rl.FileList[i].Created = FixEmptyTimestamp()
		}

		if rl.FileList[i].Modified == "" {
			rl.FileList[i].Modified = FixEmptyTimestamp()
		}

		if rl.FileList[i].Accessed == "" {
			rl.FileList[i].Accessed = FixEmptyTimestamp()
		}

		if rl.FileList[i].Changed == "" {
			rl.FileList[i].Changed = FixEmptyTimestamp()
		}

		rl.FileList[i].TlnTime = rl.FileList[i].Created
	}

}


func (rl *RlDisk) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)
}




func (rl *RlRawFile) ParseAuditData(computername string, caseInfo CaseInformation, auditInfo RlAuditType, xmlData []byte) {
	rl.ComputerName = computername
	rl.CaseInfo = caseInfo
	rl.AuditType = auditInfo
	xml.Unmarshal(xmlData, &rl)

	for i := range rl.FileList {
		if rl.FileList[i].JobCreated == "" {
			rl.FileList[i].JobCreated = FixEmptyTimestamp()
		}

		if rl.FileList[i].Created == "" {
			rl.FileList[i].Created = FixEmptyTimestamp()
		}

		if rl.FileList[i].Modified == "" {
			rl.FileList[i].Modified = FixEmptyTimestamp()
		}

		if rl.FileList[i].Accessed == "" {
			rl.FileList[i].Accessed = FixEmptyTimestamp()
		}

		if rl.FileList[i].Changed == "" {
			rl.FileList[i].Changed = FixEmptyTimestamp()
		}

		if rl.FileList[i].FilenameCreated == "" {
			rl.FileList[i].FilenameCreated = FixEmptyTimestamp()
		}

		if rl.FileList[i].FilenameModified == "" {
			rl.FileList[i].FilenameModified = FixEmptyTimestamp()
		}

		if rl.FileList[i].FilenameAccessed == "" {
			rl.FileList[i].FilenameAccessed = FixEmptyTimestamp()
		}

		if rl.FileList[i].FilenameChanged == "" {
			rl.FileList[i].FilenameChanged = FixEmptyTimestamp()
		}

		rl.FileList[i].TlnTime = rl.FileList[i].FilenameCreated
	}
}
/* __end_of_w32rawfiles__ */






