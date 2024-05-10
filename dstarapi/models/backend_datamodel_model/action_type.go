package backend_datamodel_model
import (
    "errors"
)
type ActionType int

const (
    POST_SUPER_LIKE_ACTIONTYPE ActionType = iota
    POST_LIKE_ACTIONTYPE
    POST_DISLIKE_ACTIONTYPE
    POST_VIEW_ACTIONTYPE
    POST_SHARE_ACTIONTYPE
    POST_SAVE_ACTIONTYPE
    POST_UNSAVE_ACTIONTYPE
    PROFILE_VIEW_ACTIONTYPE
)

func (i ActionType) String() string {
    return []string{"post_super_like", "post_like", "post_dislike", "post_view", "post_share", "post_save", "post_unsave", "profile_view"}[i]
}
func ParseActionType(v string) (any, error) {
    result := POST_SUPER_LIKE_ACTIONTYPE
    switch v {
        case "post_super_like":
            result = POST_SUPER_LIKE_ACTIONTYPE
        case "post_like":
            result = POST_LIKE_ACTIONTYPE
        case "post_dislike":
            result = POST_DISLIKE_ACTIONTYPE
        case "post_view":
            result = POST_VIEW_ACTIONTYPE
        case "post_share":
            result = POST_SHARE_ACTIONTYPE
        case "post_save":
            result = POST_SAVE_ACTIONTYPE
        case "post_unsave":
            result = POST_UNSAVE_ACTIONTYPE
        case "profile_view":
            result = PROFILE_VIEW_ACTIONTYPE
        default:
            return 0, errors.New("Unknown ActionType value: " + v)
    }
    return &result, nil
}
func SerializeActionType(values []ActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ActionType) isMultiValue() bool {
    return false
}
