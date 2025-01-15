package handler

import (
	"github.com/gofiber/fiber/v2"
	"samsamoohooh-api/internal/application/port"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/router"
)

type GroupHandler struct {
	router       *router.Router
	groupService port.GroupService
}

func NewGroupHandler(
	groupService port.GroupService,
	router *router.Router,
) *GroupHandler {
	groupHandler := &GroupHandler{
		router:       router,
		groupService: groupService,
	}

	groupHandler.Route()
	return groupHandler
}

func (h *GroupHandler) Route() {
	groups := h.router.ApiRouter.Group("/groups")
	{
		groups.Post("/", h.CreateGroup)
		groups.Get("/:id", h.FindGroup)
	}
}

// CreateGroup godoc
//
//	@Tags		groups
//	@Produce	json
//	@Param		CreateGroupRequest	body		presenter.CreateGroupRequest	true	"Create Group Request"
//	@Success	201					{object}	presenter.CreateGroupResponse
//	@Router		/api/groups [post]
func (h *GroupHandler) CreateGroup(c *fiber.Ctx) error {
	var _ = &presenter.CreateGroupRequest{}
	return nil
}

// FindGroup godoc
//
//	@Tags		groups
//	@Produce	json
//	@Param		id	path		int	true	"Group ID"
//	@Success	200	{object}	presenter.FindGroupResponse
//	@Router		/api/groups/{id} [get]
func (h *GroupHandler) FindGroup(c *fiber.Ctx) error {
	var _ = &presenter.FindGroupRequest{}
	return nil
}

// ListGroups godoc
//
//	@Tags		groups
//	@Produce	json
//	@Param		ListGroupsRequest	body		presenter.ListGroupsRequest	true	"List Groups Request"
//	@Success	200					{object}	presenter.ListGroupsResponse
//	@Router		/api/groups [get]
func (h *GroupHandler) ListGroups(c *fiber.Ctx) error {
	var _ = &presenter.ListGroupsRequest{}
	return nil
}
