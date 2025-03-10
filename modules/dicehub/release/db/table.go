// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import "time"

type Release struct {
	// ReleaseID Generated by the server when created
	ReleaseID string `json:"releaseId" gorm:"type:varchar(64);primary_key"`
	// ReleaseName Any string, easy for users to identify, the maximum length is 255, required
	ReleaseName string `json:"releaseName" gorm:"index:idx_release_name;not null"`
	// Desc Describe, optional
	Desc string `json:"desc" gorm:"type:text"`
	// Dice When ResourceType is diceyml, store dice.yml context, optional
	Dice string `json:"dice" gorm:"type:text"` // dice.yml
	// Addon When ResourceType is addonyml，store addon.yml context，optional
	Addon string `json:"addon" gorm:"type:text"`
	// Changelog changelog，optional
	Changelog string `json:"changelog" gorm:"type:text"`
	// IsStable not temp if ture, otherwise temp
	IsStable bool `json:"isStable" gorm:"type:tinyint(1)"`
	// IsFormal formal
	IsFormal bool `json:"isFormal" gorm:"type:tinyint(1)"`
	// IsProjectRelease .
	IsProjectRelease bool `json:"IsProjectRelease" gorm:"type:tinyint(1)"`
	// Modes list of release deployment mode
	Modes string `json:"modes" gorm:"type:text"`
	// Labels map type, the maximum length is 1000, optional
	Labels string `json:"labels" gorm:"type:varchar(1000)"`
	// GitBranch
	GitBranch string `json:"gitBranch" gorm:"type:varchar(255)"`
	// Tags
	Tags string `json:"tags" gorm:"type:varchar(100)"`
	// Version store release version, only in the same company, same project and same application，the maximum length is 100，optional
	Version string `json:"version" gorm:"type:varchar(100)"`
	// OrgID Corporate identifier，optional
	OrgID int64 `json:"orgId" gorm:"index:idx_org_id"`
	// ProjectID Project identifier，optional
	ProjectID int64 `json:"projectId"`
	// ApplicationID Application identifier，optional
	ApplicationID int64 `json:"applicationId"`
	// ProjectName Project name，optional
	ProjectName string `json:"projectName" gorm:"type:varchar(80)"`
	// ApplicationName Application name，optional
	ApplicationName string `json:"applicationName" gorm:"type:varchar(80)"`
	// UserID User identifier，the maximum length is 50，optional
	UserID string `json:"userId" gorm:"type:varchar(50)"`
	// ClusterName Cluster Name，the maximum length is 80，optional
	ClusterName string `json:"clusterName" gorm:"type:varchar(80)"` // 所属集群
	// Resources Specify the release resource type and resource storage path, optional
	Resources string `json:"resources,omitempty" gorm:"type:text"`
	// Reference Number of deployments, when it is 0, can be clear
	Reference int64 `json:"reference"`
	// CrossCluster Indicates whether the current release can cross clusters, without cluster restrictions
	CrossCluster bool `json:"crossCluster"`
	// CreatedAt Release created time
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt Release updated time
	UpdatedAt time.Time `json:"updatedAt"`
	// IsLatest 是否为分支最新
	IsLatest bool `json:"isLatest"`
}

// Set table name
func (Release) TableName() string {
	return "dice_release"
}
