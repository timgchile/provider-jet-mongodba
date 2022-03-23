/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	configuration "github.com/timgchile/provider-jet-mongodba/internal/controller/alert/configuration"
	backupschedule "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/backupschedule"
	backupsnapshot "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/backupsnapshot"
	backupsnapshotrestorejob "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/backupsnapshotrestorejob"
	provideraccess "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/provideraccess"
	provideraccessauthorization "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/provideraccessauthorization"
	provideraccesssetup "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/provideraccesssetup"
	providersnapshot "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/providersnapshot"
	providersnapshotbackuppolicy "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/providersnapshotbackuppolicy"
	providersnapshotrestorejob "github.com/timgchile/provider-jet-mongodba/internal/controller/cloud/providersnapshotrestorejob"
	dbrole "github.com/timgchile/provider-jet-mongodba/internal/controller/custom/dbrole"
	dnsconfigurationclusteraws "github.com/timgchile/provider-jet-mongodba/internal/controller/custom/dnsconfigurationclusteraws"
	lake "github.com/timgchile/provider-jet-mongodba/internal/controller/data/lake"
	user "github.com/timgchile/provider-jet-mongodba/internal/controller/database/user"
	trigger "github.com/timgchile/provider-jet-mongodba/internal/controller/event/trigger"
	clusterconfig "github.com/timgchile/provider-jet-mongodba/internal/controller/global/clusterconfig"
	configurationldap "github.com/timgchile/provider-jet-mongodba/internal/controller/ldap/configuration"
	verify "github.com/timgchile/provider-jet-mongodba/internal/controller/ldap/verify"
	window "github.com/timgchile/provider-jet-mongodba/internal/controller/maintenance/window"
	advancedcluster "github.com/timgchile/provider-jet-mongodba/internal/controller/mongodba/advancedcluster"
	project "github.com/timgchile/provider-jet-mongodba/internal/controller/mongodba/project"
	auditing "github.com/timgchile/provider-jet-mongodba/internal/controller/mongodbatlas/auditing"
	cluster "github.com/timgchile/provider-jet-mongodba/internal/controller/mongodbatlas/cluster"
	team "github.com/timgchile/provider-jet-mongodba/internal/controller/mongodbatlas/team"
	container "github.com/timgchile/provider-jet-mongodba/internal/controller/network/container"
	peering "github.com/timgchile/provider-jet-mongodba/internal/controller/network/peering"
	archive "github.com/timgchile/provider-jet-mongodba/internal/controller/online/archive"
	invitation "github.com/timgchile/provider-jet-mongodba/internal/controller/org/invitation"
	ipmode "github.com/timgchile/provider-jet-mongodba/internal/controller/private/ipmode"
	endpoint "github.com/timgchile/provider-jet-mongodba/internal/controller/privatelink/endpoint"
	endpointservice "github.com/timgchile/provider-jet-mongodba/internal/controller/privatelink/endpointservice"
	endpointserviceadl "github.com/timgchile/provider-jet-mongodba/internal/controller/privatelink/endpointserviceadl"
	invitationproject "github.com/timgchile/provider-jet-mongodba/internal/controller/project/invitation"
	ipaccesslist "github.com/timgchile/provider-jet-mongodba/internal/controller/project/ipaccesslist"
	providerconfig "github.com/timgchile/provider-jet-mongodba/internal/controller/providerconfig"
	index "github.com/timgchile/provider-jet-mongodba/internal/controller/search/index"
	partyintegration "github.com/timgchile/provider-jet-mongodba/internal/controller/third/partyintegration"
	authenticationdatabaseuser "github.com/timgchile/provider-jet-mongodba/internal/controller/x509/authenticationdatabaseuser"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.Setup,
		backupschedule.Setup,
		backupsnapshot.Setup,
		backupsnapshotrestorejob.Setup,
		provideraccess.Setup,
		provideraccessauthorization.Setup,
		provideraccesssetup.Setup,
		providersnapshot.Setup,
		providersnapshotbackuppolicy.Setup,
		providersnapshotrestorejob.Setup,
		dbrole.Setup,
		dnsconfigurationclusteraws.Setup,
		lake.Setup,
		user.Setup,
		trigger.Setup,
		clusterconfig.Setup,
		configurationldap.Setup,
		verify.Setup,
		window.Setup,
		advancedcluster.Setup,
		project.Setup,
		auditing.Setup,
		cluster.Setup,
		team.Setup,
		container.Setup,
		peering.Setup,
		archive.Setup,
		invitation.Setup,
		ipmode.Setup,
		endpoint.Setup,
		endpointservice.Setup,
		endpointserviceadl.Setup,
		invitationproject.Setup,
		ipaccesslist.Setup,
		providerconfig.Setup,
		index.Setup,
		partyintegration.Setup,
		authenticationdatabaseuser.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
