package core_dtos

type PushNotificationDto struct {
	UserEmailId     int    `json:"user_email_id" validate:"required"` // Required to specify the recipient
	Type            string `json:"type" validate:"required"`          // Required to specify the notification type
	Message         string `json:"message" validate:"required"`       // Required for the notification content
	RelatedItemId   int    `json:"related_item_id,omitempty"`         // Optional: ID of related item
	RelatedItemType string `json:"related_item_type,omitempty"`       // Optional: Type of related item
	ExtraData       string `json:"extra_data,omitempty"`              // Optional: Additional data
}
