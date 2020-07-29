// Code generated by protoc-gen-a2config. DO NOT EDIT.
// source: config/deployment/automate_config.proto

package deployment

import (
	applications "github.com/chef/automate/api/config/applications"
	authn "github.com/chef/automate/api/config/authn"
	authz "github.com/chef/automate/api/config/authz"
	backupgateway "github.com/chef/automate/api/config/backup_gateway"
	bifrost "github.com/chef/automate/api/config/bifrost"
	bookshelf "github.com/chef/automate/api/config/bookshelf"
	builderapi "github.com/chef/automate/api/config/builder_api"
	builderapiproxy "github.com/chef/automate/api/config/builder_api_proxy"
	buildermemcached "github.com/chef/automate/api/config/builder_memcached"
	cds "github.com/chef/automate/api/config/cds"
	cereal "github.com/chef/automate/api/config/cereal"
	cfgmgmt "github.com/chef/automate/api/config/cfgmgmt"
	compliance "github.com/chef/automate/api/config/compliance"
	csnginx "github.com/chef/automate/api/config/cs_nginx"
	datafeed "github.com/chef/automate/api/config/data_feed"
	dex "github.com/chef/automate/api/config/dex"
	elasticsearch "github.com/chef/automate/api/config/elasticsearch"
	erchef "github.com/chef/automate/api/config/erchef"
	essidecar "github.com/chef/automate/api/config/es_sidecar"
	esgateway "github.com/chef/automate/api/config/esgateway"
	event "github.com/chef/automate/api/config/event"
	eventfeed "github.com/chef/automate/api/config/event_feed"
	eventgateway "github.com/chef/automate/api/config/event_gateway"
	gateway "github.com/chef/automate/api/config/gateway"
	infraproxy "github.com/chef/automate/api/config/infra_proxy"
	ingest "github.com/chef/automate/api/config/ingest"
	licensecontrol "github.com/chef/automate/api/config/license_control"
	loadbalancer "github.com/chef/automate/api/config/load_balancer"
	localuser "github.com/chef/automate/api/config/local_user"
	minio "github.com/chef/automate/api/config/minio"
	nodemanager "github.com/chef/automate/api/config/nodemanager"
	notifications "github.com/chef/automate/api/config/notifications"
	pggateway "github.com/chef/automate/api/config/pg_gateway"
	pgsidecar "github.com/chef/automate/api/config/pg_sidecar"
	postgresql "github.com/chef/automate/api/config/postgresql"
	prometheus "github.com/chef/automate/api/config/prometheus"
	secrets "github.com/chef/automate/api/config/secrets"
	session "github.com/chef/automate/api/config/session"
	shared "github.com/chef/automate/api/config/shared"
	teams "github.com/chef/automate/api/config/teams"
	ui "github.com/chef/automate/api/config/ui"
	workflownginx "github.com/chef/automate/api/config/workflow_nginx"
	workflowserver "github.com/chef/automate/api/config/workflow_server"
)

