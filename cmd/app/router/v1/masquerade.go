package v1

import (
	"firewall-api/code"
	"firewall-api/utils/dbus"
	q "firewall-api/utils/query"
	"github.com/gin-gonic/gin"
)

type MasqueradeRouter struct{}

func (this *MasqueradeRouter) RegisterPortAPI(g *gin.RouterGroup) {
	portGroup := g.Group("/masquerade")

	portGroup.GET("/enable", this.enableInRuntime)
	portGroup.GET("/disable", this.disableInRuntime)
	portGroup.GET("/query", this.queryInRuntime)
	portGroup.GET("/enablepermanent", this.enableInPermanent)
	portGroup.GET("/disablepermanent", this.disableInPermanent)
	portGroup.GET("/querypermanent", this.queryInPermanent)
}

// enableInRuntime ...
// @Summary enableInRuntime
// @Produce  json
// @Success 200 {object} internal.Response
// @Router /fw/v1/port/enable [GET]
func (this *MasqueradeRouter) enableInRuntime(c *gin.Context) {

	var query = &q.Query{}
	err := c.Bind(query)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	dbusClient, err := dbus.NewDbusClientService(query.Ip)
	defer dbusClient.Destroy()
	if err != nil {
		q.ConnectDbusService(c, err)
		return
	}

	if err := dbusClient.EnableMasquerade(query.Zone, query.Timeout); err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	q.SuccessResponse(c, code.OK, nil)
}

// disableInRuntime ...
// @Summary disableInRuntime
// @Produce  json
// @Success 200 {object} internal.Response
// @Router /fw/v1/port/disable [GET]
func (this *MasqueradeRouter) disableInRuntime(c *gin.Context) {

	var query = &q.Query{}
	err := c.Bind(query)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	dbusClient, err := dbus.NewDbusClientService(query.Ip)
	defer dbusClient.Destroy()
	if err != nil {
		q.ConnectDbusService(c, err)
		return
	}

	if err := dbusClient.DisableMasquerade(query.Zone); err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	q.SuccessResponse(c, code.OK, nil)
}

// queryInRuntime ...
// @Summary queryInRuntime
// @Produce  json
// @Success 200 {object} internal.Response
// @Router /fw/v1/port/query [GET]
func (this *MasqueradeRouter) queryInRuntime(c *gin.Context) {

	var query = &q.Query{}
	err := c.Bind(query)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	dbusClient, err := dbus.NewDbusClientService(query.Ip)
	defer dbusClient.Destroy()
	if err != nil {
		q.ConnectDbusService(c, err)
		return
	}

	isenable, err := dbusClient.QueryMasquerade(query.Zone)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	if isenable == false {
		q.SuccessResponse(c, code.NETWORK_MASQUERADE_DISABLE, isenable)
		return
	}

	q.SuccessResponse(c, code.NETWORK_MASQUERADE_ENABLE, isenable)
}

// enableInPermanent ...
// @Summary enableInPermanent
// @Produce  json
// @Success 200 {object} internal.Response
// @Router /fw/v1/port/enableinpermanent [GET]
func (this *MasqueradeRouter) enableInPermanent(c *gin.Context) {

	var query = &q.Query{}
	err := c.Bind(query)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	dbusClient, err := dbus.NewDbusClientService(query.Ip)
	defer dbusClient.Destroy()
	if err != nil {
		q.ConnectDbusService(c, err)
		return
	}

	if err := dbusClient.PermanentEnableMasquerade(query.Zone); err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	q.SuccessResponse(c, code.OK, nil)
}

// disableInPermanent ...
// @Summary disableInPermanent
// @Produce  json
// @Success 200 {object} internal.Response
// @Router /fw/v1/port/disablepermanent [GET]
func (this *MasqueradeRouter) disableInPermanent(c *gin.Context) {

	var query = &q.Query{}
	err := c.Bind(query)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	dbusClient, err := dbus.NewDbusClientService(query.Ip)
	defer dbusClient.Destroy()
	if err != nil {
		q.ConnectDbusService(c, err)
		return
	}

	if err := dbusClient.PermanentDisableMasquerade(query.Zone); err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	q.SuccessResponse(c, code.OK, nil)
}

// queryInPermanent ...
// @Summary queryInPermanent
// @Produce  json
// @Success 200 {object} internal.Response
// @Router /fw/v1/port/querypermanent [GET]
func (this *MasqueradeRouter) queryInPermanent(c *gin.Context) {

	var query = &q.Query{}
	err := c.Bind(query)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	dbusClient, err := dbus.NewDbusClientService(query.Ip)
	defer dbusClient.Destroy()
	if err != nil {
		q.ConnectDbusService(c, err)
		return
	}

	isenable, err := dbusClient.PermanentQueryMasquerade(query.Zone)

	if err != nil {
		q.APIResponse(c, err, nil)
		return
	}

	if isenable == false {
		q.SuccessResponse(c, code.NETWORK_MASQUERADE_DISABLE, isenable)
		return
	}

	q.SuccessResponse(c, code.NETWORK_MASQUERADE_ENABLE, isenable)
}
