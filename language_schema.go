package glide

// https://github.com/EinStack/glide/tree/develop/pkg/api/schemas

// RouterConfig TODO.
type RouterConfig any

// RouterList TODO.
type RouterList struct {
	Routers []RouterConfig `json:"routers"`
}

// ChatRequest TODO.
type ChatRequest struct {
	Message        ChatMessage          `json:"message" validate:"required"`
	MessageHistory []ChatMessage        `json:"message_history"`
	OverrideParams *OverrideChatRequest `json:"override_params,omitempty"`
}

type OverrideChatRequest struct {
	ModelID string      `json:"model_id" validate:"required"`
	Message ChatMessage `json:"message" validate:"required"`
}

// NewChatRequest instantiates a new ChatRequest.
func NewChatRequest() ChatRequest {
	// TODO.
	return ChatRequest{}
}

// ChatResponse TODO.
type ChatResponse struct {
	ID            string        `json:"id,omitempty"`
	Created       int           `json:"created_at,omitempty"`
	Provider      string        `json:"provider_id,omitempty"`
	RouterID      string        `json:"router_id,omitempty"`
	ModelID       string        `json:"model_id,omitempty"`
	ModelName     string        `json:"model_name,omitempty"`
	Cached        bool          `json:"cached,omitempty"`
	ModelResponse ModelResponse `json:"model_response,omitempty"`
}

type ModelResponse struct {
	Metadata   map[string]string `json:"metadata,omitempty"`
	Message    ChatMessage       `json:"message"`
	TokenUsage TokenUsage        `json:"token_usage"`
}

type TokenUsage struct {
	PromptTokens   int `json:"prompt_tokens"`
	ResponseTokens int `json:"response_tokens"`
	TotalTokens    int `json:"total_tokens"`
}

type ChatMessage struct {
	// The role of the author of this message. One of system, user, or assistant.
	Role string `json:"role" validate:"required"`
	// The content of the message.
	Content string `json:"content" validate:"required"`
	// The name of the author of this message. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name string `json:"name,omitempty"`
}
