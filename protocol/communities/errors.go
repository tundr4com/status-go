package communities

import "errors"

var ErrChatNotFound = errors.New("chat not found")
var ErrCategoryNotFound = errors.New("category not found")
var ErrNoChangeInPosition = errors.New("no change in category position")
var ErrChatAlreadyAssigned = errors.New("chat already assigned to a category")
var ErrOrgNotFound = errors.New("community not found")
var ErrOrgAlreadyJoined = errors.New("community already joined")
var ErrChatAlreadyExists = errors.New("chat already exists")
var ErrCategoryAlreadyExists = errors.New("category already exists")
var ErrCantRequestAccess = errors.New("can't request access")
var ErrInvalidCommunityDescription = errors.New("invalid community description")
var ErrInvalidCommunityDescriptionNoOrgPermissions = errors.New("invalid community description no org permissions")
var ErrInvalidCommunityDescriptionNoChatPermissions = errors.New("invalid community description no chat permissions")
var ErrInvalidCommunityDescriptionUnknownChatAccess = errors.New("invalid community description unknown chat access")
var ErrInvalidCommunityDescriptionUnknownOrgAccess = errors.New("invalid community description unknown org access")
var ErrInvalidCommunityDescriptionMemberInChatButNotInOrg = errors.New("invalid community description member in chat but not in org")
var ErrInvalidCommunityDescriptionCategoryNoID = errors.New("invalid community category id")
var ErrInvalidCommunityDescriptionCategoryNoName = errors.New("invalid community category name")
var ErrInvalidCommunityDescriptionChatIdentity = errors.New("invalid community chat name, missing")
var ErrInvalidCommunityDescriptionDuplicatedName = errors.New("invalid community chat name, duplicated")
var ErrInvalidCommunityDescriptionUnknownChatCategory = errors.New("invalid community category in chat")
var ErrInvalidCommunityTags = errors.New("invalid community tags")
var ErrNotAdmin = errors.New("no admin privileges for this community")
var ErrNotOwner = errors.New("no owner privileges for this community")
var ErrNotControlNode = errors.New("no owner private key found for this community")
var ErrInvalidGrant = errors.New("invalid grant")
var ErrNotAuthorized = errors.New("not authorized")
var ErrAlreadyMember = errors.New("already a member")
var ErrAlreadyJoined = errors.New("already joined")
var ErrInvalidMessage = errors.New("invalid community description message")
var ErrMemberNotFound = errors.New("member not found")
var ErrTokenPermissionAlreadyExists = errors.New("token permission already exists")
var ErrTokenPermissionNotFound = errors.New("token permission not found")
var ErrNoPermissionToJoin = errors.New("member has no permission to join")
var ErrMemberWalletAlreadyExists = errors.New("member wallet already exists")
var ErrMemberWalletNotFound = errors.New("member wallet not found")
var ErrNotEnoughPermissions = errors.New("not enough permissions for this community")
var ErrCannotRemoveOwnerOrAdmin = errors.New("not allowed to remove admin or owner")
var ErrCannotBanOwnerOrAdmin = errors.New("not allowed to ban admin or owner")
var ErrInvalidManageTokensPermission = errors.New("no privileges to manage tokens")
