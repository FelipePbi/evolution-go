package whatsmeow_service

import (
	"testing"

	instance_model "github.com/EvolutionAPI/evolution-go/pkg/instance/model"
)

func TestResolveInstanceWebhookUrl(t *testing.T) {
	tests := []struct {
		name         string
		evolutionEnv string
		instance     *instance_model.Instance
		want         string
	}{
		{
			name:         "production uses default webhook",
			evolutionEnv: "production",
			instance: &instance_model.Instance{
				Webhook:      "https://api.example.com/webhook",
				WebhookLocal: "http://localhost:3000/webhook",
			},
			want: "https://api.example.com/webhook",
		},
		{
			name:         "local uses local webhook",
			evolutionEnv: "local",
			instance: &instance_model.Instance{
				Webhook:      "https://api.example.com/webhook",
				WebhookLocal: "http://localhost:3000/webhook",
			},
			want: "http://localhost:3000/webhook",
		},
		{
			name:         "local without local webhook falls back to default webhook",
			evolutionEnv: "local",
			instance: &instance_model.Instance{
				Webhook: "https://api.example.com/webhook",
			},
			want: "https://api.example.com/webhook",
		},
		{
			name:         "local disabled local webhook disables effective webhook",
			evolutionEnv: "local",
			instance: &instance_model.Instance{
				Webhook:      "https://api.example.com/webhook",
				WebhookLocal: "disabled",
			},
			want: "disabled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveInstanceWebhookUrl(tt.evolutionEnv, tt.instance); got != tt.want {
				t.Fatalf("resolveInstanceWebhookUrl() = %q, want %q", got, tt.want)
			}
		})
	}
}