// NewAutomateConfig returns a new instance of AutomateConfig with zero values.
func NewAutomateConfig() *AutomateConfig {
	return &AutomateConfig{
		Applications:     applications.NewConfigRequest(),
		AuthN:            authn.NewConfigRequest(),
		AuthZ:            authz.NewConfigRequest(),
		BackupGateway:    backupgateway.NewConfigRequest(),
		Bifrost:          bifrost.NewConfigRequest(),
		Bookshelf:        bookshelf.NewConfigRequest(),
		BuilderApi:       builderapi.NewConfigRequest(),
		BuilderApiProxy:  builderapiproxy.NewConfigRequest(),
		BuilderMemcached: buildermemcached.NewConfigRequest(),
		Cds:              cds.NewConfigRequest(),
		Cereal:           cereal.NewConfigRequest(),
		Compliance:       compliance.NewConfigRequest(),
		ConfigMgmt:       cfgmgmt.NewConfigRequest(),
		CsNginx:          csnginx.NewConfigRequest(),
		DataFeedService:  datafeed.NewConfigRequest(),
		Deployment:       NewConfigRequest(),
		Dex:              dex.NewConfigRequest(),
		Elasticsearch:    elasticsearch.NewConfigRequest(),
		Erchef:           erchef.NewConfigRequest(),
		EsSidecar:        essidecar.NewConfigRequest(),
		Esgateway:        esgateway.NewConfigRequest(),
		EventFeedService: eventfeed.NewConfigRequest(),
		EventGateway:     eventgateway.NewConfigRequest(),
		EventService:     event.NewConfigRequest(),
		Gateway:          gateway.NewConfigRequest(),
		Global:           shared.NewGlobalConfig(),
		InfraProxy:       infraproxy.NewConfigRequest(),
		Ingest:           ingest.NewConfigRequest(),
		LicenseControl:   licensecontrol.NewConfigRequest(),
		LoadBalancer:     loadbalancer.NewConfigRequest(),
		LocalUser:        localuser.NewConfigRequest(),
		Minio:            minio.NewConfigRequest(),
		Nodemanager:      nodemanager.NewConfigRequest(),
		Notifications:    notifications.NewConfigRequest(),
		PgGateway:        pggateway.NewConfigRequest(),
		PgSidecar:        pgsidecar.NewConfigRequest(),
		Postgresql:       postgresql.NewConfigRequest(),
		Prometheus:       prometheus.NewConfigRequest(),
		Secrets:          secrets.NewConfigRequest(),
		Session:          session.NewConfigRequest(),
		Teams:            teams.NewConfigRequest(),
		UI:               ui.NewConfigRequest(),
		Workflow:         workflowserver.NewConfigRequest(),
		WorkflowNginx:    workflownginx.NewConfigRequest(),
	}
}

// DefaultAutomateConfig returns a new instance of Automate config with default values.
func DefaultAutomateConfig() *AutomateConfig {
	return &AutomateConfig{
		Applications:     applications.DefaultConfigRequest(),
		AuthN:            authn.DefaultConfigRequest(),
		AuthZ:            authz.DefaultConfigRequest(),
		BackupGateway:    backupgateway.DefaultConfigRequest(),
		Bifrost:          bifrost.DefaultConfigRequest(),
		Bookshelf:        bookshelf.DefaultConfigRequest(),
		BuilderApi:       builderapi.DefaultConfigRequest(),
		BuilderApiProxy:  builderapiproxy.DefaultConfigRequest(),
		BuilderMemcached: buildermemcached.DefaultConfigRequest(),
		Cds:              cds.DefaultConfigRequest(),
		Cereal:           cereal.DefaultConfigRequest(),
		Compliance:       compliance.DefaultConfigRequest(),
		ConfigMgmt:       cfgmgmt.DefaultConfigRequest(),
		CsNginx:          csnginx.DefaultConfigRequest(),
		DataFeedService:  datafeed.DefaultConfigRequest(),
		Deployment:       DefaultConfigRequest(),
		Dex:              dex.DefaultConfigRequest(),
		Elasticsearch:    elasticsearch.DefaultConfigRequest(),
		Erchef:           erchef.DefaultConfigRequest(),
		EsSidecar:        essidecar.DefaultConfigRequest(),
		Esgateway:        esgateway.DefaultConfigRequest(),
		EventFeedService: eventfeed.DefaultConfigRequest(),
		EventGateway:     eventgateway.DefaultConfigRequest(),
		EventService:     event.DefaultConfigRequest(),
		Gateway:          gateway.DefaultConfigRequest(),
		Global:           shared.DefaultGlobalConfig(),
		InfraProxy:       infraproxy.DefaultConfigRequest(),
		Ingest:           ingest.DefaultConfigRequest(),
		LicenseControl:   licensecontrol.DefaultConfigRequest(),
		LoadBalancer:     loadbalancer.DefaultConfigRequest(),
		LocalUser:        localuser.DefaultConfigRequest(),
		Minio:            minio.DefaultConfigRequest(),
		Nodemanager:      nodemanager.DefaultConfigRequest(),
		Notifications:    notifications.DefaultConfigRequest(),
		PgGateway:        pggateway.DefaultConfigRequest(),
		PgSidecar:        pgsidecar.DefaultConfigRequest(),
		Postgresql:       postgresql.DefaultConfigRequest(),
		Prometheus:       prometheus.DefaultConfigRequest(),
		Secrets:          secrets.DefaultConfigRequest(),
		Session:          session.DefaultConfigRequest(),
		Teams:            teams.DefaultConfigRequest(),
		UI:               ui.DefaultConfigRequest(),
		Workflow:         workflowserver.DefaultConfigRequest(),
		WorkflowNginx:    workflownginx.DefaultConfigRequest(),
	}
}

