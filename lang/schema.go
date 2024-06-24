package lang

// https://github.com/EinStack/glide/tree/develop/pkg/api/schemas

// RouterList is a list of all router configurations.
type RouterList struct {
	Routers []RouterConfig `json:"routers"`
}

// RouterConfig is a single router configuration.
type RouterConfig map[string]any

// ChatRequest is a unified chat request across all language models.
type ChatRequest struct {
	Message        ChatMessage                     `json:"message"`
	MessageHistory *[]ChatMessage                  `json:"message_history,omitempty"`
	OverrideParams *map[string]OverrideChatRequest `json:"override_params,omitempty"`
}

// OverrideChatRequest is an override of a single chat request.
type OverrideChatRequest struct {
	Message ChatMessage `json:"message"`
}

// NewChatRequest instantiates a new ChatRequest.
func NewChatRequest(content string) ChatRequest {
	message := ChatMessage{Content: content}
	return ChatRequest{Message: message}
}

// ChatResponse is a unified chat response across all language models.
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

// Content returns the content of the response.
func (r *ChatResponse) Content() string {
	return r.ModelResponse.Message.Content
}

// ModelResponse is unified response from the provider.
type ModelResponse struct {
	Metadata   map[string]string `json:"metadata,omitempty"`
	Message    ChatMessage       `json:"message"`
	TokenUsage TokenUsage        `json:"token_usage"`
}

// TokenUsage is a list of prompt, response and total token usage.
type TokenUsage struct {
	PromptTokens   int `json:"prompt_tokens"`
	ResponseTokens int `json:"response_tokens"`
	TotalTokens    int `json:"total_tokens"`
}

// ChatMessage is content and role of the message.
type ChatMessage struct {
	// The role of the author of this message.
	// One of system, user, or assistant.
	Role *string `json:"role"`
	// The content of the message.
	Content string `json:"content"`
	// The name of the author of this message.
	// May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name string `json:"name,omitempty"`
}
