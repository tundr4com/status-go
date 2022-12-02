package communities

import "errors"

var ErrChatNotFound = errors.New("chat not found")
var ErrCategoryNotFound = errors.New("category not found")
var ErrNoChangeInPosition = errors.New("no change in category position")
var ErrChatAlreadyAssigned = errors.New("chat already assigned to a category")
var ErrOrgNotFound = errors.New("community not found")
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
var ErrInvalidGrant = errors.New("invalid grant")
var ErrNotAuthorized = errors.New("not authorized")
var ErrAlreadyMember = errors.New("already a member")
var ErrAlreadyJoined = errors.New("already joined")
var ErrInvalidMessage = errors.New("invalid community description message")
var ErrMemberNotFound = errors.New("member not found")