/*
Validate verifies that all required configuration keys are present
and enforces other invariants on configuration option values.

If the configuration is valid, the returned error is nil.
*/
func (c *AutomateConfig) Validate() error {
	err := shared.Validate(c.Global.Validate(), c.AuthN.Validate(), c.AuthZ.Validate(), c.Compliance.Validate(), c.ConfigMgmt.Validate(), c.Deployment.Validate(), c.Dex.Validate(), c.Elasticsearch.Validate(), c.Esgateway.Validate(), c.EsSidecar.Validate(), c.Gateway.Validate(), c.Ingest.Validate(), c.LoadBalancer.Validate(), c.LocalUser.Validate(), c.LicenseControl.Validate(), c.Notifications.Validate(), c.Postgresql.Validate(), c.Session.Validate(), c.Teams.Validate(), c.UI.Validate(), c.Secrets.Validate(), c.BackupGateway.Validate(), c.PgSidecar.Validate(), c.PgGateway.Validate(), c.Applications.Validate(), c.Bookshelf.Validate(), c.Bifrost.Validate(), c.Erchef.Validate(), c.CsNginx.Validate(), c.Workflow.Validate(), c.WorkflowNginx.Validate(), c.EventService.Validate(), c.Nodemanager.Validate(), c.EventGateway.Validate(), c.Prometheus.Validate(), c.DataFeedService.Validate(), c.EventFeedService.Validate(), c.Cereal.Validate(), c.BuilderApi.Validate(), c.BuilderApiProxy.Validate(), c.Minio.Validate(), c.BuilderMemcached.Validate(), c.InfraProxy.Validate(), c.Cds.Validate())
	if err == nil {
		return nil
	}
	cfgErr, ok := err.(shared.Error)
	if ok {
		if cfgErr.IsEmpty() {
			return nil
		}
	}
	return err
}

// SetGlobalConfig iterates over the AutomateConfig and applies global configuration
func (c *AutomateConfig) SetGlobalConfig() {
	c.AuthN.SetGlobalConfig(c.Global)
	c.AuthZ.SetGlobalConfig(c.Global)
	c.Compliance.SetGlobalConfig(c.Global)
	c.ConfigMgmt.SetGlobalConfig(c.Global)
	c.Deployment.SetGlobalConfig(c.Global)
	c.Dex.SetGlobalConfig(c.Global)
	c.Elasticsearch.SetGlobalConfig(c.Global)
	c.Esgateway.SetGlobalConfig(c.Global)
	c.EsSidecar.SetGlobalConfig(c.Global)
	c.Gateway.SetGlobalConfig(c.Global)
	c.Ingest.SetGlobalConfig(c.Global)
	c.LoadBalancer.SetGlobalConfig(c.Global)
	c.LocalUser.SetGlobalConfig(c.Global)
	c.LicenseControl.SetGlobalConfig(c.Global)
	c.Notifications.SetGlobalConfig(c.Global)
	c.Postgresql.SetGlobalConfig(c.Global)
	c.Session.SetGlobalConfig(c.Global)
	c.Teams.SetGlobalConfig(c.Global)
	c.UI.SetGlobalConfig(c.Global)
	c.Secrets.SetGlobalConfig(c.Global)
	c.BackupGateway.SetGlobalConfig(c.Global)
	c.PgSidecar.SetGlobalConfig(c.Global)
	c.PgGateway.SetGlobalConfig(c.Global)
	c.Applications.SetGlobalConfig(c.Global)
	c.Bookshelf.SetGlobalConfig(c.Global)
	c.Bifrost.SetGlobalConfig(c.Global)
	c.Erchef.SetGlobalConfig(c.Global)
	c.CsNginx.SetGlobalConfig(c.Global)
	c.Workflow.SetGlobalConfig(c.Global)
	c.WorkflowNginx.SetGlobalConfig(c.Global)
	c.EventService.SetGlobalConfig(c.Global)
	c.Nodemanager.SetGlobalConfig(c.Global)
	c.EventGateway.SetGlobalConfig(c.Global)
	c.Prometheus.SetGlobalConfig(c.Global)
	c.DataFeedService.SetGlobalConfig(c.Global)
	c.EventFeedService.SetGlobalConfig(c.Global)
	c.Cereal.SetGlobalConfig(c.Global)
	c.BuilderApi.SetGlobalConfig(c.Global)
	c.BuilderApiProxy.SetGlobalConfig(c.Global)
	c.Minio.SetGlobalConfig(c.Global)
	c.BuilderMemcached.SetGlobalConfig(c.Global)
	c.InfraProxy.SetGlobalConfig(c.Global)
	c.Cds.SetGlobalConfig(c.Global)
}

