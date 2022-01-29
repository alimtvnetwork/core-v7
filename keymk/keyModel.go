package keymk

type keyModel struct {
	Option        *Option  `json:"Option,omitempty"`
	MainName      string   `json:"MainName,omitempty"`
	KeyChains     []string `json:"KeyChains,omitempty"`
	CompiledChain *string  `json:"CompiledChain,omitempty"`
}
