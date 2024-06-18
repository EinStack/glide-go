package glide

// https://github.com/EinStack/glide/tree/develop/pkg/api/schemas

// RouterList is a list of all router configurations.
type RouterList struct {
	Routers []RouterConfig `json:"routers"`
}

// RouterConfig is a single router configuration.
type RouterConfig any