// PlatformServiceConfigForService gets the config for the service by name
func (c *AutomateConfig) PlatformServiceConfigForService(serviceName string) (shared.PlatformServiceConfigurable, bool) {
	switch serviceName {
	case "authn-service":
		return c.AuthN, true
	case "authz-service":
		return c.AuthZ, true
	case "compliance-service":
		return c.Compliance, true
	case "config-mgmt-service":
		return c.ConfigMgmt, true
	case "deployment-service":
		return c.Deployment, true
	case "automate-dex":
		return c.Dex, true
	case "automate-elasticsearch":
		return c.Elasticsearch, true
	case "automate-es-gateway":
		return c.Esgateway, true
	case "es-sidecar-service":
		return c.EsSidecar, true
	case "automate-gateway":
		return c.Gateway, true
	case "ingest-service":
		return c.Ingest, true
	case "automate-load-balancer":
		return c.LoadBalancer, true
	case "local-user-service":
		return c.LocalUser, true
	case "license-control-service":
		return c.LicenseControl, true
	case "notifications-service":
		return c.Notifications, true
	case "automate-postgresql":
		return c.Postgresql, true
	case "session-service":
		return c.Session, true
	case "teams-service":
		return c.Teams, true
	case "automate-ui":
		return c.UI, true
	case "secrets-service":
		return c.Secrets, true
	case "backup-gateway":
		return c.BackupGateway, true
	case "pg-sidecar-service":
		return c.PgSidecar, true
	case "automate-pg-gateway":
		return c.PgGateway, true
	case "applications-service":
		return c.Applications, true
	case "automate-cs-bookshelf":
		return c.Bookshelf, true
	case "automate-cs-oc-bifrost":
		return c.Bifrost, true
	case "automate-cs-oc-erchef":
		return c.Erchef, true
	case "automate-cs-nginx":
		return c.CsNginx, true
	case "automate-workflow-server":
		return c.Workflow, true
	case "automate-workflow-nginx":
		return c.WorkflowNginx, true
	case "event-service":
		return c.EventService, true
	case "nodemanager-service":
		return c.Nodemanager, true
	case "event-gateway":
		return c.EventGateway, true
	case "automate-prometheus":
		return c.Prometheus, true
	case "data-feed-service":
		return c.DataFeedService, true
	case "event-feed-service":
		return c.EventFeedService, true
	case "cereal-service":
		return c.Cereal, true
	case "automate-builder-api":
		return c.BuilderApi, true
	case "automate-builder-api-proxy":
		return c.BuilderApiProxy, true
	case "automate-minio":
		return c.Minio, true
	case "automate-builder-memcached":
		return c.BuilderMemcached, true
	case "infra-proxy-service":
		return c.InfraProxy, true
	case "automate-cds":
		return c.Cds, true
	default:
		return nil, false
	}
}
